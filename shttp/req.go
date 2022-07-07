package shttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

var client *fasthttp.Client

func Get(url string) ([]byte, error) {
	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")
	client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		// increase DNS cache time to an hour instead of default minute
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Add("User-Agent", getUserAgent())
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		return []byte(""), err
	}
	return resp.Body(), err
}

func Post(url string, data interface{}) {
	jsonType := []byte("application/json")
	// per-request timeout
	reqTimeout := time.Duration(5) * time.Second

	reqEntityBytes, _ := json.Marshal(data)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Add("User-Agent", getUserAgent())
	req.Header.SetContentTypeBytes(jsonType)
	req.SetBodyRaw(reqEntityBytes)
	resp := fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		fmt.Printf("DEBUG Response: %s\n", respBody)
		if statusCode == http.StatusOK {
			// respEntity := &Entity{}
			// err = json.Unmarshal(respBody, respEntity)
			// if err == io.EOF || err == nil {
			// 	fmt.Printf("DEBUG Parsed Response: %v\n", respEntity)
			// } else {
			// 	fmt.Fprintf(os.Stderr, "ERR failed to parse reponse: %v\n", err)
			// }
		} else {
			fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)
		}
	} else {
		errName, known := httpConnError(err)
		if known {
			fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
		} else {
			fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
		}
	}
	fasthttp.ReleaseResponse(resp)
}

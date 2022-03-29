package zdns

import (
	"fmt"
	"net/url"

	"github.com/gsmlg-dev/gsmlg-golang/errorhandler"
)

var exitIfError = errorhandler.CreateExitIfError("ZDNS")
var host string = "https://cloud.zdns.cn"

var api *ApiService = NewApi()

type ApiService struct {
	baseUrl string
	token   string
}

type DataStruct struct {
	ResourceType string      `json:"resource_type"`
	Zdnsuser     string      `json:"zdnsuser"`
	Attrs        interface{} `json:"attrs"`
}

func NewApi() *ApiService {
	api := &ApiService{
		baseUrl: host,
	}
	return api
}

func SetToken(t string) {
	api.SetToken(t)
}

func (api *ApiService) SetHost(h string) {
	api.baseUrl = h
}

func (api *ApiService) SetToken(t string) {
	api.token = t
}

func (api *ApiService) GetAuthUrl() string {
	s := fmt.Sprintf("%s/%s", api.baseUrl, "auth_cmd")
	return s
}

func (api *ApiService) GetRRManagerUrl() *url.URL {
	u, err := url.Parse(api.baseUrl)
	exitIfError(err)
	u.Path = "/rrmanager"
	q := u.Query()
	q.Set("zdnsuser", api.token)
	u.RawQuery = q.Encode()
	return u
}

func (api *ApiService) RRManagerRequest() (*url.URL, DataStruct) {
	u, err := url.Parse(api.baseUrl)
	exitIfError(err)
	u.Path = "/rrmanager"
	d := DataStruct{
		Zdnsuser: api.token,
	}
	return u, d
}

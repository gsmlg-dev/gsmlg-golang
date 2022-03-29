package zdns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	RR_NORMAL           uint32 = 0x0001
	RR_PAUSE                   = 0x0002
	RR_RESERVED1               = 0x0004
	RR_RESERVED2               = 0x0008
	RR_FAILOVER_ENABLE         = 0x0010
	RR_FAILOVER_DISABLE        = 0x0020
	RR_FAILOVER_BACKUP         = 0x0040
	RR_QS_ENABLE               = 0x1000
	RR_QS_BACKUP               = 0x1001

	RR_MODULE = 0x10000
)

var FlagToString = map[uint32]string{
	RR_NORMAL:           "normal",
	RR_PAUSE:            "pause",
	RR_FAILOVER_ENABLE:  "monitor-up",
	RR_FAILOVER_DISABLE: "monitor-down",
	RR_FAILOVER_BACKUP:  "backup",
	RR_QS_ENABLE:        "qs-enable",
	RR_QS_BACKUP:        "qs-backup",
}

type Rr struct {
	Id          string `json:"id"`
	Zone        string `json:"zone"`
	Name        string `json:"name"`
	ReverseName string `json:"-"`
	Type        string `json:"type"`
	Ttl         int    `json:"ttl"`
	Rdata       string `json:"rdata"`
	View        string `json:"view"`
	Flags       uint32 `json:"flags"`
	Note        string `json:"note"`
}

type Rrset struct {
	Name string `json:"name"`
	Type string `json:"type"`
	View string `json:"view"`
	Zone string `json:"zone"`
}

const (
	RrMaxTtl = 2147483647
	RrMinTtl = 10
)

var RrTypes = []string{"soa", "ns", "a", "aw", "aaaa", "aaaaw", "mx", "cname", "cnamew", "dname", "txt", "spf", "ptr", "srv", "xw", "lname", "caa", "trans", "aa", "aapf", "transaa"}

func GetRrByZone(zone string) []Rr {
	u := api.GetRRManagerUrl()
	q := u.Query()
	q.Set("resource_type", "rr")
	q.Set("zone", zone)
	u.RawQuery = q.Encode()
	s := u.String()
	resp, err := http.Get(s)
	exitIfError(err)
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	exitIfError(err)
	var rrs []Rr
	err = json.Unmarshal(data, &rrs)
	exitIfError(err)
	return rrs
}

func CreateRrInZone(zone string, name string, typ string, ttl int, rdata string) []Rr {
	u, d := api.RRManagerRequest()
	d.ResourceType = "rr"
	var fullName string = ""
	if name == "@" {
		fullName = zone
	} else {
		fullName = fmt.Sprintf("%s.%s", name, zone)
	}
	rr := Rr{
		Zone:  zone,
		Name:  fullName,
		Type:  typ,
		Ttl:   ttl,
		Rdata: rdata,
		Flags: 1,
		View:  "others",
	}
	d.Attrs = []Rr{rr}
	data, err := json.Marshal(d)
	exitIfError(err)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(data))
	exitIfError(err)
	c := http.Client{}
	resp, err := c.Do(req)
	exitIfError(err)
	data, err = io.ReadAll(resp.Body)
	resp.Body.Close()
	exitIfError(err)
	var rrs []Rr
	err = json.Unmarshal(data, &rrs)
	if err != nil {
		fmt.Printf("%s\n", data)
	}
	exitIfError(err)
	return rrs
}

func DeleteRr(ids ...string) []Rr {
	u, d := api.RRManagerRequest()
	d.ResourceType = "rr"
	attrs := make([]map[string]string, len(ids))
	for k, v := range ids {
		attrs[k] = map[string]string{"id": v}
	}
	d.Attrs = attrs
	data, err := json.Marshal(d)
	exitIfError(err)
	req, err := http.NewRequest(http.MethodDelete, u.String(), bytes.NewReader(data))
	exitIfError(err)
	c := http.Client{}
	resp, err := c.Do(req)
	exitIfError(err)
	data, err = io.ReadAll(resp.Body)
	resp.Body.Close()
	exitIfError(err)
	var rrs []Rr
	err = json.Unmarshal(data, &rrs)
	exitIfError(err)
	return rrs
}

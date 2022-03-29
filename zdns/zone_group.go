package zdns

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ZoneGroup struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Zdnsuser   string    `json:"zdnsuser"`
	CreateTime time.Time `json:"create_time"`
}

func GetZoneGroup() []ZoneGroup {
	u := api.GetRRManagerUrl()
	q := u.Query()
	q.Set("resource_type", "zone_group")
	u.RawQuery = q.Encode()
	s := u.String()
	resp, err := http.Get(s)
	exitIfError(err)
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	exitIfError(err)
	var zgs []ZoneGroup
	err = json.Unmarshal(data, &zgs)
	exitIfError(err)
	return zgs
}

func GetZoneGroupUngroupId() (string, error) {
	zgs := GetZoneGroup()
	for i := 0; i < len(zgs); i++ {
		z := zgs[i]
		if z.Name == "ungrouped" {
			return z.Id, nil
		}
	}
	return "", errors.New("Do not have Group: ungroup")
}

func CreateZoneGroup(name string) {
	u, d := api.RRManagerRequest()
	d.ResourceType = "zone_group"
	d.Attrs = ZoneGroup{
		Name: name,
	}
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
	var zg DataStruct
	err = json.Unmarshal(data, &zg)
	exitIfError(err)
	fmt.Printf("res: %v", zg)
}

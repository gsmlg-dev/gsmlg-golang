package zdns

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	ZONE_NORMAL        = 0x0001
	ZONE_PAUSE         = 0x0002
	ZONE_CONVERT       = 0x0003
	ZONE_CONVERT_AFTER = 0x0004
	ZONE_NOAUTH        = 0x0005

	SOA_RDATA_TEMPLATE = "%s %s 1 3600 3600 3600 900"
)

type Zone struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	ReverseName    string    `json:"-"`
	Zdnsuser       string    `json:"zdnsuser"`
	Content        string    `json:"content"`
	ZoneGroup      string    `json:"zone_group"`
	CreateTime     time.Time `json:"create_time"`
	Note           string    `json:"note"`
	Flags          uint32    `json:"flags"`
	PutInRecordId  string    `json:"put_in_record_id"`
	ChineseConvert bool      `json:"chinese_convert"`
}

func GetZone() []Zone {
	u := api.GetRRManagerUrl()
	q := u.Query()
	q.Set("resource_type", "zone")
	u.RawQuery = q.Encode()
	s := u.String()
	resp, err := http.Get(s)
	exitIfError(err)
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	exitIfError(err)
	var zones []Zone
	err = json.Unmarshal(data, &zones)
	exitIfError(err)
	return zones
}

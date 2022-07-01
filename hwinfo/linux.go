package hwinfo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const PRODUCT_UUID = "/sys/class/dmi/id/product_uuid"
const BOARD_SERIAL = "/sys/class/dmi/id/board_serial"

func LinuxGetInfoString() string {
	uuid, err1 := ioutil.ReadFile(PRODUCT_UUID)
	serial, err2 := ioutil.ReadFile(BOARD_SERIAL)

	if err1 != nil && err2 != nil {
		msg := fmt.Sprintf("Get device information failure. Must Provide read privilidge to %s or %s.", PRODUCT_UUID, BOARD_SERIAL)
		errorExit(msg)
	}
	if err1 != nil {
		return strings.TrimSpace(fmt.Sprintf("%s", serial))
	}
	if err2 != nil {
		return strings.TrimSpace(fmt.Sprintf("%s", uuid))
	}

	b := append(uuid, serial...)
	return strings.TrimSpace(fmt.Sprintf("%s", b))
}

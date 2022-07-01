package hwinfo

import (
	"crypto/sha256"
	"encoding/hex"
	"runtime"

	"github.com/gsmlg-dev/gsmlg-golang/errorhandler"
)

var errorExit = errorhandler.CreateExitIfError("hwinfo")

func GetInfoString() string {
	var info string
	switch runtime.GOOS {
	case "linux":
		info = LinuxGetInfoString()
	case "darwin":
		info = DarwinGetInfoString()
	default:
		errorExit("OS not support")
	}
	return info
}

func GetCode() string {
	info := GetInfoString()
	code := sha256.Sum256([]byte(info))
	return hex.EncodeToString(code[:])
}

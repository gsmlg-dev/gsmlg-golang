package hwinfo

import "os/exec"

func DarwinGetInfoString() string {
	out, err := exec.Command("system_profiler", "SPHardwareDataType").Output()
	if err != nil {
		errorExit(err)
	}
	return string(out)
}

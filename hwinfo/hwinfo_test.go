package hwinfo_test

import (
	"testing"

	"github.com/gsmlg-dev/gsmlg-golang/hwinfo"
)

func TestService(t *testing.T) {
	t.Logf("Print Machine info: %s", hwinfo.GetInfoString())
	t.Logf("Print Machine Code: %s", hwinfo.GetCode())
}

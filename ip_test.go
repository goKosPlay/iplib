package iplib_test

import (
	"testing"
	"github.com/cc123123/iplib"
)

func TestGetRealIpAddress(t *testing.T) {
	k := iplib.getVersion()
	if k != "0.1" {
		t.Fatal("version is wrong")
	}
}

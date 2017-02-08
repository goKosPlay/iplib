package iplib_test

import (
	"testing"
	"github.com/cc123123/iplib"
)

func TestGetRealIpAddress(t *testing.T) {
	k := iplib.GetRealIpAddress()
	if k.City != "Taipei" {
		t.Fatal("city is fail")
	}

}

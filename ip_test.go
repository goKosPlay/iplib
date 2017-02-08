package iplib_test

import (
	"testing"
	"../iplib"
)

func TestGetRealIpAddress(t *testing.T) {
	k := iplib.GetRealIpAddress()
	if k.City != "Taipei" {
		t.Fatal("city is fail")
	}

}

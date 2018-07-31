package main

import (
	"github.com/goKosPlay/iplib"
	"fmt"
)
func main()  {
	k := iplib.NewIpMod()
	fmt.Println(k.GetIp()) // your remote ip
	fmt.Println(k.GetIpDetail()) // your remote ip detail
}

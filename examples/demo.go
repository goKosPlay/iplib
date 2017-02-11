package main

//author Kos
import (
	 "github.com/cc123123/iplib"
	"fmt"
)
func main() {
	fmt.Println(iplib.GetRealIpAddress())
	fmt.Println(iplib.GetIpInformation("0.0.0.0"))
}
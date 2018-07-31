# iplib [![Release](https://img.shields.io/github/release/cc123123/iplib.svg)](https://github.com/cc123123/iplib/releases) [![Build Status](https://travis-ci.org/cc123123/iplib.svg?branch=master)](https://travis-ci.org/cc123123/iplib)

簡易取得實際IP


## Install

```
go get github.com/goKosPlay/iplib
```

## Example

```
package main

//author Kos
import (
	 "github.com/goKosPlay/iplib"
	 "fmt"
)
func main() {
	k := iplab.NewIpMod()
	fmt.Println(k.GetIp())
	fmt.Println(k.GetIpDetail()["country_name"])
	fmt.Println(k.GetSelectIpDetail("0.0.0.0"))
}
```

## License

MIT.

# iplib [![Release](https://img.shields.io/github/release/goKosPlay/iplib.svg)](https://github.com/goKosPlay/iplib/releases) [![Build Status](https://travis-ci.org/goKosPlay/iplib.svg?branch=master)](https://travis-ci.org/goKosPlay/iplib)

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

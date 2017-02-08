# iplib

## Overview [![GoDoc](https://godoc.org/github.com/cc123123/iplib?status.svg)](https://godoc.org/github.com/cc123123/iplib)

簡易取得實際IP


## Install

```
go get github.com/cc123123/iplib
```

## Example

```
package main

//author Kos
import (
	 "github.com/cc123123/iplib"
	"fmt"
)
func main() {
	fmt.Println(iplib.GetRealIpAddress())
	iplib.GetLocalIAddress()
}
```

## License

MIT.

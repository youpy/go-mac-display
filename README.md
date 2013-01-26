# go-mac-display

A Go library to control displays for Mac

## Installation

```
go get github.com/youpy/go-mac-display
```

## Synopsis

```go
package main

import (
	"github.com/youpy/go-mac-display"
	"fmt"
)

func main() {
	d := display.MainDisplay()
	d.SetBrightness(1.0)
	fmt.Printf("brightness: %f\n", d.Brightness());
}
```

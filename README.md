gorduino
========

Simple Arduino wrapper for Go, based on go-firmata by kraman.

Usage:
```go
package main

import (
  "github.com/yanzay/gorduino"
  "time"
)

func main() {
  g := gorduino.NewGorduino("/dev/tty.usbmodem1411", 13)
  for {
    g.On(13)
    time.Sleep(1 * time.Second)
    g.Off(13)
    time.Sleep(1 * time.Second)
  }
}
```

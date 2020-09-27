package main

import (
	"github.com/Microsoft/go-winio"
)

func main() {
	winio.ListenPipe(`\\.\pipe\testpipe`, nil)
}

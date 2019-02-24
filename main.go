package main

import (
	"github.com/dereknex/ip-changed/cmd"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	resp := cmd.Execute(os.Args[1:])
	if resp.Err != nil {
		os.Exit(-1)
	}
}

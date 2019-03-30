package main

import (
	"os"
	"runtime"

	"github.com/dereknex/ddns/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := cmd.Execute(os.Args[1:])
	if err != nil {
		os.Exit(-1)
	}
}

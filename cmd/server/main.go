package main

import (
	"github.com/dereknex/ip-changed/pkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func runServer() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", pkg.Ping)
	}
	r.Run(":8080")
}

func main() {

}

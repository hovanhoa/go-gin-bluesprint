package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", getHello)
	_ = r.Run(":80")
}

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

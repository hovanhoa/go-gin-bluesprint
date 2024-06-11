package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/hovanhoa/go-gin-bluesprint/configs"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs: ", err)
	}

	gin.SetMode(cfg.Mode)

	r := gin.New()
	r.Use(gin.Logger())
	r.SetTrustedProxies(nil)

	r.GET("/", getHello)

	_ = r.Run(cfg.ServerAddress)
}

func getHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

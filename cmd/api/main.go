package main

import (
	"github.com/gin-gonic/gin"
	config "github.com/hovanhoa/go-gin-bluesprint/configs"
	"github.com/hovanhoa/go-gin-bluesprint/handlers"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs: ", err)
	}

	gin.SetMode(cfg.Mode)

	e := handlers.Gin(&cfg)
	handlers.SetDefault(e)

	_ = e.Run(cfg.ServerAddress)
}

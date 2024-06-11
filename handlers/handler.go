package handlers

import (
	"github.com/gin-gonic/gin"
	config "github.com/hovanhoa/go-gin-bluesprint/configs"
)

func SetDefault(e *gin.Engine) {
	e.GET("/health", HealthCheck)
}

func Gin(config *config.Config) *gin.Engine {
	e := gin.New()
	_ = e.SetTrustedProxies(nil)

	e.Use(gin.Logger())

	return e
}

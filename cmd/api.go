package main

import (
	"everdale-wiki/pkg/buildings"
	"everdale-wiki/pkg/config"
	"everdale-wiki/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	router.HEAD("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "")
	})

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/buildings", buildings.GetBuildings)
	}

	err := router.Run(config.CfgInstance.Api.Endpoint)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
}

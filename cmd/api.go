package main

import (
	"everdale-wiki/pkg/buildings"
	"everdale-wiki/pkg/config"
	"everdale-wiki/pkg/logger"
	"everdale-wiki/pkg/page"
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

	router.Static("static", "static")
	router.StaticFile("/favicon.ico", "static/100.ico")

	router.GET("/", page.Page)
	router.GET("/buildings", page.Page)

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

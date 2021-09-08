package main

import (
	"everdale-wiki/pkg/buildings"
	"everdale-wiki/pkg/challenges"
	"everdale-wiki/pkg/config"
	"everdale-wiki/pkg/logger"
	"everdale-wiki/pkg/nation_buildings"
	"everdale-wiki/pkg/page"
	"everdale-wiki/pkg/specialties"
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
	router.GET("/nation_buildings", nation_buildings.NationBuildingsPage)
	router.GET("/specialties", specialties.SpecialtiesPage)
	router.GET("/challenges", challenges.ChallengesPage)

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

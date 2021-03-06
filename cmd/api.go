package main

import (
	"everdale-wiki/pkg/buildings"
	"everdale-wiki/pkg/challenges"
	"everdale-wiki/pkg/config"
	"everdale-wiki/pkg/logger"
	"everdale-wiki/pkg/nation_buildings"
	"everdale-wiki/pkg/page"
	"everdale-wiki/pkg/recipes"
	"everdale-wiki/pkg/specialties"
	"everdale-wiki/pkg/tools"
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
	router.StaticFile("/google561dd461b94fd509.html", "static/google561dd461b94fd509.html")

	router.GET("/", page.Page)
	router.GET("/buildings", page.Page)
	router.GET("/nation_buildings", nation_buildings.NationBuildingsPage)
	router.GET("/specialties", specialties.SpecialtiesPage)
	router.GET("/challenges", challenges.ChallengesPage)
	router.GET("/tools", tools.ToolsPage)
	router.GET("/recipes", recipes.RecipesPage)

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

package server

import (
	"github.com/Tomdango/retrotool-go/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	teamsGroup := router.Group("teams")
	{
		teams := controllers.TeamsController{}
		teamsGroup.POST("/create", teams.CreateTeam)
	}

	return router
}

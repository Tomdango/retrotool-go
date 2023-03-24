package server

import (
	"github.com/Tomdango/retrotool-go/controllers"
	"github.com/Tomdango/retrotool-go/db"
	"github.com/Tomdango/retrotool-go/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(cfg aws.Config, connection *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cognitoParams := utils.GetCognitoParameters(cfg)

	authGroup := router.Group("auth")
	{
		cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)
		auth := &controllers.AuthController{
			CognitoClient: cognitoClient,
			CognitoParams: cognitoParams,
		}
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/confirmation", auth.RegisterConfirmation)
	}

	teamsGroup := router.Group("teams")
	{
		teams := &controllers.TeamsController{
			TeamsRepository: &db.TeamsRepository{DB: connection},
			CognitoParams:   &cognitoParams,
		}
		teamsGroup.GET("/", teams.GetTeams)
		teamsGroup.GET("/:teamID", teams.GetTeamByID)
		teamsGroup.POST("/create", teams.CreateTeam)
	}

	return router
}

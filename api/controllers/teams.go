package controllers

import (
	"fmt"

	"github.com/Tomdango/retrotool-go/db"
	"github.com/Tomdango/retrotool-go/utils"
	"github.com/gin-gonic/gin"
)

type TeamsController struct {
	TeamsRepository *db.TeamsRepository
	CognitoParams   *utils.CognitoParameters
}

/**
 * Create Team
 * POST /teams/create
 */
type CreateTeamRequestBody struct {
	TeamName string `json:"name" binding:"required"`
}

type CreateTeamResponseBody struct {
	Message string   `json:"message"`
	Team    *db.Team `json:"team"`
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}

func (t *TeamsController) CreateTeam(ctx *gin.Context) {
	var body CreateTeamRequestBody
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSONP(500, &ErrorResponseBody{
			Message: fmt.Sprintf("Invalid Request Body, %v", err),
		})
		return
	}

	token, err := utils.ParseJWTFromHeaders(ctx.Request.Header, t.CognitoParams.UserPoolID)
	if err != nil {
		ctx.JSONP(401, &ErrorResponseBody{
			Message: fmt.Sprintf("Invalid Access Token, %v", err),
		})
		return
	}

	newTeam := &db.Team{
		Name:  body.TeamName,
		Owner: token.Subject(),
	}
	err = t.TeamsRepository.CreateTeam(newTeam)
	if err != nil {
		ctx.JSONP(500, &ErrorResponseBody{
			Message: fmt.Sprintf("Failed to create team, %v", err),
		})
		return
	}

	ctx.JSONP(200, &CreateTeamResponseBody{
		Message: "Successfully Created New Team",
		Team:    newTeam,
	})
}

/**
 * GetTeamByID Handler
 * GET /teams/:teamID
 */
type GetTeamByIDResponseBody struct {
	Message string   `json:"message"`
	Team    *db.Team `json:"team"`
}

func (t *TeamsController) GetTeamByID(ctx *gin.Context) {
	teamID := ctx.Param("teamID")

	token, err := utils.ParseJWTFromHeaders(ctx.Request.Header, t.CognitoParams.UserPoolID)
	if err != nil {
		ctx.JSONP(401, &ErrorResponseBody{
			Message: fmt.Sprintf("Invalid Access Token, %v", err),
		})
		return
	}

	ownerID := token.Subject()
	team := t.TeamsRepository.GetTeamByID(teamID)

	if team == nil || team.Owner != ownerID {
		ctx.JSONP(400, &ErrorResponseBody{
			Message: "Team Not Found",
		})
		return
	}

	ctx.JSONP(200, &GetTeamByIDResponseBody{
		Message: "Successfully found team",
		Team:    team,
	})
}

/**
 * GetTeams
 * GET /teams/
 */
type GetTeamsResponseBody struct {
	Message string     `json:"message"`
	Teams   *[]db.Team `json:"teams"`
}

func (t *TeamsController) GetTeams(ctx *gin.Context) {
	token, err := utils.ParseJWTFromHeaders(ctx.Request.Header, t.CognitoParams.UserPoolID)
	if err != nil {
		ctx.JSONP(401, &ErrorResponseBody{
			Message: fmt.Sprintf("Invalid Access Token, %v", err),
		})
		return
	}

	ownerID := token.Subject()
	teams := t.TeamsRepository.GetAllByOwnerID(ownerID)

	ctx.JSONP(200, &GetTeamsResponseBody{
		Message: "Retrieved all teams for current user",
		Teams:   teams,
	})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamsController struct{}

func (h TeamsController) CreateTeam(c *gin.Context) {
	c.String(http.StatusOK, "Hello There!")
}

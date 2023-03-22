package db

import "github.com/aws/aws-sdk-go-v2/aws/session"

type TeamsRepository struct {
	table string
}

func (t TeamsRepository) Init(session *session.Session) {
	session
}

func (t TeamsRepository) GetTeamByID(teamID string) () {
}

// func (t TeamsRepository)

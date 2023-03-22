package models

type Member struct {
	MemberID string `json:"id"`
}

type Team struct {
	TeamID  string   `json:"id"`
	Name    string   `json:"name"`
	Members []Member `json:"members"`
}

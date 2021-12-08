package models

type Suggestion struct {
	Priority   int    `json:"priority"`
	Username   string `json:"username"`
	Suggestion int    `json:"suggestion"`
}

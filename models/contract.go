package models

type Contract struct {
	Id       int    `json:"id"`
	Clientid int    `json:"clientid"`
	Points   int    `json:"points"`
	Finished int    `json:"finished"`
	Comment  string `json:"comment"`
}

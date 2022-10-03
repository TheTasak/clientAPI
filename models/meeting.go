package models

type Meeting struct {
	Id       int    `json:"id"`
	Clientid int    `json:"clientid"`
	Type     int    `json:"type"`
	Date     string `json:"date"`
	Comment  string `json:"comment"`
}

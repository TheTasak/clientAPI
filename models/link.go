package models

type Link struct {
	Id            int    `json:"id"`
	Idrecommended int    `json:"idrecommended"`
	Idsource      int    `json:"idsource"`
	Comment       string `json:"comment"`
}

package models

type Skins struct {
	ID     int    `json:"id"`
	Name   string `json:"nama"`
	Hero   string `json:"hero"`
	Tag    string `json:"tag"`
	Config byte   `json:"config"`
}

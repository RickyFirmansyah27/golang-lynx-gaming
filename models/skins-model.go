package models

type Skins struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"nama" db:"nama"`
	Hero       string `json:"hero" db:"hero"`
	Tag        string `json:"tag" db:"tag"`
	Desciption string `json:"desc" db:"deskripsi"`
	ImageUrl   string `json:"image_url" db:"image_url"`
	Config     string `json:"config" db:"config"`
}

type Arenas struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"nama" db:"nama"`
	Desciption string `json:"desc" db:"deskripsi"`
	Tag        string `json:"tag" db:"tag"`
	ImageUrl   string `json:"image_url" db:"image_url"`
	Config     string `json:"config" db:"config"`
}

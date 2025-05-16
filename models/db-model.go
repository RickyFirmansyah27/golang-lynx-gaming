package models

type Skins struct {
	ID          int     `json:"id" db:"id"`
	Name        *string `json:"nama" db:"nama"`
	Hero        *string `json:"hero" db:"hero"`
	Tag         *string `json:"tag" db:"tag"`
	Description *string `json:"desc" db:"deskripsi"`
	ImageUrl    *string `json:"image_url" db:"image_url"`
	Config      *string `json:"config" db:"config"`
}

type Arenas struct {
	ID          int     `json:"id" db:"id"`
	Name        *string `json:"nama" db:"nama"`
	Description *string `json:"desc" db:"deskripsi"`
	Tag         *string `json:"tag" db:"tag"`
	ImageUrl    *string `json:"image_url" db:"image_url"`
	Config      *string `json:"config" db:"config"`
}

type User struct {
	ID       uint   `json:"id" db:"id"`
	GameID   string `json:"gameId" db:"gameId"`
	ServerID string `json:"serverId" db:"serverId"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

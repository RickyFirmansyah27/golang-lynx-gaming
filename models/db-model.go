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
	Nickname string `json:"nickname" db:"nickname"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type AccountRequest struct {
	TypeName string `json:"type_name"`
	UserID   string `json:"userId"`
	ZoneID   string `json:"zoneId"`
}

type AccountResponse struct {
	Message    string `json:"message"`
	Nickname   string `json:"nickname"`
	ServerTime string `json:"server_time"`
	Status     bool   `json:"status"`
	TypeName   string `json:"type_name"`
}

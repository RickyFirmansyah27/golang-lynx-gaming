package models

type Item struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	Stock      int    `json:"stock"`
	Unit       string `json:"unit"`
	MinStock   int    `json:"min_stock"`
}

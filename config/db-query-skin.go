package config

import (
	"strconv"
	"strings"

	"go-fiber-vercel/models" // Import the models package
)

func GetAllskins(params map[string]string) ([]models.Skins, int, error) {
	page, err := strconv.Atoi(params["page"])
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(params["size"])
	if err != nil {
		size = 10
	}

	if size != 10 && size != 20 && size != 50 {
		size = 10
	}

	query := "SELECT id, nama, tag, hero, image_url, config FROM lynx.skins"
	queryCount := "SELECT COUNT(*) FROM lynx.skins"

	whereClauses := []string{}
	queryParams := []interface{}{}
	paramIndex := 1

	if nama, ok := params["nama"]; ok && nama != "" {
		whereClauses = append(whereClauses, "nama ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+nama+"%")
		paramIndex++
	}

	if categoryID, ok := params["tag"]; ok && categoryID != "" {
		whereClauses = append(whereClauses, "tag = $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, categoryID)
		paramIndex++
	}

	if hero, ok := params["hero"]; ok && hero != "" {
		whereClauses = append(whereClauses, "hero = $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, hero)
		paramIndex++
	}

	if len(whereClauses) > 0 {
		whereStr := strings.Join(whereClauses, " AND ")
		query += " WHERE " + whereStr
		queryCount += " WHERE " + whereStr
	}

	sortBy := params["sort_by"]
	sortOrder := params["sort_order"]

	allowedSortFields := map[string]bool{
		"id": true, "nama": true, "tag": true, "hero": true,
	}

	if _, ok := allowedSortFields[sortBy]; !ok {
		sortBy = "id"
	}

	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "ASC"
	}

	query += " ORDER BY " + sortBy + " " + sortOrder

	offset := (page - 1) * size
	query += " LIMIT $" + strconv.Itoa(paramIndex) + " OFFSET $" + strconv.Itoa(paramIndex+1)
	queryParams = append(queryParams, size, offset)

	countRows, err := ExecuteSQLWithParams(queryCount, queryParams[:paramIndex-1]...)
	if err != nil {
		return nil, 0, err
	}
	defer countRows.Close()

	var totalData int
	if countRows.Next() {
		if err := countRows.Scan(&totalData); err != nil {
			return nil, 0, err
		}
	}

	rows, err := ExecuteSQLWithParams(query, queryParams...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	skins := []models.Skins{}
	for rows.Next() {
		var skin models.Skins
		if err := rows.Scan(&skin.ID, &skin.Name, &skin.Tag, &skin.Hero, &skin.ImageUrl, &skin.Config); err != nil {
			return nil, 0, err
		}
		skins = append(skins, skin)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return skins, totalData, nil
}
func UpdateSkin(id int, skin models.Skins) (models.Skins, error) {
	// Update the skin
	query := `UPDATE lynx.skins 
              SET nama = $1, tag = $2, hero = $3, image_url = $4, config = $5 
              WHERE id = $6 
              RETURNING id, nama, tag, hero, image_url, config`

	row, err := ExecuteSQLWithParams(query, skin.Name, skin.Tag, skin.Hero, skin.ImageUrl, skin.Config, id)
	if err != nil {
		return models.Skins{}, err
	}
	defer row.Close()
	var updatedSkin models.Skins
	if row.Next() {
		if err := row.Scan(&updatedSkin.ID, &updatedSkin.Name, &updatedSkin.Tag, &updatedSkin.Hero, &updatedSkin.ImageUrl, &updatedSkin.Config); err != nil {
			return models.Skins{}, err
		}
	} else {
		return models.Skins{}, nil
	}

	return updatedSkin, nil
}

func CreateSkin(skin models.Skins) (models.Skins, error) {
	query := `INSERT INTO lynx.skins (nama, tag, hero, image_url, config)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, nama, tag, hero, image_url, config`

	row, err := ExecuteSQLWithParams(query, skin.Name, skin.Tag, skin.Hero, skin.ImageUrl, skin.Config)
	if err != nil {
		return models.Skins{}, err
	}
	defer row.Close()

	var newSkin models.Skins
	if row.Next() {
		if err := row.Scan(&newSkin.ID, &newSkin.Name, &newSkin.Tag, &newSkin.Hero, &newSkin.ImageUrl, &newSkin.Config); err != nil {
			return models.Skins{}, err
		}
	}

	return newSkin, nil
}

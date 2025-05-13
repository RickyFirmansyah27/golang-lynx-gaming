package config

import (
	"strconv"
	"strings"

	"go-fiber-vercel/models"
)

func GetAllArenas(params map[string]string) ([]models.Arenas, int, error) {
	page, err := strconv.Atoi(params["page"])
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(params["size"])
	if err != nil {
		size = 1
	}

	query := "SELECT id, nama, deskripsi, tag, image_url, config FROM lynx.arenas"
	queryCount := "SELECT COUNT(*) FROM lynx.arenas"

	whereClauses := []string{}
	queryParams := []interface{}{}
	paramIndex := 1

	if nama, ok := params["nama"]; ok && nama != "" {
		whereClauses = append(whereClauses, "nama ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+nama+"%")
		paramIndex++
	}

	if tag, ok := params["tag"]; ok && tag != "" {
		whereClauses = append(whereClauses, "tag = $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, tag)
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
		"id": true, "nama": true, "tag": true,
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

	arenas := []models.Arenas{}
	for rows.Next() {
		var arena models.Arenas
		if err := rows.Scan(&arena.ID, &arena.Name, &arena.Desciption, &arena.Tag, &arena.ImageUrl, &arena.Config); err != nil {
			return nil, 0, err
		}
		arenas = append(arenas, arena)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return arenas, totalData, nil
}

func CreateArena(arena models.Arenas) (models.Arenas, error) {
	query := `INSERT INTO lynx.arenas (nama, deskripsi, tag, image_url, config)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, nama, deskripsi, tag, image_url, config`

	row, err := ExecuteSQLWithParams(query, arena.Name, arena.Desciption, arena.Tag, arena.ImageUrl, arena.Config)
	if err != nil {
		return models.Arenas{}, err
	}
	defer row.Close()

	var newArena models.Arenas
	if row.Next() {
		if err := row.Scan(&newArena.ID, &newArena.Name, &newArena.Desciption, &newArena.Tag, &newArena.ImageUrl, &newArena.Config); err != nil {
			return models.Arenas{}, err
		}
	}

	return newArena, nil
}

func UpdateArena(id int, arena models.Arenas) (models.Arenas, error) {
	query := `UPDATE lynx.arenas 
              SET nama = $1, deskripsi = $2, tag = $3, image_url = $4, config = $5
              WHERE id = $6
              RETURNING id, nama, deskripsi, tag, image_url, config`

	row, err := ExecuteSQLWithParams(query, arena.Name, arena.Desciption, arena.Tag, arena.ImageUrl, arena.Config, id)
	if err != nil {
		return models.Arenas{}, err
	}
	defer row.Close()

	var updatedArena models.Arenas
	if row.Next() {
		if err := row.Scan(&updatedArena.ID, &updatedArena.Name, &updatedArena.Desciption, &updatedArena.Tag, &updatedArena.ImageUrl, &updatedArena.Config); err != nil {
			return models.Arenas{}, err
		}
	} else {
		return models.Arenas{}, nil
	}

	return updatedArena, nil
}

package config

import (
	"fmt"
	"go-fiber-vercel/models"
	"strconv"
	"strings"
)

func GetAllItems(params map[string]string) ([]map[string]interface{}, int, error) {
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

	query := "SELECT * FROM items"
	queryCount := "SELECT COUNT(*) FROM items"

	whereClauses := []string{}
	queryParams := []interface{}{}
	paramIndex := 1

	if name, ok := params["name"]; ok && name != "" {
		whereClauses = append(whereClauses, "name ILIKE $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, "%"+name+"%")
		paramIndex++
	}

	if categoryID, ok := params["category_id"]; ok && categoryID != "" {
		whereClauses = append(whereClauses, "category_id = $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, categoryID)
		paramIndex++
	}

	if stock, ok := params["stock"]; ok && stock != "" {
		whereClauses = append(whereClauses, "stock = $"+strconv.Itoa(paramIndex))
		queryParams = append(queryParams, stock)
		paramIndex++
	}

	if len(whereClauses) > 0 {
		whereStr := strings.Join(whereClauses, " AND ")
		query += " WHERE " + whereStr
		queryCount += " WHERE " + whereStr
	}

	sortBy := params["sort_by"]
	sortOrder := params["sort_order"]

	// Update allowed sort fields
	allowedSortFields := map[string]bool{
		"id": true, "name": true, "category_id": true,
		"stock": true, "unit": true, "min_stock": true,
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

	items := []map[string]interface{}{}
	for rows.Next() {
		var id, categoryID, stock, minStock int
		var name, unit string

		if err := rows.Scan(&id, &name, &categoryID, &stock, &unit, &minStock); err != nil {
			return nil, 0, err
		}

		items = append(items, map[string]interface{}{
			"id":          id,
			"name":        name,
			"category_id": categoryID,
			"stock":       stock,
			"unit":        unit,
			"min_stock":   minStock,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return items, totalData, nil
}

func CreateItem(item *models.Item) (*models.Item, error) {
	sql := `
        INSERT INTO items (name, category_id, stock, unit, min_stock)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, name, category_id, stock, unit, min_stock`

	row, err := ExecuteSQLWithParams(sql,
		item.Name,
		item.CategoryID,
		item.Stock,
		item.Unit,
		item.MinStock,
	)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, fmt.Errorf("no rows returned after insert")
	}

	createdItem := &models.Item{}
	err = row.Scan(
		&createdItem.ID,
		&createdItem.Name,
		&createdItem.CategoryID,
		&createdItem.Stock,
		&createdItem.Unit,
		&createdItem.MinStock,
	)
	if err != nil {
		return nil, err
	}

	return createdItem, nil
}

func UpdateItem(id int, item *models.Item) (*models.Item, error) {
	sql := `
        UPDATE items 
        SET name = $1, category_id = $2, stock = $3, unit = $4, min_stock = $5
        WHERE id = $6
        RETURNING id, name, category_id, stock, unit, min_stock`

	row, err := ExecuteSQLWithParams(sql,
		item.Name,
		item.CategoryID,
		item.Stock,
		item.Unit,
		item.MinStock,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if !row.Next() {
		return nil, fmt.Errorf("item with id %d not found", id)
	}

	updatedItem := &models.Item{}
	err = row.Scan(
		&updatedItem.ID,
		&updatedItem.Name,
		&updatedItem.CategoryID,
		&updatedItem.Stock,
		&updatedItem.Unit,
		&updatedItem.MinStock,
	)
	if err != nil {
		return nil, err
	}

	return updatedItem, nil
}

func DeleteItem(id int) error {
	sql := "DELETE FROM items WHERE id = $1"
	rows, err := ExecuteSQLWithParams(sql, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

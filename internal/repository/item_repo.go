package repository

import (
	"database/sql"
	"project-app-inventaris-cli-iay/internal/entity"
	"time"
)

type ItemRepository struct {
	DB *sql.DB
}

func (r *ItemRepository) GetAll() ([]entity.Item, error) {
	query := `
		SELECT i.id, i.name, c.name, i.price, i.purchase_date 
		FROM items i 
		LEFT JOIN categories c ON i.category_id = c.id
	`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entity.Item
	for rows.Next() {
		var i entity.Item
		var catName sql.NullString // Handle null category
		err := rows.Scan(&i.ID, &i.Name, &catName, &i.Price, &i.PurchaseDate)
		if err != nil {
			return nil, err
		}
		i.CategoryName = catName.String

		// Hitung hari penggunaan
		days := time.Since(i.PurchaseDate).Hours() / 24
		i.DaysUsed = int(days)

		items = append(items, i)
	}
	return items, nil
}

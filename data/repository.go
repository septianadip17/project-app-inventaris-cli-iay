package data

import (
	"database/sql"
	"project-app-inventaris-cli-iay/model"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// === KATEGORI ===

func (r *Repository) AddCategory(c model.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, c.Name, c.Description)
	return err
}

func (r *Repository) GetAllCategories() ([]model.Category, error) {
	rows, err := r.DB.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// === BARANG (ITEMS) ===
// (Nanti tambahkan fungsi AddItem, GetAllItems di sini)

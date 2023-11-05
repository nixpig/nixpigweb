package queries

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/models"
)

type CategoryQueries struct {
	*sql.DB
}

func (q *CategoryQueries) GetCategories() ([]models.Category, error) {
	query := "select id, name_ from category_"

	var categories []models.Category

	rows, err := q.Query(query)
	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		category := models.Category{}

		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

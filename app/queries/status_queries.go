package queries

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/models"
)

type StatusQueries struct {
	*sql.DB
}

func (q *StatusQueries) GetStatuses() ([]models.Status, error) {
	query := "select id, name_ from status_"

	var statuses []models.Status

	rows, err := q.Query(query)
	if err != nil {
		return statuses, err
	}

	defer rows.Close()

	for rows.Next() {
		status := models.Status{}

		if err := rows.Scan(&status.Id, &status.Name); err != nil {
			return statuses, err
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}

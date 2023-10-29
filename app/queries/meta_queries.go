package queries

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/models"
)

type MetaQueries struct {
	*sql.DB
}

func (q *MetaQueries) GetMeta() ([]models.Meta, error) {
	meta := []models.Meta{}

	query := "select id, name_, value_ from meta_"

	rows, err := q.Query(query)
	if err != nil {
		return meta, err
	}

	defer rows.Close()

	for rows.Next() {
		metaItem := models.Meta{}

		if err := rows.Scan(&metaItem.Id, &metaItem.Name, &metaItem.Value); err != nil {
			return meta, err
		}

		meta = append(meta, metaItem)
	}

	return meta, nil
}

func (q *MetaQueries) GetMetaById(id int) (models.Meta, error) {
	meta := models.Meta{}

	query := "select id, name_, value_ from meta_ where id = $1"

	row := q.QueryRow(query, id)

	if err := row.Scan(&meta.Id, &meta.Name, &meta.Value); err != nil {
		return meta, err
	}

	return meta, nil
}

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

func (q *MetaQueries) CreateMeta(meta models.Meta) error {
	query := "insert into meta_ (name_, value_) values ($1, $2)"

	_, err := q.Exec(query, meta.Name, meta.Value)
	if err != nil {
		return err
	}

	return nil
}

func (q *MetaQueries) DeleteMeta(id int) error {
	query := "delete from meta_ where id = $1"

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *MetaQueries) UpdateMeta(meta models.Meta) error {
	query := "update meta_ set name_ = $2, value_ = $3 where id = $1"

	_, err := q.Exec(query, meta.Id, meta.Name, meta.Value)
	if err != nil {
		return err
	}

	return nil
}

package queries

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/api/models"
)

type ConfigQueries struct {
	db *sql.DB
}

func (q *ConfigQueries) GetConfigs() ([]models.Config, error) {
	configs := []models.Config{}

	query := "select * from config_"

	rows, err := q.db.Query(query)
	if err != nil {
		return configs, err
	}

	defer rows.Close()

	for rows.Next() {
		config := models.Config{}

		if err := rows.Scan(&config.Name, &config.Value, &config.Id); err != nil {
			return configs, err
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func (q *ConfigQueries) GetConfig(id int) (models.Config, error) {
	config := models.Config{}

	query := "select * from config_ where id = $1"

	row := q.db.QueryRow(query, id)

	if err := row.Scan(&config.Id, &config.Name, &config.Value); err != nil {
		return config, err
	}

	return config, nil
}

func (q *ConfigQueries) CreateConfig(config *models.Config) error {
	query := "insert into config_ (name_, value_) values ($1, $2)"

	_, err := q.db.Exec(query, &config.Name, &config.Value)
	if err != nil {
		return err
	}

	return nil
}

func (q *ConfigQueries) DeleteConfig(id int) error {
	query := "delete from config_ where_ id = $1"

	_, err := q.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (q *ConfigQueries) UpdateConfig(config *models.Config) error {
	query := "update config_ set name_ = $2, value_ = $3 where id = $1"

	_, err := q.db.Exec(query, &config.Name, &config.Value, &config.Id)
	if err != nil {
		return err
	}

	return nil
}

package queries

import (
	"database/sql"
	"fmt"

	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

type Content struct {
	*sql.DB
}

func (c *Content) CreateContent(content *models.Content) (int64, error) {
	query := `insert into content_ (title_, subtitle_, slug_, body_, type_) values ($1, $2, $3, $4, $5)`

	res, err := c.Exec(query, &content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.Type)
	if err != nil {
		return 0, fmt.Errorf("error inserting new content item\n%v", err)
	}

	id, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Content) GetContent() ([]models.Content, error) {
	var contents []models.Content

	query := `select id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_ from content_`

	rows, err := c.Query(query)
	if err != nil {
		return contents, fmt.Errorf("error fetching content from database\n%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var content models.Content
		if err := rows.Scan(&content.Id, &content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.CreatedAt, &content.UpdatedAt, &content.Type); err != nil {
			return contents, fmt.Errorf("error scanning data to struct\n%v", err)
		}

		contents = append(contents, content)
	}

	return contents, nil
}

func (c *Content) GetContentById(id int) (models.Content, error) {
	query := `select id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_ from content_ where id_ = $1`

	var content models.Content

	row := c.QueryRow(query, id)
	if err := row.Scan(&content.Id, &content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.CreatedAt, &content.UpdatedAt, &content.Type); err != nil {
		return content, fmt.Errorf("error scanning data\n%v", err)
	}

	return content, nil
}

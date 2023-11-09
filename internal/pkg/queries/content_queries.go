package queries

import (
	"database/sql"
	"fmt"

	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

type Content struct {
	*sql.DB
}

func (c *Content) GetContent() ([]models.Content, error) {
	var contents []models.Content

	query := `select id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_ from nixpigweb_`

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

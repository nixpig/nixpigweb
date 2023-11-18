package queries

import (
	"fmt"

	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
)

func CreateContent(content *models.Content) (int64, error) {
	query := `insert into content_ (title_, subtitle_, slug_, body_, type_, user_id_) values ($1, $2, $3, $4, $5, $6)`

	res, err := database.DB.Exec(query, &content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.Type, &content.UserId)
	if err != nil {
		return 0, fmt.Errorf("error inserting new content item\n%v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func GetContent() ([]models.Content, error) {
	var contents []models.Content

	query := `select id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_, user_id_ from content_`

	rows, err := database.DB.Query(query)
	if err != nil {
		return contents, fmt.Errorf("error fetching content from database\n%v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var content models.Content
		if err := rows.Scan(&content.Id, &content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.CreatedAt, &content.UpdatedAt, &content.Type, &content.UserId); err != nil {
			return contents, fmt.Errorf("error scanning data to struct\n%v", err)
		}

		contents = append(contents, content)
	}

	return contents, nil
}

func GetContentById(id int) (models.Content, error) {
	query := `select id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_, user_id_ from content_ where id_ = $1`

	var content models.Content

	row := database.DB.QueryRow(query, id)
	if err := row.Scan(&content.Id, &content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.CreatedAt, &content.UpdatedAt, &content.Type, &content.UserId); err != nil {
		return content, fmt.Errorf("error scanning data\n%v", err)
	}

	return content, nil
}

func DeleteContentById(id int) (int64, error) {
	query := `delete from content_ where id_ = $1`

	res, err := database.DB.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("error deleting record with id: %v\n%v", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func UpdateContent(content *models.Content) (int64, error) {
	query := `update content_ set title_ = $2, subtitle_ = $3, slug_ = $4, body_ = $5, updated_at_ = $6, type_ = $7 where id_ = $1`

	res, err := database.DB.Exec(query, &content.Id, &content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.UpdatedAt, &content.Type)
	if err != nil {
		return 0, fmt.Errorf("error updating record\n%v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

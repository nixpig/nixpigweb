package queries_test

import (
	"regexp"
	"testing"

	_ "database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/nixpig/nixpigweb/internal/pkg/database"
	"github.com/nixpig/nixpigweb/internal/pkg/models"
	"github.com/nixpig/nixpigweb/internal/pkg/queries"
	"github.com/stretchr/testify/assert"
)

func TestCreateContent(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error while creating database mock: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectedQuery := regexp.QuoteMeta(`insert into content_ (title_, subtitle_, slug_, body_, type_, user_id_) values ($1, $2, $3, $4, $5, $6)`)

	mockResult := sqlmock.NewResult(1, 1)

	content := models.Content{
		Title:    "Lorem ipsum",
		Subtitle: "Dolar sit amet",
		Slug:     "lorem-ipsum",
		Body:     "Lorem ipsum dolar sit amet",
		Type:     "post",
		UserId:   23,
	}

	mock.ExpectExec(expectedQuery).WithArgs(&content.Title, &content.Subtitle, &content.Slug, &content.Body, &content.Type, &content.UserId).WillReturnResult(mockResult)

	affectedRows, err := queries.CreateContent(&content)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), affectedRows)
}

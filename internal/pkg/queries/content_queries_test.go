package queries_test

import (
	"regexp"
	"testing"
	"time"

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

func TestGetContent(t *testing.T) {
	var content []models.Content
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error trying to create mock database: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectedQuery := regexp.QuoteMeta(`select id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_, user_id_ from content_`)

	mockTime := time.Now()

	mockResult := sqlmock.
		NewRows([]string{"id_", "title_", "subtitle_", "slug_", "body_", "created_at_", "updated_at_", "type_", "user_id"}).
		AddRow(1, "title one", "subtitle one", "slug-one", "body one", mockTime, mockTime, "post", 23).
		AddRow(1, "title two", "subtitle two", "slug-two", "body two", mockTime, mockTime, "page", 23)

	mock.ExpectQuery(expectedQuery).WillReturnRows(mockResult)

	expectedContent := []models.Content{
		{
			Id:        1,
			Title:     "title one",
			Subtitle:  "subtitle one",
			Slug:      "slug-one",
			Body:      "body one",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			Type:      "post",
			UserId:    23,
		},
		{
			Id:        1,
			Title:     "title two",
			Subtitle:  "subtitle two",
			Slug:      "slug-two",
			Body:      "body two",
			CreatedAt: mockTime,
			UpdatedAt: mockTime,
			Type:      "page",
			UserId:    23,
		},
	}

	content, err = queries.GetContent()

	assert.Nil(t, err)
	assert.Equal(t, expectedContent, content)
}

func TestGetContentById(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectQuery := regexp.QuoteMeta(`select id_, title_, subtitle_, slug_, body_, created_at_, updated_at_, type_, user_id_ from content_ where id_ = $1`)

	mockTime := time.Now()

	mockResult := sqlmock.
		NewRows([]string{"id_", "title_", "subtitle_", "slug_", "body_", "created_at_", "updated_at_", "type_", "user_id"}).
		AddRow(1, "title one", "subtitle one", "slug-one", "body one", mockTime, mockTime, "post", 23)

	mock.ExpectQuery(expectQuery).WithArgs(1).WillReturnRows(mockResult)

	expectedResult := models.Content{
		Id:        1,
		Title:     "title one",
		Subtitle:  "subtitle one",
		Slug:      "slug-one",
		Body:      "body one",
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		Type:      "post",
		UserId:    23,
	}

	content, err := queries.GetContentById(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, content)
}

func TestDeleteContentById(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectExec := regexp.QuoteMeta(`delete from content_ where id_ = $1`)

	mock.ExpectExec(expectExec).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	affectedRows, err := queries.DeleteContentById(1)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), affectedRows)
}

func TestUpdateContent(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db instance: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectExec := regexp.QuoteMeta(`update content_ set title_ = $2, subtitle_ = $3, slug_ = $4, body_ = $5, updated_at_ = $6, type_ = $7 where id_ = $1`)

	mockTime := time.Now()

	mock.ExpectExec(expectExec).
		WithArgs(1, "title changed", "subtitle changed", "slug-changed", "body changed", mockTime, "post").
		WillReturnResult(sqlmock.NewResult(1, 1))

	contentUpdate := models.Content{
		Id:        1,
		Title:     "title changed",
		Subtitle:  "subtitle changed",
		Slug:      "slug-changed",
		Body:      "body changed",
		UpdatedAt: mockTime,
		Type:      "post",
	}

	rowsAffected, err := queries.UpdateContent(&contentUpdate)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)

}

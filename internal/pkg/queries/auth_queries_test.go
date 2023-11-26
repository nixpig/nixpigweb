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

func TestGetUserByUserName(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create database mock: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectQuery := regexp.QuoteMeta(`select id_, username_, email_, password_, is_admin_ from users_ where username_ = $1`)

	userRows := sqlmock.NewRows([]string{"id_", "username_", "email_", "password_", "is_admin_"}).AddRow(23, "np1", "2@email.com", "p4ssw0rd", true)

	mock.ExpectQuery(expectQuery).WithArgs("np1").WillReturnRows(userRows)

	user, err := queries.GetUserByUsername("np1")

	assert.Nil(t, err)
	assert.Equal(t, models.User{
		Id:       23,
		Username: "np1",
		Email:    "2@email.com",
		Password: "p4ssw0rd",
		IsAdmin:  true,
	}, user)

}

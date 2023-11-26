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

func TestCreateUser(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error creating database mock: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectedQuery := regexp.QuoteMeta(`insert into users_ (username_, email_, password_) values ($1, $2, $3)`)

	mockResult := sqlmock.NewResult(1, 1)

	user := models.User{
		Username: "test_username",
		Password: "test_password",
		Email:    "test@example.com",
	}

	mock.ExpectExec(expectedQuery).WithArgs(user.Username, user.Email, user.Password).WillReturnResult(mockResult)

	rowsAffected, err := queries.CreateUser(&user)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), rowsAffected)
}

func TestGetUserById(t *testing.T) {
	var user models.User
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error creating database mock: %s", err)
	}

	defer db.Close()

	database.DB = db

	expectedQuery := regexp.QuoteMeta(`select id_, username_, email_, is_admin_ from users_ where id_ = $1`)

	userRows := sqlmock.NewRows([]string{"id_", "username_", "email_", "is_admin_"}).AddRow(23, "np1", "2@email.com", true)

	id := 23

	mock.ExpectQuery(expectedQuery).WithArgs(id).WillReturnRows(userRows)

	user, err = queries.GetUserById(id)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user, models.User{
		Id:       23,
		Username: "np1",
		Email:    "2@email.com",
		IsAdmin:  true,
	})

	mock.ExpectQuery(expectedQuery).WithArgs(50).WillReturnRows(sqlmock.NewRows([]string{}))

	user, err = queries.GetUserById(50)
	assert.NotNil(t, err)
	assert.Empty(t, user)

}

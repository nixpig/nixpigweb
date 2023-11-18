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

func TestGetUserById(t *testing.T) {
	var user models.User
	var err error

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' opening stub connection", err)
	}

	defer db.Close()

	database.DB = db

	query := regexp.QuoteMeta(`select id_, username_, email_, is_admin_ from users_ where id_ = $1`)

	userRows := sqlmock.NewRows([]string{"id_", "username_", "email_", "is_admin_"}).AddRow(23, "np1", "2@email.com", true)

	id := 23

	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(userRows)

	user, err = queries.GetUserById(id)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user, models.User{
		Id:       23,
		Username: "np1",
		Email:    "2@email.com",
		IsAdmin:  true,
	})

	mock.ExpectQuery(query).WithArgs(50).WillReturnRows(sqlmock.NewRows([]string{}))

	user, err = queries.GetUserById(50)
	assert.NotNil(t, err)
	assert.Empty(t, user)

}

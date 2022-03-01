package repository

import (
	"gochicoba/models"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockRepo() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	dbGorm, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return dbGorm, mock
}

func TestGetAllUsersRepo(t *testing.T) {
	db, mock := NewMockRepo()
	repo := &userRepository{db}

	var user = &models.User{
		ID:   1,
		Name: "test",
	}

	query := `SELECT * FROM "users"`

	rows := sqlmock.NewRows([]string{"id", "nama"}).
		AddRow(user.ID, user.Name)

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	results, err := repo.GetAllUsers(models.UserFilter{})

	assert.NotEmpty(t, results)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
}

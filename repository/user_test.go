package repository

import (
	"fmt"
	"gochicoba/db"
	"gochicoba/models"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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

	var users = []models.User{{
		ID:   1,
		Name: "First",
	}, {
		ID:   2,
		Name: "Second",
	}}

	query := `SELECT * FROM "users"`

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(users[0].ID, users[0].Name).AddRow(users[1].ID, users[1].Name)

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	results, err := repo.GetAllUsers(models.UserFilter{})

	assert.NotEmpty(t, results)
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestGetUserRepo(t *testing.T) {
	db, mock := NewMockRepo()
	repo := &userRepository{db}

	var user models.User = models.User{
		ID:        1,
		Name:      "First",
		Age:       19,
		Status:    "First",
		CreatedAt: time.Time{},
	}

	query := `SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`

	rows := sqlmock.NewRows([]string{"id", "name", "age", "status", "created_at"}).
		AddRow(user.ID, user.Name, user.Age, user.Status, user.CreatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(user.ID).WillReturnRows(rows)

	result, err := repo.GetUser(1)

	assert.NotEmpty(t, result)
	assert.NoError(t, err)
	assert.Equal(t, &user, result)
}

func TestAddUserRepo(t *testing.T) {
	db, mock := NewMockRepo()
	repo := &userRepository{db}

	var user models.User = models.User{
		ID:        1,
		Name:      "First",
		Age:       19,
		Status:    "First",
		CreatedAt: time.Time{},
	}

	// query := `INSERT INTO "users" ("name","age","status") VALUES ($1,$2,$3)`
	// query := `INSERT INTO "users" (.+) RETURNING`

	// rows := sqlmock.NewRows([]string{"id", "name", "age", "status", "created_at"}).
	// 	AddRow(user.ID, user.Name, user.Age, user.Status, user.CreatedAt)

	// mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(sqlmock.AnyArg(), user.Name, user.Age, user.Status).WillReturnRows(rows)

	mock.ExpectQuery(`INSERT INTO "users" (.+) RETURNING`).WithArgs(sqlmock.AnyArg(), user.Name, user.Age, user.Status).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "status", "created_at"}).AddRow(user.ID, user.Name, user.Age, user.Status, user.CreatedAt))

	result, err := repo.AddUser(&models.User{
		Name:   "First",
		Age:    19,
		Status: "First",
	})

	fmt.Println(result)
	fmt.Println("Hallo")
	fmt.Println(err)

	// assert.NotEmpty(t, result)
	// assert.NoError(t, err)
	// assert.Equal(t, &user, result)
}

func TestUpdateUserRepo(t *testing.T) {
	db, mock := NewMockRepo()
	repo := &userRepository{db}

	// mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(sqlmock.AnyArg(), user.Name, user.Age, user.Status).WillReturnRows(rows)

	fmt.Println(mock)
	fmt.Println(repo)
}

func TestDeleteUserRepo(t *testing.T) {
	db, mock := NewMockRepo()
	repo := &userRepository{db}

	query := `DELETE FROM "users" WHERE id = $1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnError(nil)

	err := repo.DeleteUser(1)

	fmt.Println("Hallo")
	fmt.Println(err)

	assert.Error(t, err) // ada error
}

/////////////////////////////////////////////
/////////// REAL DATABASE SECTION ///////////
/////////////////////////////////////////////

func TestDeleteUserRepoReal(t *testing.T) {
	db := db.DatabaseInitialize()
	repo := &userRepository{db: db}

	err := repo.DeleteUser(1000)

	assert.Error(t, err, "Harusnya Error")
}

func TestCreateDeleteRepoReal(t *testing.T) {
	db := db.DatabaseInitialize()
	repo := &userRepository{db: db}

	var userData = &models.User{
		Name:   "First",
		Age:    100,
		Status: "First",
	}

	res, err := repo.AddUser(userData)

	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, userData.Name, res.Name, "Name must be same")

	err = repo.DeleteUser(res.ID)

	assert.Nil(t, err, "Error must be nil")
}

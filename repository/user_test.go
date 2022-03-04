package repository

import (
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

func TestUserRepository_GetAllUsers(t *testing.T) {
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
	assert.Equal(t, *results[0], users[0])
	assert.Equal(t, results[1], &users[1])
}

func TestUserRespository_GetUser(t *testing.T) {
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
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Status, result.Status)
}

// func TestAddUserRepo(t *testing.T) {
// 	db, mock := NewMockRepo()
// 	repo := &userRepository{db}

// 	var user models.User = models.User{
// 		ID:        1,
// 		Name:      "First",
// 		Age:       19,
// 		Status:    "First",
// 		CreatedAt: time.Time{},
// 	}

// 	mock.ExpectQuery(`INSERT INTO "users" (.+) RETURNING`).WithArgs(sqlmock.AnyArg(), user.Name, user.Age, user.Status).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "status", "created_at"}).AddRow(user.ID, user.Name, user.Age, user.Status, user.CreatedAt))

// 	result, err := repo.AddUser(&models.User{
// 		Name:   "First",
// 		Age:    19,
// 		Status: "First",
// 	})

// 	fmt.Println(result)
// 	fmt.Println("Hallo")
// 	fmt.Println(err)

// 	assert.NotEmpty(t, result)
// 	assert.NoError(t, err)
// 	assert.Equal(t, &user, result)
// }

// func TestUpdateUserRepo(t *testing.T) {
// 	db, mock := NewMockRepo()
// 	repo := &userRepository{db}

// 	// mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(sqlmock.AnyArg(), user.Name, user.Age, user.Status).WillReturnRows(rows)

// 	fmt.Println(mock)
// 	fmt.Println(repo)
// }

// func TestDeleteUserRepo(t *testing.T) {
// 	db, mock := NewMockRepo()
// 	repo := &userRepository{db}

// 	query := `DELETE FROM "users" WHERE id = $1`
// 	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnError(nil)

// 	err := repo.DeleteUser(1)

// 	fmt.Println("Hallo")
// 	fmt.Println(err)

// 	assert.Error(t, err)
// }

/////////////////////////////////////////////
/////////// REAL DATABASE SECTION ///////////
/////////////////////////////////////////////

func TestUserRepositoryReal_DeleteUser(t *testing.T) {
	db := db.DatabaseInitialize()
	repo := &userRepository{db: db}

	err := repo.DeleteUser(1000)

	assert.Error(t, err, "Harusnya Error")
}

func TestUserRepositoryReal_CreateDeleteUser(t *testing.T) {
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

func TestUserRepositoryReal_GetAllUsers_FilterNameNoExist(t *testing.T) {
	db := db.DatabaseInitialize()
	repo := &userRepository{db: db}

	res, err := repo.GetAllUsers(models.UserFilter{
		Name:    "Zahra",
		AgeUp:   0,
		AgeDown: 0,
		Status:  "",
	})

	// fmt.Println(res)
	assert.NotNil(t, res, "must be not nil")
	assert.Nil(t, err, "must be nil")
	assert.Equal(t, 0, len(res), "data not found")

	res, err = repo.GetAllUsers(models.UserFilter{
		Name:    "",
		AgeUp:   0,
		AgeDown: 0,
		Status:  "",
	})

	assert.NotNil(t, res)
	assert.Nil(t, err)

}

func TestUserRepositoryReal_GetAllUsers_FilterAge(t *testing.T) {
	db := db.DatabaseInitialize()
	repo := &userRepository{db: db}

	res, err := repo.GetAllUsers(models.UserFilter{
		Name:    "",
		AgeUp:   30,
		AgeDown: 10,
		Status:  "",
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)

	for i := range res {
		var actual = res[i].Age >= 10 && res[i].Age <= 30
		assert.Equal(t, true, actual)
	}
}

func TestUserRepositoryReal_GetAllUsers_FilterStatus(t *testing.T) {
	db := db.DatabaseInitialize()
	repo := &userRepository{db: db}

	res, err := repo.GetAllUsers(models.UserFilter{
		Name:    "",
		AgeUp:   0,
		AgeDown: 0,
		Status:  "First",
	})

	assert.Nil(t, err)
	assert.NotNil(t, res)

	for i := range res {
		assert.Equal(t, "First", res[i].Status)
	}

}

func TestUserRepositoryReal_UpdateUser(t *testing.T) {
	db := db.DatabaseInitialize()
	repo := &userRepository{db: db}

	var userData = &models.User{
		ID:        1,
		Name:      "First",
		Age:       100,
		Status:    "First",
		CreatedAt: time.Now(),
	}

	res, err := repo.UpdateUser(userData.ID, userData)

	assert.Nil(t, err, "error must nil")
	assert.NotNil(t, res, "res must not nil")
	assert.Equal(t, userData.ID, res.ID)
	assert.Equal(t, userData.Name, res.Name)
	assert.Equal(t, userData.Age, res.Age)
	assert.Equal(t, userData.Status, res.Status)
	assert.Equal(t, userData.CreatedAt, res.CreatedAt)
}

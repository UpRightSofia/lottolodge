package user_store

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models/utils"
	"github.com/google/uuid"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestUserPsqlStore(t *testing.T) {

	utils.WithPostgres(t, func(db *sql.DB, wg *sync.WaitGroup) {
		store := UserPostgresStore{db}

		utils.WithParallel(wg, func() {
			t.Run("GetUser returns User", func(t *testing.T) {
				uuId := uuid.New().String()

				_, err := db.Exec(`insert into users (id, balance_e5) values ($1, $2);`, uuId, 100)
				if err != nil {
					t.Error(err)
				}

				user, getErr := store.GetUser(uuId)
				if getErr != nil {
					t.Errorf("GetUser failed: %s\n", getErr)
				}

				expectedUser := User{ID: uuId, BalanceE5: 100}

				compareUsers(t, expectedUser, user)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("GetUser returns error if User not found", func(t *testing.T) {
				_, getErr := store.GetUser("not-found")
				if getErr == nil {
					t.Errorf("GetUser should've failed")
				}
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("Create user creates User", func(t *testing.T) {
				id := uuid.New().String()
				createdUser, err := store.CreateUser(CreateUserRequest{ID: id, BalanceE5: 100})
				if err != nil {
					t.Errorf("CreateUser failed: %s\n", err)
				}

				user, getErr := store.GetUser(createdUser.ID)
				if getErr != nil {
					t.Errorf("GetUser failed: %s\n", err)
				}

				compareUsers(t, createdUser, user)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("CreateUser returns error if User with same userId exists", func(t *testing.T) {
				id := uuid.New().String()
				_, err := store.CreateUser(CreateUserRequest{ID: id})
				if err != nil {
					t.Errorf("CreateUser failed: %s\n", err)
				}

				_, err = store.CreateUser(CreateUserRequest{ID: id})
				if err == nil {
					t.Errorf("CreateUser should've failed")
				}
			})
		})
	})
}

func compareUsers(t *testing.T, expected, actual User) {
	t.Helper()

	if actual.ID != expected.ID || actual.BalanceE5 != expected.BalanceE5 {
		t.Errorf("Expected %+v, got %+v\n", expected, actual)
	}
}

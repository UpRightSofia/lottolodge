package pool_store

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models/utils"
	"github.com/google/uuid"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestPoolPsqlStore(t *testing.T) {

	utils.WithPostgres(t, func(db *sql.DB, wg *sync.WaitGroup) {
		store := PoolPostgresStore{db}

		utils.WithParallel(wg, func() {
			t.Run("GetPool returns Pool", func(t *testing.T) {
				uuId := uuid.New().String()
				details := sql.NullString{String: `{"key": "value"}`, Valid: true} // Sample JSON for demonstration

				_, err := db.Exec(`insert into pools (id, details) values ($1, $2);`, uuId, details)

				if err != nil {
					t.Error(err)
				}

				pool, getErr := store.GetPool(uuId)
				if getErr != nil {
					t.Errorf("GetPool failed: %s\n", getErr)
				}

				// The UpdatedAt field will be set to the current time when the record is inserted, so we don't set it here.
				expectedPool := Pool{
					ID:       uuId,
					Details:  details,
					IsActive: true,
					// UpdatedAt: This field will be set by the database, so we can skip it in this comparison or handle it separately.
				}

				comparePools(t, expectedPool, pool)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("GetLastCompleted returns Pool", func(t *testing.T) {
				uuId1 := uuid.New().String()
				uuId2 := uuid.New().String()
				details := sql.NullString{String: `{"key": "value"}`, Valid: true} // Sample JSON for demonstration
				isActive := false

				_, err := db.Exec(`insert into pools (id, details, is_active) values ($1, $2, $3);`, uuId1, details, isActive)
				if err != nil {
					t.Error(err)
				}


				_, err = db.Exec(`insert into pools (id, details, is_active) values ($1, $2, $3);`, uuId2, details, isActive)
				if err != nil {
					t.Error(err)
				}

				pool, getErr := store.GetLastCompleted()
				if getErr != nil {
					t.Errorf("GetPool failed: %s\n", getErr)
				}

				// The UpdatedAt field will be set to the current time when the record is inserted, so we don't set it here.
				expectedPool := Pool{
					ID:       uuId2,
					Details:  details,
					IsActive: false,
					// UpdatedAt: This field will be set by the database, so we can skip it in this comparison or handle it separately.
				}

				comparePools(t, expectedPool, pool)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("Create pool creates Pool", func(t *testing.T) {
				createdPool, err := store.CreatePool()
				if err != nil {
					t.Errorf("CreatePool failed: %s\n", err)
				}

				pool, getErr := store.GetPool(createdPool.ID)
				if getErr != nil {
					t.Errorf("GetPool failed: %s\n", err)
				}

				comparePools(t, createdPool, pool)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("Mark pool as finished updates relevant fields", func(t *testing.T) {
				createdPool, err := store.CreatePool()
				if err != nil {
					t.Errorf("CreatePool failed: %s\n", err)
				}

				pool, err := store.MarkPoolAsDone(createdPool.ID, `{"key": "value"}`)
				if err != nil {
					t.Errorf("MarkPoolAsDone failed: %s\n", err)
				}

				expectedPool := Pool{
					ID:       createdPool.ID,
					Details:  sql.NullString{String: `{"key": "value"}`, Valid: true},
					IsActive: false,
				}

				comparePools(t, expectedPool, pool)
			})
		})
	})

	utils.WithPostgres(t, func(db *sql.DB, wg *sync.WaitGroup) {
		store := PoolPostgresStore{db}

		utils.WithParallel(wg, func() {
			t.Run("GetActivePool returns Pool", func(t *testing.T) {
				createdPool, err := store.CreatePool()
				if err != nil {
					t.Errorf("CreatePool failed: %s\n", err)
				}

				pool, getErr := store.GetActivePool()
				if getErr != nil {
					t.Errorf("GetTodayPool failed: %s\n", err)
				}

				comparePools(t, createdPool, pool)
			})
		})
	})
}

func comparePools(t *testing.T, expected, got Pool) {
	if expected.ID != got.ID ||
		expected.Details.Valid != got.Details.Valid ||
		expected.Details.String != got.Details.String ||
		expected.IsActive != got.IsActive {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

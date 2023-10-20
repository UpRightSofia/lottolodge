package setting_store

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models/utils"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestSettingPsqlStore(t *testing.T) {

	utils.WithPostgres(t, func(db *sql.DB, wg *sync.WaitGroup) {
		store := SettingPostgresStore{db}

		utils.WithParallel(wg, func() {
			t.Run("GetLastSetting returns Setting", func(t *testing.T) {
				ticketPrizeE5 := 1
				payoutPercent := 2
				_, err := db.Exec(`insert into settings (ticket_prize_e5, payout_percent) values ($1, $2);`, ticketPrizeE5, payoutPercent)
				if err != nil {
					t.Error(err)
				}

				ticketPrizeE5 = 3
				payoutPercent = 4
				_, err = db.Exec(`insert into settings (ticket_prize_e5, payout_percent) values ($1, $2);`, ticketPrizeE5, payoutPercent)
				if err != nil {
					t.Error(err)
				}

				ticketPrizeE5 = 5
				payoutPercent = 6
				_, err = db.Exec(`insert into settings (ticket_prize_e5, payout_percent) values ($1, $2);`, ticketPrizeE5, payoutPercent)
				if err != nil {
					t.Error(err)
				}

				setting, getErr := store.GetLastSetting()
				if getErr != nil {
					t.Errorf("GetLastSetting failed: %s\n", getErr)
				}

				expectedSetting := Setting{TicketPrizeE5: ticketPrizeE5, PayoutPercent: payoutPercent}
				compareSetting(t, expectedSetting, setting)
			})
		})

		utils.WithParallel(wg, func() {
			t.Run("Create setting creates setting", func(t *testing.T) {
				createdSetting, err := store.CreateSetting(CreateSettingRequest{TicketPrizeE5: 1, PayoutPercent: 2})
				if err != nil {
					t.Errorf("CreateSetting failed: %s\n", err)
				}

				setting, getErr := store.GetLastSetting()
				if getErr != nil {
					t.Errorf("GetLastSetting failed: %s\n", err)
				}

				compareSetting(t, createdSetting, setting)
			})
		})
	})
}

func compareSetting(t *testing.T, expected, actual Setting) {
	t.Helper()

	if actual.PayoutPercent != expected.PayoutPercent ||
		actual.TicketPrizeE5 != expected.TicketPrizeE5 {
		t.Errorf("Expected %+v, got %+v\n", expected, actual)
	}
}

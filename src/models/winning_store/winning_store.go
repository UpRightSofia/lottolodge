package winning_store

type WinningStore interface {
	GetWinning(id string) (Winning, error)
	CreateWinning(request CreateWinningRequest) (Winning, error)
	GetWinningsForUserAndPool(user_id string, pool_id string) ([]Winning, error)
}

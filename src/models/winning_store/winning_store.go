package winning_store

type WinningStore interface {
	GetWinning(id string) (Winning, error)
	CreateWinning(request CreateWinningRequest) (Winning, error)
}

package pool_store

type PoolStore interface {
	GetPool(id string) (Pool, error)
	CreatePool(request CreatePoolRequest) (Pool, error)
	GetActivePool() (Pool, error)
	MarkPoolAsDone(id string, details string) (Pool, error)
	GetLastCompleted() (Pool, error)
}

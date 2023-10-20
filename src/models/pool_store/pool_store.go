package pool_store

type PoolStore interface {
	GetPool(id string) (Pool, error)
	CreatePool(request CreatePoolRequest) (Pool, error)
}

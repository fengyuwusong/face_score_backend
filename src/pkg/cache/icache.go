package cache

type ICache interface {
	CanEliminate() bool
	Update() error
	Reload() error
	Close(bDrop bool) error
}

package components

type Locker interface {
	AcquireLock(string) Mutex
	ReleaseLock(Mutex) (bool, error)
}
type Mutex interface {
	Lock() error
	Unlock() (bool, error)
}

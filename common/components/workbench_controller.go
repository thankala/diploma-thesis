package components

type WorkbenchController interface {
	AcquireLock() error
	ReleaseLock() error
	AcquirePlaceLock(WorkbenchPlace) error
	ReleasePlaceLock() error
	Rotate() error
}

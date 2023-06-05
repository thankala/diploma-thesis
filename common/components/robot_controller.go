package components

type RobotController interface {
	AcquireLock() error
	ReleaseLock() error
	Move() error
}

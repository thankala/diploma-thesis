package components

import (
	"github.com/google/uuid"
)

type Workbench struct {
	id                  uuid.UUID
	locker              Locker
	workbenchMutex      Mutex
	workbenchPlaceMutex Mutex
}

type WorkbenchPlace int64

const (
	UNUSED WorkbenchPlace = iota
	P1
	P2
	P3
)

func (p WorkbenchPlace) String() string {
	switch p {
	case P1:
		return "P1"
	case P2:
		return "P2"
	case P3:
		return "P3"
	}
	return "unknown"
}

func NewWorkbench(id uuid.UUID, locker Locker) *Workbench {
	return &Workbench{
		id:     id,
		locker: locker,
	}
}

func (w *Workbench) AcquireLock() error {
	w.workbenchMutex = w.locker.AcquireLock(w.id.String() + "WorkbenchLock")
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := w.workbenchMutex.Lock(); err != nil {
		return err
	}
	return nil
}

func (w *Workbench) ReleaseLock() error {
	// Do your work that requires the lock.
	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := w.workbenchMutex.Unlock(); !ok || err != nil {
		return err
	}
	return nil
}

func (w *Workbench) AcquirePlaceLock(place WorkbenchPlace) error {
	w.workbenchPlaceMutex = w.locker.AcquireLock(w.id.String() + place.String() + "WorkbenchLock")
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := w.workbenchPlaceMutex.Lock(); err != nil {
		return err
	}
	return nil
}

func (w *Workbench) ReleasePlaceLock() error {
	// Do your work that requires the lock.
	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := w.workbenchPlaceMutex.Unlock(); !ok || err != nil {
		return err
	}
	return nil
}

func (w *Workbench) Rotate() error {
	return nil
	//val, err := l.redis.Get(context.TODO(), "WorkbenchPosition").Result()
	//if err != nil {
	//	return &err
	//}
	//switch val {
	//case "R3":
	//	logger.Info("Rotating head from R3 to R1", log.M{})
	//	err := l.redis.Set(context.TODO(), "WorkbenchPosition", "R1", 0).Err()
	//	if err != nil {
	//		return &err
	//	}
	//	logger.Info("Head rotated from R3 to R1", log.M{})
	//case "R1":
	//	logger.Info("Head at R1, no need to rotate", log.M{})
	//	return nil
	//case "R2":
	//	logger.Info("Rotating head from R2 to R1", log.M{})
	//	err := l.redis.Set(context.TODO(), "WorkbenchPosition", "R1", 0).Err()
	//	if err != nil {
	//		return &err
	//	}
	//	logger.Info("Head rotated from R2 to R1", log.M{})
	//}
	//return nil
}

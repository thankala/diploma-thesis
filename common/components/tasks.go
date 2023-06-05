package components

type Task int64

const (
	AT1 Task = iota
	AT2
	AT3
	AT4
	AT5
	AT6
	AT7
	AT8
)

func (t Task) String() string {
	switch t {
	case AT1:
		return "AT1"
	case AT2:
		return "AT2"
	case AT3:
		return "AT3"
	case AT4:
		return "AT4"
	case AT5:
		return "AT5"
	case AT6:
		return "AT6"
	case AT7:
		return "AT7"
	case AT8:
		return "AT7"
	}
	return "unknown"
}

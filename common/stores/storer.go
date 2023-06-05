package stores

type Storer interface {
	Store(string, []byte) error
	Load(string) ([]byte, error)
}

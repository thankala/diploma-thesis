package middlewares

type Persister interface {
	State() ([]byte, error)
	LoadState(string) error
}

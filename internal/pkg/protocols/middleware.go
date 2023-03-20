package protocols

type Middleware interface {
	Handle(*HttpRequest) error
}

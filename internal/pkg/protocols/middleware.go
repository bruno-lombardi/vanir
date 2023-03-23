package protocols

type Middleware interface {
	Handle(req *HttpRequest) error
}

package protocols

type HttpRequest struct {
	Body        interface{}
	PathParams  map[string]string
	QueryParams map[string][]string
	Headers     interface{}
}

type HttpResponse struct {
	StatusCode int
	Body       interface{}
	Headers    map[string][]string
}

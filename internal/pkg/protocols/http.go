package protocols

import (
	"net/http"
	"vanir/internal/pkg/data/models"
)

type HttpRequest struct {
	Body              interface{}
	PathParams        map[string]string
	QueryParams       map[string][]string
	HttpReq           *http.Request
	AuthenticatedUser *models.User
}

type HttpResponse struct {
	StatusCode int
	Body       interface{}
	Headers    map[string][]string
}

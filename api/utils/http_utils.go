package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var TraceHeaders = [...]string{
	"x-request-id",
}

type Request struct {
	Transport http.RoundTripper
	Ctx       *gin.Context
}

func (t *Request) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, d := range TraceHeaders {
		if value, ok := t.Ctx.Get(d); ok {
			req.Header.Add(d, value.(string))
		}
	}
	response, err := t.Transport.RoundTrip(req)
	return response, err
}

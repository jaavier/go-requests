package gorequests

import (
	"net/http"
	"net/url"
)

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Url      string
	Method   string
	Headers  []Header
	BodyJSON interface{}
	Form     url.Values
}

type Response struct {
	Body          string
	StatusCode    int
	Headers       http.Header
	Cookies       []*http.Cookie
	ContentLength int64
}

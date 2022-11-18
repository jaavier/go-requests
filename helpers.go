package gorequests

import (
	"io"
	"net/http"
	"net/url"
)

func isForm(form url.Values) bool {
	for key := range form {
		if len(key) > 0 {
			return true
		}
	}
	return false
}

func buildResponse(response *http.Response) Response {
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{
			StatusCode: 500,
		}
	}
	return Response{
		Body:          string(data),
		StatusCode:    response.StatusCode,
		Headers:       response.Header,
		Cookies:       response.Cookies(),
		ContentLength: response.ContentLength,
	}
}

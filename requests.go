package gorequests

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendRequest(req Request) (Response, error) {
	if isForm(req.Form) {
		httpResponse, err := http.PostForm(req.Url, req.Form)
		if err != nil {
			return Response{
				StatusCode: 400,
			}, err
		}
		return buildResponse(httpResponse), nil
	}

	var structAsBytes, err = json.Marshal(req.BodyJSON)
	if err != nil {
		return Response{
			StatusCode: 400,
		}, err
	}

	var structAsBuffer = bytes.NewBuffer(structAsBytes)
	httpRequest, err := http.NewRequest(req.Method, req.Url, structAsBuffer)
	if err != nil {
		return Response{
			StatusCode: 500,
		}, err
	}

	for _, header := range req.Headers {
		httpRequest.Header.Add(header.Key, header.Value)
	}

	var client http.Client
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return Response{
			StatusCode: 500,
		}, err
	}

	return buildResponse(httpResponse), nil
}

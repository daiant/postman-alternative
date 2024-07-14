package http

import (
	"bytes"
	"net/http"

	http_types "jerry.io/core/http/types"
)

func MakeRequest(req http_types.Request) (*http.Response, error) {
	client := http.Client{}
	request, err := http.NewRequest(req.Method, req.Url, bytes.NewReader([]byte(req.Body)))
	// request, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return nil, err
	}
	setHeaders(request, req.Headers)
	return client.Do(request)
}
func setHeaders(request *http.Request, headers []http_types.Header) {
	for _, header := range headers {
		request.Header.Add(header.Name, header.Value)
	}
}

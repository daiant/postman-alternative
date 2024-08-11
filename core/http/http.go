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
	setParams(request, req.Params)
	return client.Do(request)
}
func setParams(request *http.Request, params [][]string) {
	q := request.URL.Query()
	for i := range params[0] {
		q.Add(params[0][i], params[1][i])
	}
	request.URL.RawQuery = q.Encode()
}
func setHeaders(request *http.Request, headers [][]string) {
	if len(headers) <= 0 {
		return
	}
	for i := range headers[0] {
		request.Header.Add(headers[0][i], headers[1][i])
	}
}

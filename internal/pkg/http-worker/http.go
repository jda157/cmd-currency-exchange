package http_worker

import (
	"io/ioutil"
	"net/http"
)

type HttpWorker struct {
}

func NewHttpWorker() *HttpWorker {
	return &HttpWorker{
	}
}

func (pw *HttpWorker) GetResponse(request string) ([]byte, error) {
	resp, err := http.Get(request)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

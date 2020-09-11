package http_worker

import (
	"context"
)

type HttpWorker struct {
	apiPath string
}

// NewHTTPPostWorker returns new instance of http post worker
func NewHttpWorker(api string) *HttpWorker {
	return &HttpWorker{
		apiPath: api,
	}
}

// SetContentType inject content type to worker
func (pw *HttpWorker) SetContentType(contentType string) {
	//pw.contentType = contentType
}

// GetResponse returns response after request
func (pw *HttpWorker) GetResponse(ctx context.Context) []byte {
	//url := pw.path + pw.request
	//resp, err := http.Post(url, pw.contentType, bytes.NewBuffer(pw.bodyPost))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//return body
}

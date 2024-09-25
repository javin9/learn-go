package real

import (
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	StatusCode int
}

func (m *Retriever) Get(url string) (htmlContent string) {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}

	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	return string(dumpResponse)
}

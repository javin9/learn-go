package infra

import (
	"fmt"
	"io"
	"net/http"
)

// Retrieve 当成一个类了
type Retrieve struct {
}

func (me *Retrieve) Get(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("%s", "失败了")
	}

	defer response.Body.Close()
	byteList, _ := io.ReadAll(response.Body)
	return string(byteList), nil
}

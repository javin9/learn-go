package realgo

import (
	"fmt"
	"io"
	"net/http"
)

type Retriever struct {
	Header string
}

func (me *Retriever) String() string {
	return fmt.Sprintln("Retrieve toString")
}

func (me *Retriever) Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("%s id error", url)
	}

	defer resp.Body.Close()
	byteList, ioErr := io.ReadAll(resp.Body)
	if ioErr != nil {
		return "", fmt.Errorf("%s id error", url)
	}

	return string(byteList), nil
}

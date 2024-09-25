package infra

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Retriever struct{}

// 是没有指明实现了接口
func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	defer resp.Body.Close()

	bytesContent, _ := ioutil.ReadAll(resp.Body)
	return string(bytesContent)
}

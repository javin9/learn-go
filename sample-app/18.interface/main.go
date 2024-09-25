package main

import (
	"fmt"
	"io"
	"net/http"
)

func RetrieveGet(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("%s", "失败了")
	}

	defer response.Body.Close()
	byteList, _ := io.ReadAll(response.Body)
	return string(byteList), nil
}

func main() {
	fmt.Println("interface demo")
	response, err := http.Get("https://www.xueersi.com")
	if err != nil {
		return
	}
	fmt.Println(response.Status, response.Body)

	defer response.Body.Close()

	byteList, _ := io.ReadAll(response.Body)
	//  []byte,可以直接转字符串
	fmt.Printf("%s \n", byteList)
	// resString := fmt.Sprintf("%s", byteList)
	// 将 fmt.Sprintf("%s", data) 替换为 string(data)，其中 data 是 []byte。
	// 性能上也会更为优越，因为避免了不必要的格式化函数调用
	resString := string(byteList)

	fmt.Println("---------")
	fmt.Println(resString)

	// 提取成公共函数
	res, _ := RetrieveGet("https://www.baidu.com")
	fmt.Println(res)
}

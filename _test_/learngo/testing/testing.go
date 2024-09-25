package testing

type Retriever struct{}

//不用省略
func (Retriever) Get(url string) string {
	return "ok"
}

package mock

type Retriever struct {
}

func (me *Retriever) Get(url string) (string, error) {
	return "mock data", nil
}

func (me *Retriever) Post(url string) string {
	return "mock data"
}

package mock

type Retriever struct {
	Content string
}

func (r Retriever) Get(url string) (htmlContent string) {
	//TODO implement me
	return r.Content
}

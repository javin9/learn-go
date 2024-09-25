package tsetse

type Retrieve struct {
}

// 实现接口的时候，这里不能用 *Retrieve
func (me Retrieve) Get(url string) (string, error) {
	return "test memeber", nil
}

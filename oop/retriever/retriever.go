package retriever

type Retriever interface {
	Get(url string) string
}

func Download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}



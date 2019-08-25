package retriever

type Retriever interface {
	Get(url string) string
}

func Download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func Post(p Poster) string {
	return p.Post("http://www.imooc.cocm", map[string]string{
		"Name":   "snowdreams1006",
		"Course": "golang",
	})
}

type RetrieverAndPoster interface {
	Retriever
	Poster
}

func Session(rp RetrieverAndPoster){
	//rp.Get()
	//rp.Post()
}
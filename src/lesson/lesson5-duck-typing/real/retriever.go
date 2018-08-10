package real

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":   "suhanyu",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
	Connect(config map[string]string) interface{}
}

// todo
func myfuncPostRetriever(_this RetrieverPoster) string {
	return ""
}

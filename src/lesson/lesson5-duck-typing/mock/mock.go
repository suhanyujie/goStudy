package mock

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	Timtout time.Duration
}

// todo
func (_this Retriever) Post() string {
	return "this is post content...."
}

func (r Retriever) Get (url string) string {
	res,err := http.Get(url)
	if err != nil{
		panic(err)
	}
	result,err := httputil.DumpResponse(res,true)
	if err != nil{
		panic(err)
	}
	res.Body.Close()
	return string(result)
}

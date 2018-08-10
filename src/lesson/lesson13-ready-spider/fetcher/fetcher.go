package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"bufio"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"log"
)

func Fetch(url string) ([]byte,error) {
	res, err := http.Get(url);
	if err != nil {
		return nil,err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		//fmt.Printf("Error:status code (%d)is not 200\n", res.StatusCode);
		return nil,fmt.Errorf("Error:fetch is error:%v", res.StatusCode)
	}
	bodyReader := bufio.NewReader(res.Body)
	e := determineDecoder(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// 通过读取字节，判断字节的编码
func determineDecoder(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:determineDecoder func has error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
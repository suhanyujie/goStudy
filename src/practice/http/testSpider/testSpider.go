package testSpider

import (
	"net/http"
	"errors"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
)

func GetHttpResponse(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(err.Error() + "-------[30015]")
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Max OS X 10_13_6) AppleWebKit/537.36 (KHTML,Like Gecko)Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New(err.Error() + "-------[30021]")
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.New(fmt.Sprintf("改请求的状态码为：%d\n", response.StatusCode))
	}
	utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())

	return ioutil.ReadAll(utf8Content)
}

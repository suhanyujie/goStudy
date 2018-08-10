package helpers

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"bytes"
	"io/ioutil"
)

type ShortUrl struct {
	Url string
}

const token = "AIzaSyDbmWuecTNuNU3MV2ItGbVdKOssVKTPpp8"
const apiUrl = "https://www.googleapis.com/urlshortener/v1/url"

// todo
func (_this *ShortUrl) GetToken() string {
	return token
}

// todo
func (_this *ShortUrl) Do(longUrl string) {
	apiUrl := apiUrl + "?key=" + token
	param := map[string]string{
		"longUrl": longUrl,
	}
	jsonResult, err := json.Marshal(param)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(apiUrl)
	resp, err := http.Post(apiUrl, "application/json", bytes.NewReader(jsonResult))
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
}

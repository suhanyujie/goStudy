package parser

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"bytes"
	"strings"
	"os"
	url2 "net/url"
	"net/http/httputil"
)

type requestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condition"`
}

func requestBodyDecoder(request *http.Request) *json.Decoder {
	body,err := ioutil.ReadAll(request.Body)
	if err!=nil {
		log.Printf("Error reading body:%v\n", err)
		panic(err)
	}
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

func parseRequestBody(request *http.Request) requestPayloadStruct {
	decoder := requestBodyDecoder(request)
	var requestPayload requestPayloadStruct
	err := decoder.Decode(&requestPayload)
	if err!=nil {
		panic(err)
	}
	return requestPayload
}

func getProxyUrl(proxyConditionRaw string) string {
	proxyCondition := strings.ToUpper(proxyConditionRaw)
	aUrl := os.Getenv("A_CONDITION_URL")
	bUrl := os.Getenv("B_CONDITION_URL")
	defaultUrl := os.Getenv("DEFAULT_CONDITION_URL")
	if proxyCondition == "A" {
		return aUrl
	}
	if proxyCondition == "B" {
		return bUrl
	}
	return defaultUrl
}

func logRequestPayload(requestPayload requestPayloadStruct, proxyUrl string) {
	log.Printf("proxy_condition:%s, proxy_url:%s\n", requestPayload.ProxyCondition, proxyUrl)
}

func serveReverseProxy(target string,res http.ResponseWriter, req *http.Request) {
	url,_:= url2.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host",req.Header.Get("Host"))
	req.Host = url.Host
	proxy.ServeHTTP(res,req)
}

func HandleRequestAndRedirect(res http.ResponseWriter,req *http.Request) {
	requestPayload := parseRequestBody(req)
	url := getProxyUrl(requestPayload.ProxyCondition)
	logRequestPayload(requestPayload, url)
	serveReverseProxy(url, res, req)
}


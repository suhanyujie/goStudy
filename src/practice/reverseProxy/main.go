package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"practice/reverseProxy/parser"
)

//获取env中的配置
func getEnv(key, fallback string) string {
	if value,ok := os.LookupEnv(key);ok {
		return value;
	}

	return fallback;
}
//获取监听的端口
func getListenAddress() string {
	port := getEnv("PORT", "2000");

	return ":"+port;
}
//记录引入的环境变量
func logsetup() {
	aConditionUrl := os.Getenv("A_CONDITION_URL")
	bConditionUrl := os.Getenv("B_CONDITION_URL")
	defaultUrl := os.Getenv("DEFAULT_CONDITION_URL")

	log.Printf("Server will run on:%s\n", getListenAddress());
	log.Printf("redirecting to A url:%s\n", aConditionUrl)
	log.Printf("redirecting to B url:%s\n", bConditionUrl)
	log.Printf("redirecting to default url:%s\n", defaultUrl)
}

func main() {
	logsetup();
	port := getEnv("PORT", "1330");
	fmt.Println(port);

	http.HandleFunc("/", parser.HandleRequestAndRedirect)
	err := http.ListenAndServe(getListenAddress(), nil)
	if err != nil {
		panic(err)
	}
}

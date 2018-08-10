package main

import "practice/urlShort/helpers"

// 入口函数
func main() {
	shortUrl := helpers.ShortUrl{}
	longUrl := "http://www.google.com"
	shortUrl.Do(longUrl)
}

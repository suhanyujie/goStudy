package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const tpl = `
<html>
	<head>
		<title>test form</title>
	</head>
	<body>
		<form method="post" action=".">
			username:<input type="text" name="username">
			passwd：<input type="password" name="password">
			<button type="submit">登录</button>
		</form>


	</body>
</html>`

func main() {
	http.HandleFunc("/", Hey)
	http.ListenAndServe(":8080", nil)

}

func Hey(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.New("hey")
		t.Parse(tpl)
		t.Execute(rw, nil)
	} else {
		fmt.Println(r.FormValue("username"))
	}
}

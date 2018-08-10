package main

import (
	"testing"
	"net/http"
	"fmt"
)

func errWrapper() {

}

type testData struct {
	a,b string
}

func errPanic(writer http.Response, request *http.Request) error {
	panic(123)
}

func TestErrWrapperInServer(t *testing.T) {
	tests := []struct {
		h appHandler
		code int
		message string
	}{
		{errPanic,500,""},
	}
	fmt.Println(tests)
}
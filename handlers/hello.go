package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	hello.logger.Println("Hello World")

	data, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(responseWriter, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(responseWriter, "Hello, %s", data)
}

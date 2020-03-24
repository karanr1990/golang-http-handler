package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello  {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter,r *http.Request)  {
	h.l.Println("hello world")
		data, err :=ioutil.ReadAll(r.Body)
		if err != nil{
			http.Error(rw,"opps",http.StatusBadRequest)
			return
		}

		fmt.Fprint(rw,"hello world 1111",data)

}

package main

import (
	"github.com/Quest-CIO/go-micro-app/handlers"
	"log"
	"net/http"
	"os"
)

func main()  {
	l := log.New(os.Stdout,"product-api",log.LstdFlags)
	
	hh :=handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	
	sm := http.NewServeMux()
	
	sm.Handle("/",hh)
	sm.Handle("/goodbye",gh)
	
	http.ListenAndServe(":4000",sm)
}


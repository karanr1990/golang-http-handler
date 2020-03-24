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

        //s := &http.Server{
	//	Addr: ":4000",
	//	Handler: sm,
	//	IdleTimeout: 120*time.Second,
	//	ReadTimeout: 1*time.Second,
	//	WriteTimeout: 1*time.Second,
	//}
	//
	//go func() {
	//	err := s.ListenAndServe()
	//
	//	if err != nil {
	//		l.Fatal(err)
	//	}
	//
	//}()
	//
	//sigchan := make(chan os.Signal)
	//signal.Notify(sigchan,os.Interrupt)
	//signal.Notify(sigchan,os.Kill)
	//
	//sig := <- sigchan
	//
	//l.Println("recieved terminate...",sig)
	//
	//tc, _ := context.WithTimeout(context.Background(),30*time.Second)
	//
	//s.Shutdown(tc)


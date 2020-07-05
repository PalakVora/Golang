package main

import (
	"GitVsGolang/Working/handlers"
	"log"
	"net/http"
	"os"
	"time"
	"context"
	"os/signal"
	
)

func main() {
	//HandleFunc : registers a function on a path to default serve mux
	//Default serve mux is a http handler
	//handler is an interface
	//HandleFunc converts the func passed as parameter into a handler for registering to the server
	//http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	//})
	//URL is a greedy match , Anything other than the word goodbye will match the above
	//	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	//		log.Println("Goodbye World")
	//	})
	
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	//create handlers
	//hh := handlers.Disp(l)
	//gh :=handlers.NewGoodbye(l)
	ph:=handlers.NewProduct(l)
	//create a new serve mux and register the handlers
	sm:=http.NewServeMux()
	sm.Handle("/",ph)
	//sm.Handle("/goodbye",gh)

	//create a new server
	s:=&http.Server{
		Addr:			":9090",
		Handler:		sm,
		ErrorLog:		l,
		IdleTimeout: 	10*time.Second,
		ReadTimeout: 	1*time.Second,
		WriteTimeout: 	1*time.Second,
	}   //Created a manual server
	//http.ListenAndServe(":9090", sm) //Constructs http server and registers default handler
	
	//start server
	go func(){
		err:=s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//interrupt and gracefully shutdown server
	sigChan:= make(chan os.Signal)
	signal.Notify(sigChan,os.Interrupt)
	signal.Notify(sigChan,os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown",sig)

	tc,_:=context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc) //waits for the exit of all the existing clients and completion of requests without accepting any more request
}

package main

import (
	"net/http"
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

	http.ListenAndServe(":9090", nil) //Constructs http server and registers default handler

}

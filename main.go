package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//HandleFunc : registers a function on a path to default serve mux
	//Default serve mux is a http handler
	//handler is an interface
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body) //ioutil returns data and err if any ; read Nic from request body in command line using ioutil
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest) //Allows the customization of the kind of err to show to the user
			//rw.Write([]byte("Oops"))
			return //return needed cause http.error doesnt terminate go program
		}
		log.Printf("Data %s \n", d) //while executing curl make sure to use -d 'Data' ; Invoke webrequest can be removed by Remove-item alias:curl command, ensures that curl is running
		fmt.Fprintf(rw, "Hello %s \n", d)
	})
	//URL is a greedy match , Anything other than the word goodbye will match the above
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":9090", nil) //Constructs http server and registers default handler

}

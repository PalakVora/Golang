package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Handle1 struct {
	l *log.Logger
}

func Disp(l *log.Logger) *Handle1 {
	return &Handle1{l}
}
func (h *Handle1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body) //ioutil returns data and err if any ; read Nic from request body in command line using ioutil
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		//rw.WriteHeader(http.StatusBadRequest) //Allows the customization of the kind of err to show to the user
		//rw.Write([]byte("Oops"))
		fmt.Fprintf(rw, "Bold")
		return //return needed cause http.error doesnt terminate go program
	}
	//log.Printf("Data %s \n", d) //while executing curl make sure to use -d 'Data' ; Invoke webrequest can be removed by Remove-item alias:curl command, ensures that curl is running
	fmt.Fprintf(rw, "Hello %s", d)
}

package handlers
 import (
	 "log"
	 "net/http"
	 "GitVsGolang/Working/data"
	 "encoding/json"
 )

type Product struct{
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product{
	return &Product{l}
}
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	lp:= data.GetProducts()   //lp is a list, use encoding json to convert list into json for returning it to user
	d,err:=json.Marshal(lp)
	if err != nil {
		http.Error(rw,"Unable to marshal json",http.StatusInternalServerError)
	}
	rw.Write(d)
}
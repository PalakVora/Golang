package handlers
import (
	"log"
	"net/http"
)
type Goodbye struct{
	l *log.Logger
}
func NewGoodbye(l *log.Logger) *Goodbye{
	return &Goodbye{l}
}
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request){   // Response writeris an interface and Request is a struct, hence the syntax
	rw.Write([]byte("byee"))
}
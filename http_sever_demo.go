package main 

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)


func logging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		handler(w, r)
	}
}
func main() {
	r := mux.NewRouter()


	r.HandleFunc("/hello/{name}", func (w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		fmt.Fprintf(w, "hello %s", vars["name"])
	})
	r.HandleFunc("/hi/{name}/{age}", logging(func (w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		fmt.Fprintf(w, "hi %s and age is %s", vars["name"], vars["age"])
	}))
	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":80", r)
}
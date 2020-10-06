package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"

func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/statics/{index}", func(w http.ResponseWriter, r *http.Request) {
    vars:=mux.Vars(r)
    fmt.Fprintf(w,"Hello to %s!\n",vars["index"])
})
    http.ListenAndServe(":80", r)
}

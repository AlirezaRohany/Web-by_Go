package main

import "fmt"
import "net/http"

func main(){
	fmt.Println("Hello Web By Go!")
	http.HandleFunc("/myindex",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"You have requested: %s\n", r.URL.Path)
	})	
	http.ListenAndServe(":80",nil)
}
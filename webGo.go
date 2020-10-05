package main

import "fmt"
import "net/http"

func main(){
	fmt.Println("Hello Web By Go!")
	http.HandleFunc("/myindex",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w,"You have requested: %s\n", r.URL.Path)
		fmt.Fprintf(w, "Welcome to my website!\n")
		fmt.Fprintf(w,"Get parameters: %s\n",r.URL.Query().Get("token"))

	})	

	fs:=http.FileServer(http.Dir("static/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	
	http.ListenAndServe(":80",nil)


}
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	
	p1:= Page{Title:"TestPage",Body:[]byte("This is a go page!")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}


type Page struct{
	Title string
	Body []byte
}

func(p *Page) save() error {
	filename:= p.Title + ".txt"
	return ioutil.WriteFile(filename,p.Body,0600)
}

func loadPage(title string) (*Page, error){
	filename := title +".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil,err
	}
	return &Page{Title:title, Body:body},nil
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello!, I hate %s!\n", r.URL.Path[1:])
}
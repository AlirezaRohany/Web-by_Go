package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"regexp"
	 "errors"
)


var templates= template.Must(template.ParseFiles("statics/edit.html","statics/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func main(){
	
	p1:= Page{Title:"TestPage",Body:[]byte("This is a go page!")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/",saveHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))

	// todo:Introducing Function Literals and Closures
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
	fmt.Fprintf(w, "Hello!, I have %s!\n", r.URL.Path[1:])
}

// func viewHandler(w http.ResponseWriter, r *http.Request){
// 	title:=r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>",p.Title,p.Body)
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/edit/"):]
//     p, err := loadPage(title)
//     if err != nil {
//         p = &Page{Title: title}
//     }
//     fmt.Fprintf(w, "<h1>Editing %s</h1>"+
//         "<form action=\"/save/%s\" method=\"POST\">"+
//         "<textarea name=\"body\">%s</textarea><br>"+
//         "<input type=\"submit\" value=\"Save\">"+
//         "</form>",
//         p.Title, p.Title, p.Body)
// }

func editHandler(w http.ResponseWriter, r *http.Request){
	// title:= r.URL.Path[len("/edit/"):]
	title, err := getTitle(w,r)
	if err != nil{
		return
	}
	p, err:=loadPage(title)
	if err!=nil{
		p=&Page{Title:title}
	}
	renderTemplate(w,"edit",p)
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	// title:= r.URL.Path[len("/view/"):]
	title, err := getTitle(w,r)
	if err!= nil{
		return
	}
	p, err := loadPage(title)
	if err!=nil{
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func renderTemplate(w http.ResponseWriter , templ string, p *Page){
	// t, err := template.ParseFiles(templ+".html")
	// if err!=nil{
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = t.Execute(w,p)
	// if err!=nil{
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	err:=templates.ExecuteTemplate(w, templ+".html", p)
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	// title:= r.URL.Path[len("/save/"):]
	title, errr := getTitle(w,r)
	if errr != nil{
		return
	}
	body:= r.FormValue("body")
	p:= &Page{Title: title, Body: []byte(body)}
	err:=p.save()
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}
package main

// import "fmt"
// import "net/http"
// import "github.com/gorilla/mux"


import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}


func main(){
// 	r:=mux.NewRouter()
// 	r.HandleFunc("/statics/{index}", func(w http.ResponseWriter, r *http.Request) {
//     vars:=mux.Vars(r)
//     fmt.Fprintf(w,"Hello to %s!\n",vars["index"])
// })
//     http.ListenAndServe(":80", r)


    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)

        fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
    })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        peter := User{
            Firstname: "John",
            Lastname:  "Doe",
            Age:       25,
        }

        json.NewEncoder(w).Encode(peter)
    })

    http.ListenAndServe(":8080", nil)
}

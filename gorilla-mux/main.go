package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/{name}", HelloName)
    router.HandleFunc("/person/{name}", GetPerson)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}

func HelloName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    fmt.Fprint(w, "hello, ", name)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    person := Person{Name: name, Gender: "male"}

    json.NewEncoder(w).Encode(person)
}

type Person struct {
    Name string
    Gender string
}

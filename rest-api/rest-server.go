package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/")
  router.HandleFunc("/todos", TodoIndex)
  router.HandleFunc("/todos/{todoId}", ToDoShow)
  
  log.Fatal(http.ListenAndServe(":8080", nil))
}

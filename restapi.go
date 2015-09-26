package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "time"
  "encoding/json"

  "github.com/gorilla/mux"
)

type Todo struct {
  Name string `json:"name"`
  Completed bool `json:"completed"`
  Due time.Time `json:"due"`
}

type Todos []Todo

func main() {

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/todos", TodoIndex)
  router.HandleFunc("/todos/{todoId}", TodoShow)

  log.Fatal(http.ListenAndServe(":3000", router))

}


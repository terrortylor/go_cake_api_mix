package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Bowl struct {
  Name    string `json:"name,omitempty"`
  Ingredients []Ingredient
}

type Ingredient struct {
  Name     string `json:"name,omitempty"`
  Quantity string `json:"quantity,omitempty"`
}

type NewBowl struct {
  Name     string `json:"name,omitempty"`
}

type Message struct {
  Message     string `json:"message,omitempty"`
}

var bowls = make(map[string]Bowl)

func main() {
  log.Println("Starting go_cake_api_mix")
  router := mux.NewRouter()
  router.HandleFunc("/bowl", PostBowl).Methods("POST")
  router.HandleFunc("/bowl/{name}", GetBowl).Methods("GET")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func PostBowl(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var t NewBowl
  err := decoder.Decode(&t)
  message := ""
  if err != nil {
    rw.WriteHeader(http.StatusInternalServerError)
    message =  "Couldn't parse form data"
  } else {
    if _, ok := bowls[t.Name]; ok {
      rw.WriteHeader(http.StatusInternalServerError)
      message = "Bowl already exists with that name"
    } else {
      newBowl := Bowl{Name: t.Name, Ingredients: []Ingredient{}}
      bowls[t.Name] = newBowl
      message = "Created bowl"
      rw.WriteHeader(http.StatusCreated)
    }
  }
  json.NewEncoder(rw).Encode(Message{Message: message})
  log.Println(message)
}

func GetBowl(rw http.ResponseWriter, r *http.Request) {
  bowlName := mux.Vars(r)["name"]
  log.Println("Looking for bowl:", bowlName)
  if bowl, ok := bowls[bowlName]; ok {
    rw.WriteHeader(http.StatusOK)
    json.NewEncoder(rw).Encode(bowl)
  } else {
    message := "Bowl not found"
    rw.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(rw).Encode(Message{Message: message})
    log.Println(message)
  }
}

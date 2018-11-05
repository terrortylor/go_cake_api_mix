package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

type Bowl struct {
  Name    string `json:"name,omitempty"`
  Ingredients []Ingredient
}

type NewBowl struct {
  Name     string `json:"name,omitempty"`
}

type Ingredient struct {
  Name     string `json:"name,omitempty"`
  Quantity string `json:"quantity,omitempty"`
}

type NewIngredient struct {
  BowlName     string `json:"bowlName,omitempty"`
  Name     string `json:"name,omitempty"`
  Quantity string `json:"quantity,omitempty"`
}

type Message struct {
  Message     string `json:"message,omitempty"`
}

var bowls = make(map[string]Bowl)

func main() {
  log.Println("Starting go_cake_api_mix")
  router := mux.NewRouter()
  router.HandleFunc("/ingredient", PutIngredient).Methods("PUT")
  router.HandleFunc("/bowl", PostBowl).Methods("POST")
  router.HandleFunc("/bowl/{name}", GetBowl).Methods("GET")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func PutIngredient(rw http.ResponseWriter, req *http.Request) {
  rw.Header().Set("Content-Type", "application/json")
  decoder := json.NewDecoder(req.Body)
  var t NewIngredient
  err := decoder.Decode(&t)
  message := ""
  if err != nil {
    rw.WriteHeader(http.StatusInternalServerError)
    message =  "Couldn't parse form data"
  } else {
    if bowl, ok := bowls[t.BowlName]; ok {
      ingredients := &bowl.Ingredients
      NewIngredient := Ingredient{Name: t.Name, Quantity: t.Quantity} 
      *ingredients = append(*ingredients, NewIngredient)
      // Reassign map item
      bowls[t.BowlName] = bowl
      message = "Ingredient added"
      rw.WriteHeader(http.StatusCreated)
      json.NewEncoder(rw).Encode(bowl)
    } else {
      message = "Bowl not found"
      rw.WriteHeader(http.StatusInternalServerError)
      json.NewEncoder(rw).Encode(Message{Message: message})
    }
  }
  log.Println(message)
}

func PostBowl(rw http.ResponseWriter, req *http.Request) {
  rw.Header().Set("Content-Type", "application/json")
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
  rw.Header().Set("Content-Type", "application/json")
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

package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "os"
  "fmt"
)

type Bowl struct {
  CakeName    string `json:"cakeName,omitempty"`
  Ingredients []Ingredient
}

type Ingredient struct {
  Name     string `json:"name,omitempty"`
  Quantity string `json:"quantity,omitempty"`
}

var logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

func main() {
  logger.Output(2, "Starting go_cake_api_mix")
  router := mux.NewRouter()
  router.HandleFunc("/bowl", GetBowl).Methods("GET")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func GetBowl(w http.ResponseWriter, r *http.Request) {
  logger.Output(2, "Responding with a new Bowl JSON")
  ingredientArray := []Ingredient{}
  newBowl := Bowl{CakeName: "A Cake Name", Ingredients: ingredientArray}
  jsonObj, _ := json.Marshal(newBowl)
  fmt.Println(string(jsonObj))
  json.NewEncoder(w).Encode(newBowl)
}

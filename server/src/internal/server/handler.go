package server

import (
  "encoding/json"
  "net/http"
  "fmt"

  "internal/reversi"
)

func UserRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  gameData, correct := reversi.JsonToGameData(w, r)
  if (correct == false) {
    return
  }

  reversi.Transition(&gameData)
  fmt.Print(gameData)

  if err := json.NewEncoder(w).Encode(gameData); err != nil {
    fmt.Println(err);
  }
}

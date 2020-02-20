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
  fmt.Println(gameData)

  if err := json.NewEncoder(w).Encode(gameData); err != nil {
    fmt.Println(err);
  }
}

func SkipTurnHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  gameData, correct := reversi.JsonToGameData(w, r)
  if (correct == false) {
    return
  }

  gameData.Turn += 1
  gameData.Status = "skipped"
  reversi.CanPutPoints(gameData.Turn % 2 + 1, &gameData)
  fmt.Println(gameData)

  if err := json.NewEncoder(w).Encode(gameData); err != nil {
    fmt.Println(err);
  }
}
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
    w.WriteHeader(http.StatusInternalServerError)
    return
  } else {
    w.WriteHeader(http.StatusOK)
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
    w.WriteHeader(http.StatusInternalServerError)
    return
  } else {
    w.WriteHeader(http.StatusOK)
  }

  gameData.Turn += 1
  gameData.Status = "skipped"
  reversi.CanPutPoints(gameData.Turn % 2 + 1, &gameData)
  fmt.Println(gameData)

  if err := json.NewEncoder(w).Encode(gameData); err != nil {
    fmt.Println(err);
  }
}

func SolversHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  solvers := reversi.GetSolvers();

  if err := json.NewEncoder(w).Encode(solvers); err != nil {
    fmt.Println(err)
  }
}

func CPUHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  gameData, correct := reversi.JsonToGameData(w, r)
  if (correct == false) {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  if correct := reversi.RunSolver(&gameData); correct == false {
    w.WriteHeader(http.StatusInternalServerError)    
  } else {
    w.WriteHeader(http.StatusOK)
  }
  reversi.Transition(&gameData)
  fmt.Println(gameData)

  if err := json.NewEncoder(w).Encode(gameData); err != nil {
    fmt.Println(err);
  }
}
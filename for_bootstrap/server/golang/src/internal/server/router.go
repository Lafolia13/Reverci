package server

import (
  "net/http"

  "github.com/gorilla/mux"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}
type Routes []Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  router.
    Methods("POST").
    Path("/userRequest").
    Name("User").
    HandlerFunc(UserRequestHandler)
  
  router.
    Methods("POST").
    Path("/skip").
    Name("Skip").
    HandlerFunc(SkipTurnHandler)

  router.
    Methods("POST").
    Path("/solverRequest").
    Name("Solvers").
    HandlerFunc(SolversHandler)

  router.
    Methods("POST").
    Path("/cpuRequest").
    Name("CPU").
    HandlerFunc(CPUHandler)

  return router
}
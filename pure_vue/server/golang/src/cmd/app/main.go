package main

import (
  "net/http"
  "log"

  "internal/server"
)

func main() {
  router := server.NewRouter()
  log.Fatal(http.ListenAndServe(":8888", router))
}

package server

import (
  "encoding/json"
  "net/http"
  "strconv"
  "io"
  "fmt"
  "log"

  "api/request"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "http://localhost")

  length, err := strconv.Atoi(r.Header.Get("Content-Length"))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  
  bytes := make([]byte, length)
  length, err = r.Body.Read(bytes)
  if err != nil && err != io.EOF {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  
  var jsonBody request.GameData
  if err := json.Unmarshal(bytes, &jsonBody); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    log.Fatal(err)
    return
  }
  jsonBody.Turn += 1
  jsonBody.Field[jsonBody.Request.H][jsonBody.Request.W] = 1
  jsonBody.Status = "ok"
  fmt.Println(jsonBody)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(jsonBody); err != nil {
    log.Fatal(err)
  }
}

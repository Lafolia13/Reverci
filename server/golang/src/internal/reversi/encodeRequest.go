package reversi

import (
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"
	"io"

	"api/request"
)

func JsonToGameData(w http.ResponseWriter, r *http.Request) (gameData request.GameData, correct bool) {
	correct = false

	// Content-Length が nil でないことを確認 (そんな状態あるんですかね)
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Println(err)
		return
  }
	
	// Content-Length 分 文字列を取得
  bytes := make([]byte, length)
  length, err = r.Body.Read(bytes)
  if err != nil && err != io.EOF {
    fmt.Println(err)
		return
  }
	
	// request.GameData にエンコーディング
  if err := json.Unmarshal(bytes, &gameData); err != nil {
    fmt.Println(err)
		return
	}

	correct = len(gameData.Field) <= 20

	return
}
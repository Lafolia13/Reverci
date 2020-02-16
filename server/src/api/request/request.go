package request

type Position struct {
  H int `json:"h"`
  W int `json:"w"`
}

type GameData struct {
  Turn    int       `json:"turn"`
  Field   [][]int   `json:"field"`
  Request Position  `json:"request"`
  Status  string    `json:"status"`
}
package request

type Point struct {
  H int `json:"h"`
  W int `json:"w"`
}

type Scores struct {
  Black int `json:"black"`
  White int `json:"white"`
}

type GameData struct {
  Turn    int       `json:"turn"`
  Score   Scores     `json:"score"`
  Field   [][]int   `json:"field"`
  Request Point  `json:"request"`
  Status  string    `json:"status"`
}
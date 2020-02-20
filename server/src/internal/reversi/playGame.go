package reversi

import (
  "api/request"
)	
  
var next_to_h = [8] int{1, 1, 1, 0, 0, -1, -1, -1}
var next_to_w = [8] int{-1, 0, 1, -1, 1, -1, 0, 1}
  
func Transition(gameData *request.GameData) {
  now := request.Point{
    H: gameData.Request.H,
    W: gameData.Request.W,
  }	

  if (gameData.Field[now.H][now.W] != 3) {
    gameData.Status = "failed"
    return
  }

  player_color := gameData.Turn % 2 + 1

  gameData.Field[now.H][now.W] = player_color;
  gameData.Turn += 1
  if (player_color == 1) {
    gameData.Score.Black += 1
  } else {
    gameData.Score.White += 1
  }

  for i := 0; i < 8; i++ {
    next := request.Point{
      H: now.H + next_to_h[i],
      W: now.W + next_to_w[i],
    }	
    diff := TurnOver(next, now, i, player_color, gameData)
    if (player_color == 1) {
      gameData.Score.Black += diff
      gameData.Score.White -= diff
    } else {
      gameData.Score.Black -= diff
      gameData.Score.White += diff
    }
  }

  able_now_player := CanPutPoints(player_color, gameData)
  able_next_player := CanPutPoints(player_color^3, gameData)
  if (able_next_player == true) {
    gameData.Status = "success"
  } else if (able_now_player == false) {
    gameData.Status = "finish"
  } else {
    gameData.Status = "skip"
  }	
    
  return
}	
      
func TurnOver(now request.Point, start request.Point,
    dir int, player_color int, gameData *request.GameData) int {
  width := len(gameData.Field)
  if (now.H < 0 || now.W < 0 || now.H >= width || now.W >= width) {
    return 0
  }

  color := gameData.Field[now.H][now.W]
  if (color == player_color) {
    return Max(Abs(now.H - start.H), Abs(now.W - start.W)) - 1;
  } else if (color == player_color ^ 3) {
    next := request.Point{
      H: now.H + next_to_h[dir],
      W: now.W + next_to_w[dir],
    }
    ret := TurnOver(next, start, dir, player_color, gameData);
    if (ret > 0) {
      gameData.Field[now.H][now.W] = player_color
    }

    return ret
  } else {
    return 0
  }
}

      
func Max(a int, b int) int {
  if (a > b) {
    return a
  } else {
    return b
  }	
}	
            
func Abs(a int) int {
  if (a > 0) {
    return a
  } else {
    return -a
  }	
}

func CanPutPoints(player_color int, gameData *request.GameData) bool {
  ok := false
  width := len(gameData.Field)
  for h := 0; h < width; h++ {
    for w := 0; w < width; w++ {
      now := request.Point{
        H: h,
        W: w,
      }
      gameData.Field[h][w] = CanPutPoint(now, player_color, gameData)
      ok = ok || gameData.Field[h][w] == 3
    }
  }

  return ok
}

func CanPutPoint(now request.Point, player_color int,
    gameData *request.GameData) int {
  if (gameData.Field[now.H][now.W] == 1 || gameData.Field[now.H][now.W] == 2) {
    return gameData.Field[now.H][now.W];
  }

  width := len(gameData.Field)
  var gameData_copy request.GameData
  gameData_copy.Field = make([][]int, width, width)
  for h := 0; h < width; h++ {
    gameData_copy.Field[h] = make([]int, width, width)
    for w := 0; w < width; w++ {
      gameData_copy.Field[h][w] = gameData.Field[h][w]
    }
  }

  ok := false
  for i := 0; i < 8; i++ {
    next := request.Point{
      H: now.H + next_to_h[i],
      W: now.W + next_to_w[i],
    }
    ok = ok || TurnOver(next, now, i, player_color, &gameData_copy) > 0
  }

  if (ok == true) {
    return 3
  } else {
    return 0
  }
}


package reversi

import (
	"api/request"
)

var next_to_h = [8] int{1, 1, 1, 0, 0, -1, -1, -1}
var next_to_w = [8] int{-1, 0, 1, -1, 1, -1, 0, 1}

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

func TurnOver(now request.Position, start request.Position,
		dir int, player_color int, gameData *request.GameData) int {
	width := len(gameData.Field[0])
	if (now.H < 0 || now.W < 0 || now.H >= width || now.W >= width) {
		return 0
	}

	color := gameData.Field[now.H][now.W]
	if (color == player_color) {
		return Max(Abs(now.H - start.H), Abs(now.W - start.W)) - 1;
	} else if (color == player_color ^ 3) {
		next := request.Position{
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

func ablePutPosition(now request.Position, player_color int,
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
		next := request.Position{
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

func ablePutPositions(player_color int, gameData *request.GameData) bool {
	ok := false
	for h := range gameData.Field {
		for w := range gameData.Field {
			now := request.Position{
				H: h,
				W: w,
			}
			gameData.Field[h][w] = ablePutPosition(now, player_color, gameData)
			ok = ok || gameData.Field[h][w] == 3
		}
	}

	return ok
}

func Transition(gameData *request.GameData) {
	now := request.Position{
		H: gameData.Request.H,
		W: gameData.Request.W,
	}

	if (gameData.Field[now.H][now.W] != 3) {
		gameData.Status = "failed"
		return
	}

	var player_color int
	if (gameData.Turn % 2 == 0) {
		player_color = 1;
	} else {
		player_color = 2;
	}
	
	gameData.Field[now.H][now.W] = player_color;
	gameData.Turn += 1

	for i := 0; i < 8; i++ {
		next := request.Position{
			H: now.H + next_to_h[i],
			W: now.W + next_to_w[i],
		}
		_ = TurnOver(next, now, i, player_color, gameData)
	}

	able_now_player := ablePutPositions(player_color, gameData)
	able_next_player := ablePutPositions(player_color^3, gameData)

	if (able_next_player == true) {
		gameData.Status = "success"
	} else if (able_now_player == false) {
		gameData.Status = "finish"
	} else {
		gameData.Status = "skip"
	}

	return
}
#ifndef _GAME_CPP_
#define _GAME_CPP_

#include <vector>
#include <cmath>

using namespace std;

struct point {
  int h;
  int w;

  point() = default;
  point(int h, int w) :
    h(h),
    w(w)
  {}

  point operator+(const point &another) const {
    return point(h + another.h, w + another.w);
  }
  bool operator==(const point &another) const {
    return h == another.h && w == another.w;
  }
};

struct game_data {
  int turn = {};
  int width = {};
  vector<vector<int>> area = {};

  game_data() = default;
};

// 列挙型
// BLACK = 0, WHITE = 1, NONE = 2 になります
enum MasuState {
  BLACK,
  WHITE,
  NONE
};

vector<point> neighborhood = {
  {1, -1}, {1, 0}, {1, 1},
  {0, -1},         {0, 1},
  {-1, -1}, {-1, 0}, {-1, 1}
};

// 引数が異なる場合は同じ関数名で定義できます
int put(point now, point start, int dir, int color, game_data &gm) {
  if (now.h < 0 || now.w < 0 || now.h >= gm.width || now.w >= gm.width) return 0;

  if (gm.area[now.h][now.w] == NONE) {
    return 0;
  }
  else if (gm.area[now.h][now.w] == color) {
  	// 石を置いた場所からの距離 - 1 がひっくり返せる石の個数
    return max(abs(now.h - start.h), abs(now.w - start.w)) - 1;
  }
  else {
    int ret = put(now + neighborhood[dir], start, dir, color, gm);

    if (ret > 0) {
      gm.area[now.h][now.w] = color;
    }

    return ret;
  }
}


void put(point pos, game_data &gm) {
  int color, ret = 0;

  if (gm.turn % 2 == 0) {
    color = BLACK;
  }
  else {
    color = WHITE;
  }

  gm.area[pos.h][pos.w] = color;
  for (int i = 0; i < 8; ++i) {
    put(pos + neighborhood[i], pos, i, color, gm);
  }

  return;
}

void put(int h, int w, game_data &gm) {
	return put(point(h, w), gm);
}

bool able_to_put(int th, int tw, game_data gm) {
  if (th < 0 || tw < 0 || th >= gm.width || tw >= gm.width) return false;

  int color, ret = 0;

  if (gm.turn % 2 == 0) {
    color = BLACK;
  }
  else {
    color = WHITE;
  }

  if (gm.area[th][tw] != NONE) return false;

  for (int i = 0; i < 8; ++i) {
    ret += put(point(th, tw) + neighborhood[i], point(th, tw), i, color, gm);
  }

  // 石が置ける → 置くことでひっくり返る石が存在する
  return ret > 0;
}

int get_score(int color, game_data &gm) {
  int ret = 0;
  for (auto &h: gm.area) {
    for (auto &w: h) {
      ret += w == color;
    }
  }

  return ret;
}

#endif
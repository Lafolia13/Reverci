#ifndef _BEAM_CPP_
#define _BEAM_CPP_

#include <queue>

#include "evaluation.cpp"
#include "game.cpp"

struct node {
  game_data gm;
  double evaluation;
  point first_action;

  node() = default;
  node(game_data gm, double evaluation):
    gm(gm),
    evaluation(evaluation)
  {}

  bool operator>(const node &right) const {
    return evaluation > right.evaluation;
  }
};

vector<point> get_actions(game_data gm) {
  vector<point> ret;

  for (int h = 0; h < gm.width; ++h) {
    for (int w = 0; w < gm.width; ++w) {
      if (able_to_put(h, w, gm)) {
        ret.push_back(point(h, w));
      }
    }
  }

  return ret;
}

double get_evaluation(int decide_color, point action, node &check_node, int width) {
  int color;
  if (check_node.gm.turn % 2 == 0) {
    color = BLACK;
  }
  else {
    color = WHITE;
  }

  double s_diff = score_difference(color, check_node.gm);
  double n_conner = near_corner(action, width);

  double ret = s_diff +
      n_conner;

  if (color != decide_color) {
    ret = rival_evaluation(ret);
  }

  return ret + before_evaluation(check_node.evaluation);
}

point beam_search(game_data gm) {
  constexpr int beam_width = 100;
  constexpr int beam_depth = 10;

  priority_queue<node, vector<node>, greater<node>> now_que, next_que;

  int decide_color;
  if (gm.turn % 2 == 0) {
    decide_color = BLACK;
  }
  else {
    decide_color = WHITE;
  }

  node root_node(gm, 0);
  now_que.push(root_node);
  for (int turn = 0; turn < beam_width; ++turn) {
    while (now_que.size()) {
      // node を一つ取り出す
      const node now_node = now_que.top();
      now_que.pop();

      // node からできる選択肢をすべて取得する
      vector<point> actions = get_actions(now_node.gm);

      if (actions.size() == 0) {
        next_que.push(now_node);
        continue;
      }
      for (auto action: actions) {
        node next_node = now_node;
        // それぞれの選択肢を実行する
        put(action, next_node.gm);
        // 評価値を得る
        next_node.evaluation = get_evaluation(decide_color, action, next_node, gm.width);
        next_node.gm.turn++;

        if (turn == 0) {
          next_node.first_action = action;
        }

        next_que.push(next_node);
        if (next_que.size() > beam_width) {
          next_que.pop();
        }
      }
    }

    swap(now_que, next_que);
  }

  while (now_que.size() > 2) {
    now_que.pop();
  }

  return now_que.top().first_action;
}

#endif

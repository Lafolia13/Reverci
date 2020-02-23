#ifndef _EVALUATION_CPP_
#define _EVALUATION_CPP_

#include "game.cpp"

double before_evaluation(double evaluation) {
  double bias = 1.0;

  return evaluation * bias;
}

double score_difference(int color, game_data &gm) {
  double bias = 1.0;

  double ret;
  double black_score = get_score(BLACK, gm);
  double white_score = get_score(WHITE, gm);
  if (color == BLACK) {
    ret = black_score - white_score;
  }
  else {
    ret = white_score - black_score;
  }

  return ret * bias;
}

double rival_evaluation(double evaluation) {
  double bias = -0.6;

  return evaluation * bias;
}

double near_corner(point action, int width) {
  double bias = 1000;

  double ret;
  if (action == point(0, 0) ||
      action == point(0, width - 1) ||
      action == point(width - 1, 0) ||
      action == point(width - 1, width - 1)) {
    ret = 1;
  }
  else if (action == point(0, 1) ||
      action == point(0, width - 2) ||
      action == point(1, 0) ||
      action == point(1, width - 1) ||
      action == point(width - 2, 0) ||
      action == point(width - 2, width - 1) ||
      action == point(width - 1, 1) ||
      action == point(width - 1, width - 2)) {
    ret = -1;
  }

  return ret * bias;
}

#endif
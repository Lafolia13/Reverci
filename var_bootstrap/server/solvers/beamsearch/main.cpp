#include <iostream>

#include "game.cpp"
#include "beam.cpp"

using namespace std;

bool output(point p) {
  cout << p.h << " " << p.w << endl;

  return true;
}

bool input(game_data &gm) {
  cin >> gm.turn;
  cin >> gm.width;
  gm.area.resize(gm.width, vector<int>(gm.width));
  for (auto &h: gm.area) {
    for (auto &w: h) {
      cin >> w;
      if (w == 1) {
        w = BLACK;
      }
      else if (w == 2) {
        w = WHITE;
      } else {
        w = NONE;
      }
    }
  }

  return true;
}

int main() {
  game_data gm;
  input(gm);
  output(beam_search(gm));

  return 0;
}
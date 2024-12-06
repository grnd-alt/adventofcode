#include "fstream"
#include "iostream"
#include <algorithm>
#include <string>
#include <tuple>
#include <vector>
using namespace std;

// "^" -1,0
// ">"  0,1
// "v"  1,0
// "<"  0,-1

vector<int> getDirection(int oldx, int oldy) {
  if (oldx == -1) {
    return {0, 1};
  }
  if (oldx == 1) {
    return {0, -1};
  }
  if (oldy == 1) {
    return {1, 0};
  }
  if (oldy == -1) {
    return {-1, 0};
  }
  return {0, 0};
}

tuple<vector<string>, int> solve1(int x, int y, vector<string> grid) {
  int count = 0;
  vector<int> dir = {-1, 0};
  do {
    if (grid[x][y] != 'X') {
      grid[x][y] = 'X';
      count++;
    }
    int newx = dir[0] + x;
    int newy = dir[1] + y;
    if (newx < 0 || newy < 0) {
      break;
    }
    if (newx >= grid.size() || newy >= grid[x].length()) {
      break;
    }
    if (grid[newx][newy] == '#') {
      dir = getDirection(dir[0], dir[1]);
    } else {
      x = newx;
      y = newy;
    }
  } while (true);
  return make_tuple(grid, count);
}

tuple<vector<string>, int> solve2(int sx, int sy, vector<string> grid,vector<string> solvedGrid) {
  int loops = 0;
  vector<vector<int>> dirs = {{-1, 0}, {0, 1}, {1, 0}, {0, -1}};
  for (int blockx = 0; blockx < grid.size(); blockx++) {
    for (int blocky = 0; blocky < grid[0].length(); blocky++) {
      if (solvedGrid[blockx][blocky] != 'X') {
        continue;
      }
      int x = sx;
      int y = sy;
      int dir = 0;
      vector<vector<int>> seen;
      while (true) {
        vector<int> element = {x, y, dir};
        if (find(seen.begin(), seen.end(), element) != seen.end()) {
          loops++;
          break;
        }
        seen.push_back({x, y, dir});
        int newx = x + dirs[dir][0];
        int newy = y + dirs[dir][1];
        if (newx < 0 || newy < 0) {
          break;
        }
        if (newx >= grid.size() || newy >= grid[x].length()) {
          break;
        }

        if (grid[newx][newy] == '#' || (newx == blockx && newy == blocky)) {
          dir = (dir + 1) % 4;
        } else {
          x = newx;
          y = newy;
        }
      }
    }
  }
  return make_tuple(grid, loops);
}

int main(int argc, char *argv[]) {
  string text;
  ifstream file("input.txt");
  vector<string> lines;
  while (getline(file, text)) {
    lines.push_back(text);
  }

  for (int i = 0; i < lines.size(); i++) {
    string line = lines[i];
    for (int j = 0; j < lines[i].length(); j++) {
      if (line[j] == '^') {
        auto [sovled, count] = solve1(i, j, lines);
        cout << count << endl;
        auto [grid, loops] = solve2(i, j, lines, sovled);
        cout << loops << endl;
        break;
      }
    }
  }

  return 0;
}

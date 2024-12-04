def countword(grid, x, y, word):
    d = [True, True, True, True]
    h = [True, True]
    v = [True, True]
    for i in range(1, len(word)):
        if y + i > len(grid[0]) - 1 or not grid[x][y + i] == word[i]:
            h[0] = False
        if y - i < 0 or not grid[x][y - i] == word[i]:
            h[1] = False
        if x + i > len(grid) - 1 or not grid[x + i][y] == word[i]:
            v[0] = False
        if x - i < 0 or not grid[x - i][y] == word[i]:
            v[1] = False
        if (
            x + i > len(grid) - 1
            or y + i > len(grid[0]) - 1
            or not grid[x + i][y + i] == word[i]
        ):
            d[0] = False
        if x + i > len(grid) - 1 or y - i < 0 or not grid[x + i][y - i] == word[i]:
            d[1] = False
        if x - i < 0 or y + i > len(grid[0]) - 1 or not grid[x - i][y + i] == word[i]:
            d[2] = False
        if x - i < 0 or y - i < 0 or not grid[x - i][y - i] == word[i]:
            d[3] = False

    return sum(d) + sum(h) + sum(v)


def getValueFromMiddle(ind, word):
    print(ind, word)
    return word[int(len(word) / 2) + ind]


def checkCross(grid, x, y, word):
    if len(word) % 3 != 0:
        return 0
    topleft = True
    topright = True

    for i in range(1, 2):
        if x + i >= len(grid) or x - i < 0 or y + i >= len(grid[x + 1]) or y - i < 0:
            topleft = False
            topright = False
            return 0

        if grid[x - i][y - i] != "M":
            if grid[x + i][y + i] != "M":
                topleft = False
            else:
                if grid[x - i][y - i] != "S":
                    topleft = False
        else:
            if grid[x + i][y + i] != "S":
                topleft = False

        if grid[x + i][y - i] != "M":
            if grid[x - 1][y + 1] != "M":
                topright = False
            else:
                if grid[x + i][y - i] != "S":
                    topright = False
        else:
            if grid[x - i][y + 1] != "S":
                topright = False
    return topright and topleft


def main():
    count = 0
    countv2 = 0
    with open("input.txt", "r") as file:
        input = file.read()
        lines = input.split("\n")
        empty = lambda x: len(x) == 0
        lines = [x for x in lines if not empty(x)]
        for i, line in enumerate(lines):
            for j, letter in enumerate(line):
                if letter == "X":
                    count = count + countword(lines, i, j, "XMAS")
                if letter == "A":
                    countv2 = countv2 + checkCross(lines, i, j, "MAS")
    print("version 1: ", count)
    print("version 2: ", countv2)


main()

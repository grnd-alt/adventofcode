import re


def solve1():
    with open("input.txt", "r") as file:
        data = file.read().replace("\n", "")
        muls = re.findall(r"mul\(\d+,\d+\)", data)
        res = 0
        for mul in muls:
            ints = re.findall(r"\d+", mul)
            res += int(ints[0]) * int(ints[1])
        return res


def solve2():
    with open("input.txt", "r") as file:
        data = file.read().replace("\n", "")
        muls = re.findall(r"mul\(\d+,\d+\)|don't|do", data)
        res = 0
        do = True
        for mul in muls:
            if mul == "do":
                do = True
                continue
            if mul == "don't":
                do = False
                continue
            if not do:
                continue
            ints = re.findall(r"\d+", mul)
            res += int(ints[0]) * int(ints[1])
        return res


print(solve1())
print(solve2())

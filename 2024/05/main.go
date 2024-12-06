package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput() (map[int][]int, [][]int) {
	cont, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(cont), "\n")

	rules := make(map[int][]int)
	inRules := true

	updates := make([][]int, 0)
	for _, line := range lines {
		if line == "" {
			inRules = false
			continue
		}
		if inRules {
			smaller, err := strconv.Atoi(strings.Split(line, "|")[0])
			bigger, err := strconv.Atoi(strings.Split(line, "|")[1])
			if err != nil {
				panic(err)
			}
			rules[smaller] = append(rules[smaller], bigger)
		} else {
			update := make([]int, 0)
			for _, str := range strings.Split(line, ",") {
				num, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func checkAllowed(x int, y int, i int, j int, rules map[int][]int) bool {
	if i == j {
		return true
	}

	if i < j {
		if slices.Contains(rules[y], x) {
			return false
		}
	}
	if j < i {
		if slices.Contains(rules[x], y) {
			return false
		}
	}
	return true
}

func fixUpdate(update []int, rules map[int][]int) {
}

func solve1(rules map[int][]int, updates [][]int) int {
	valid := 0
	for _, update := range updates {
		validUpdate := true
		for i, val := range update {
			if !validUpdate {
				break
			}
			for j, val2 := range update {
				if !checkAllowed(val, val2, i, j, rules) {
					validUpdate = false
					break
				}
			}
		}
		if validUpdate {
			valid += update[len(update)/2]
		}
	}
	return valid
}

func solve2(rules map[int][]int, updates [][]int) int {
  res := 0
	for _, update := range updates {
    validUpdate := true
    fixed := false
		for true {
			for i, x := range update {
        if !validUpdate {
          break
        }
				for j, y := range update {
					if !checkAllowed(x, y, i, j, rules) {
            validUpdate = false
            fixed = true
            update[i] = y
            update[j] = x
            break
					}
				}
			}
      if validUpdate {
        if fixed {
          res += update[len(update)/2]
        }
        break;
      }
      validUpdate = true
		}
	}
	return res
}
func main() {
	rules, updates := parseInput()
	fmt.Println(solve1(rules, updates))
	fmt.Println(solve2(rules, updates))
}

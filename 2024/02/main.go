package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	cont, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(cont), "\n")

	reports := [][]int{}
	for _, element := range lines {
		arr := []int{}
		nums := strings.Split(element, " ")
		for _, stringVal := range nums {
			num, err := strconv.Atoi(stringVal)
			if err != nil {
				fmt.Println(err)
				break
			}
			arr = append(arr, num)
		}
		if len(arr) > 0 {
			reports = append(reports, arr)
		}
	}
	return reports
}

func checkTwo(val1 int, val2 int, decreasing bool) bool {
	big := val2
	small := val1
	if decreasing {
		big = val1
		small = val2
	}
	if big > small && big-small <= 3 {
		return true
	}
	return false
}

func checkReport(report []int) int {
	decreasing := false
	for i := 0; i < len(report)-1; i++ {
		item := report[i]
		next := report[i+1]
		if i == 0 {
			if next < item {
				decreasing = true
			} else if next > item {
				decreasing = false
			} else {
				return i
			}
		}
		if checkTwo(item, next, decreasing) {
			continue
		} else {
			return i
		}
	}
	return -1
}

func solve1(reports [][]int) int {
	savereports := 0
	for _, report := range reports {
		if checkReport(report) == -1 {
			savereports++
		}
	}
	return savereports
}

func solve2(reports [][]int) int {
	savereports := 0
	for _, report := range reports {
		check := checkReport(report)
		if check == -1 {
			savereports++
		} else {
			for i := 0; check != -1 && i < len(report); i++ {
				stripped := make([]int, len(report))
				copy(stripped, report)

				stripped = append(stripped[:i], stripped[i+1:]...)
				check := checkReport(stripped)
				if check == -1 {
					savereports++
					break
				}
			}
		}
	}
	return savereports
}

func main() {
	reports := parseInput()
	// fmt.Println(solve1(reports))
	fmt.Println(solve2(reports))
	// fmt.Println("HELLO ")
}

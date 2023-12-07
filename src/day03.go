package src

import (
	"os"
	"strconv"
	"strings"
)

func (s *Service) CalculateDay03FirstPart() int {
	fileName := "./data/day03/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	symbolMap := make(map[int]map[int]rune)
	for i, l := range lines {
		symbolMap[i] = make(map[int]rune)
		for j, r := range []rune(l) {
			symbolMap[i][j] = r
		}
	}

	res := 0
	for i := 0; i <= len(symbolMap)-1; i++ {
		for j := 0; j <= len(symbolMap[0])-1; j++ {
			if s.isDigit(symbolMap[i][j]) {
				startJ := j
				if startJ > 0 {
					startJ--
				}
				num := 0
				for ; j <= len(symbolMap[0])-1 && s.isDigit(symbolMap[i][j]); j++ {
					num *= 10
					num += s.runeToDigit(symbolMap[i][j])
				}
				finishJ := j - 1
				if finishJ < len(symbolMap[0])-1 {
					finishJ++
				}
				startI := i
				if startI > 0 {
					startI--
				}
				finishI := i
				if finishI < len(symbolMap)-1 {
					finishI++
				}

				near := true
				for ci := startI; near && ci <= finishI; ci++ {
					for cj := startJ; near && cj <= finishJ; cj++ {
						if !s.isDigit(symbolMap[ci][cj]) && symbolMap[ci][cj] != '.' {
							near = false
						}
					}
				}

				if !near {
					res += num
				}
			}
		}
	}

	return res
}

func (s *Service) CalculateDay03SecondPart() int {
	fileName := "./data/day03/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	symbolMap := make(map[int]map[int]rune)
	for i, l := range lines {
		symbolMap[i] = make(map[int]rune)
		for j, r := range []rune(l) {
			symbolMap[i][j] = r
		}
	}

	res := 0
	gears := map[string][]int{}
	for i := 0; i <= len(symbolMap)-1; i++ {
		for j := 0; j <= len(symbolMap[0])-1; j++ {
			if s.isDigit(symbolMap[i][j]) {
				startJ := j
				if startJ > 0 {
					startJ--
				}
				num := 0
				for ; j <= len(symbolMap[0])-1 && s.isDigit(symbolMap[i][j]); j++ {
					num *= 10
					num += s.runeToDigit(symbolMap[i][j])
				}
				finishJ := j - 1
				if finishJ < len(symbolMap[0])-1 {
					finishJ++
				}
				startI := i
				if startI > 0 {
					startI--
				}
				finishI := i
				if finishI < len(symbolMap)-1 {
					finishI++
				}

				for ci := startI; ci <= finishI; ci++ {
					for cj := startJ; cj <= finishJ; cj++ {
						if symbolMap[ci][cj] == '*' {
							key := strconv.Itoa(ci) + "_" + strconv.Itoa(cj)
							gears[key] = append(gears[key], num)
						}
					}
				}
			}
		}
	}

	for _, gearSet := range gears {
		if len(gearSet) == 2 {
			res += gearSet[0] * gearSet[1]
		}
	}

	return res
}

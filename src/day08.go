package src

import (
	"os"
	"strings"
)

type FieldOf struct {
	L string
	R string
}

func (s *Service) CalculateDay08FirstPart() int {
	fileName := "./data/day08/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	instructions := map[string]FieldOf{}
	for _, l := range lines {
		f := strings.Split(l, " = ")
		filed := strings.Split(strings.Trim(f[1], "()"), ", ")
		instructions[f[0]] = FieldOf{L: filed[0], R: filed[1]}
	}

	way := strings.Split("LRRLRRRLRRRLLRRLRRLRLRLRRLLRRLRRLRRRLLLRRRLRRRLRRRLLRRRLRRLLRRLRRLRLRRRLRRLRLRRLRRRLLRRLLRLRRRLLRRLRRLLLRLRRRLRLRLRLLRRRLRLLRRRLRLRRRLRRRLLRRLRRRLLRRLRLLRLRRLLLRRLRRLLLRLLRLRRRLRLRLRRRLRRLLRRRLRLRLRRLRRRLRLRRLRRLRRRLRRRLRRRLRRRLRRLLRRLRLLRRLLRRRLRLLRLRLRRLRRLRLRLRRRLRLRLRRLRLRRLRRRR", "")
	res := 0
	currentPosition := "AAA"
	for i := 0; ; i++ {
		if i >= len(way) {
			i = 0
		}
		step := way[i]
		if step == "R" {
			currentPosition = instructions[currentPosition].R
		} else {
			currentPosition = instructions[currentPosition].L
		}
		res += 1
		if currentPosition == "ZZZ" {
			break
		}
	}

	return res
}

func (s *Service) CalculateDay08SecondPart() int64 {
	fileName := "./data/day08/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	instructions := map[string]FieldOf{}
	for _, l := range lines {
		f := strings.Split(l, " = ")
		filed := strings.Split(strings.Trim(f[1], "()"), ", ")
		instructions[f[0]] = FieldOf{L: filed[0], R: filed[1]}
	}

	way := strings.Split("LRRLRRRLRRRLLRRLRRLRLRLRRLLRRLRRLRRRLLLRRRLRRRLRRRLLRRRLRRLLRRLRRLRLRRRLRRLRLRRLRRRLLRRLLRLRRRLLRRLRRLLLRLRRRLRLRLRLLRRRLRLLRRRLRLRRRLRRRLLRRLRRRLLRRLRLLRLRRLLLRRLRRLLLRLLRLRRRLRLRLRRRLRRLLRRRLRLRLRRLRRRLRLRRLRRLRRRLRRRLRRRLRRRLRRLLRRLRLLRRLLRRRLRLLRLRLRRLRRLRLRLRRRLRLRLRRLRLRRLRRRR", "")

	currentPositions := []string{}

	for pos := range instructions {
		if strings.Split(pos, "")[2] == "A" {
			currentPositions = append(currentPositions, pos)
		}
	}

	res := int64(0)
	for _, cPos := range currentPositions {
		resPos := int64(0)
		for i := 0; ; i++ {
			if i >= len(way) {
				i = 0
			}
			step := way[i]
			if step == "R" {
				cPos = instructions[cPos].R
			} else {
				cPos = instructions[cPos].L
			}
			resPos += 1
			if strings.Split(cPos, "")[2] == "Z" {
				break
			}
		}
		res = nok(res, resPos)
	}

	return res
}

func nok(a, b int64) int64 {
	res := int64(1)
	if a == 0 {
		return b
	}
	for i := int64(2); i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			a = a / i
			b = b / i
			res *= i
			i--
		}
	}
	res *= a * b
	return res
}

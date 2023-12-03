package src

import (
	"os"
	"strconv"
	"strings"
)

func (s *Service) CalculateDay02FirstPart() int {
	fileName := "./data/day02/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	res := 0

	for i, l := range lines {
		game := strings.Split(l, ":")[1]
		sets := strings.Split(game, ";")
		gameOk := true
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cubeColors := range cubes {
				colorCount := strings.Split(strings.Trim(cubeColors, " "), " ")
				count, _ := strconv.Atoi(colorCount[0])
				switch strings.Trim(colorCount[1], " ") {
				case "red":
					if count > 12 {
						gameOk = false
					}
					break
				case "green":
					if count > 13 {
						gameOk = false
					}
					break
				case "blue":
					if count > 14 {
						gameOk = false
					}
					break
				}
			}
		}
		if gameOk {
			res += i + 1
		}

	}
	return res
}
func (s *Service) CalculateDay02SecondPart() int {
	fileName := "./data/day02/input2"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	res := 0

	for _, l := range lines {
		game := strings.Split(l, ":")[1]
		sets := strings.Split(game, ";")
		gameCubes := map[string]int{
			"red": 0,
			"blue": 0,
			"green": 0,
		}
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cubeColors := range cubes {
				colorCount := strings.Split(strings.Trim(cubeColors, " "), " ")
				count, _ := strconv.Atoi(colorCount[0])
				if gameCubes[strings.Trim(colorCount[1], " ")] < count {
					gameCubes[strings.Trim(colorCount[1], " ")] = count
				}
			}
		}
		power := gameCubes["red"] * gameCubes["green"] * gameCubes["blue"]
		res += power
	}
	return res
}
package src

import (
	"os"
	"strings"
)

func (s *Service) CalculateDay04FirstPart() int {
	fileName := "./data/day04/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	res := 0
	for _, l := range lines {
		card := strings.Split(l, ":")[1]
		numbers := strings.Split(card, "|")
		winings := strings.Split(strings.Trim(numbers[0], " "), " ")
		got := strings.Split(strings.Trim(numbers[1], " "), " ")
		winMap := map[string]struct{}{}
		for _, w := range winings {
			if w != "" {
				winMap[w] = struct{}{}
			}
		}

		cardRes := 0
		for _, g := range got {
			if _, ok := winMap[g]; ok {
				if cardRes == 0 {
					cardRes = 1
				} else {
					cardRes *= 2
				}
			}
		}

		res += cardRes
	}

	return res
}

func (s *Service) CalculateDay04SecondPart() int {
	fileName := "./data/day04/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	cardNumberMap := map[int]int{}

	for i := range lines {
		cardNumberMap[i] = 1
	}

	res := 0
	for i, l := range lines {
		card := strings.Split(l, ":")[1]
		numbers := strings.Split(card, "|")
		winings := strings.Split(strings.Trim(numbers[0], " "), " ")
		got := strings.Split(strings.Trim(numbers[1], " "), " ")
		winMap := map[string]struct{}{}
		for _, w := range winings {
			if w != "" {
				winMap[w] = struct{}{}
			}
		}

		winNumbers := 0
		for _, g := range got {
			if _, ok := winMap[g]; ok {
				winNumbers++
			}
		}

		for j := 1; j <= winNumbers; j++ {
			cardNumberMap[i+j] += cardNumberMap[i]
		}
	}

	for _, num := range cardNumberMap {
		res += num
	}

	return res
}

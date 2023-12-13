package src

import (
	"os"
	"strings"
)

func (s *Service) CalculateDay13FirstPart() int {
	fileName := "./data/day13/input1"
	file, _ := os.ReadFile(fileName)

	puzzles := strings.Split(string(file), "\n\n")

	res := 0
	for _, puzzle := range puzzles {
		mutrix := [][]rune{}
		lines := strings.Split(puzzle, "\n")
		for _, l := range lines {
			mutrix = append(mutrix, []rune(l))
		}

		resM := 0
		for i := 1; i < len(mutrix); i++ {
			isMirror := true
			for j := i - 1; j >= 0 && i+i-j-1 < len(mutrix); j-- {
				if string(mutrix[j]) != string(mutrix[i+i-j-1]) {
					isMirror = false
					break
				}
			}
			if isMirror {
				resM += 100 * i
				break
			}
		}
		if resM != 0 {
			res += resM
			continue
		}
		invertM := s.invert(mutrix)
		for i := 1; i < len(invertM); i++ {
			isMirror := true
			for j := i - 1; j >= 0 && i+i-j-1 < len(invertM); j-- {
				if string(invertM[j]) != string(invertM[i+i-j-1]) {
					isMirror = false
					break
				}
			}
			if isMirror {
				resM += i
				break
			}
		}
		if resM != 0 {
			res += resM
			continue
		}
	}

	return res
}

func (s *Service) invert(m [][]rune) [][]rune {
	res := make([][]rune, len(m[0]))
	for i := 0; i < len(res); i++ {
		res[i] = make([]rune, len(m))
	}
	for i, r := range m {
		for j, c := range r {
			res[j][i] = c
		}
	}
	return res
}

func (s *Service) CalculateDay13SecondPart() int {
	fileName := "./data/day13/input1"
	file, _ := os.ReadFile(fileName)

	puzzles := strings.Split(string(file), "\n\n")

	res := 0
	for _, puzzle := range puzzles {
		mutrix := [][]rune{}
		lines := strings.Split(puzzle, "\n")
		for _, l := range lines {
			mutrix = append(mutrix, []rune(l))
		}

		resM := 0
		for i := 1; i < len(mutrix); i++ {
			diff := 0
			for j := i - 1; j >= 0 && i+i-j-1 < len(mutrix); j-- {
				for k := 0; k < len(mutrix[j]); k++ {
					if mutrix[j][k] != mutrix[i+i-j-1][k] {
						diff++
					}
				}
			}
			if diff == 1 {
				resM += 100 * i
				break
			}
		}
		if resM != 0 {
			res += resM
			continue
		}
		invertM := s.invert(mutrix)
		for i := 1; i < len(invertM); i++ {
			diff := 0
			for j := i - 1; j >= 0 && i+i-j-1 < len(invertM); j-- {
				for k := 0; k < len(invertM[j]); k++ {
					if invertM[j][k] != invertM[i+i-j-1][k] {
						diff++
					}
				}
			}
			if diff == 1 {
				resM += i
				break
			}
		}
		if resM != 0 {
			res += resM
			continue
		}
	}

	return res
}

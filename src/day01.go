package src

import (
	"os"
	"strings"
)

var runesSetToDigits = map[int][]rune{
	1: {'o', 'n', 'e'},
	2: {'t', 'w', 'o'},
	3: {'t', 'h', 'r', 'e', 'e'},
	4: {'f', 'o', 'u', 'r'},
	5: {'f', 'i', 'v', 'e'},
	6: {'s', 'i', 'x'},
	7: {'s', 'e', 'v', 'e', 'n'},
	8: {'e', 'i', 'g', 'h', 't'},
	9: {'n', 'i', 'n', 'e'},
}

func (s *Service) CalculateDay01FirstPart() int {
	fileName := "./data/day01/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	res := 0
	for _, l := range lines {
		lineRunes := []rune(l)
		for _, r := range lineRunes {
			if r >= '0' && r <= '9' {
				res += s.runeToDigit(r) * 10
				break
			}
		}
		for i := len(lineRunes) - 1; i >= 0; i-- {
			if lineRunes[i] >= '0' && lineRunes[i] <= '9' {
				res += s.runeToDigit(lineRunes[i])
				break
			}
		}
	}

	return res
}

func (s *Service) CalculateDay01SecondPart() int {
	fileName := "./data/day01/input2"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	res := 0
	for _, l := range lines {
		lineRunes := []rune(l)
		for i, r := range lineRunes {
			if r >= '0' && r <= '9' {
				res += s.runeToDigit(r) * 10
				break
			}
			digit := s.getDigit(lineRunes, i, s.checkIsDigitDirect)
			if digit != 0 {
				res += digit * 10
				break
			}
		}
		for i := len(lineRunes) - 1; i >= 0; i-- {
			if lineRunes[i] >= '0' && lineRunes[i] <= '9' {
				res += s.runeToDigit(lineRunes[i])
				break
			}
			digit := s.getDigit(lineRunes, i, s.checkIsDigitReverse)
			if digit != 0 {
				res += digit
				break
			}
		}
	}

	return res
}

func (s *Service) getDigit(input []rune, fromIndex int, checker func(digitRunes []rune, input []rune, fromIndex int) bool) int {
	for j, digitRunes := range runesSetToDigits {
		if checker(digitRunes, input, fromIndex) {
			return j
		}
	}
	return 0
}

func (s *Service) checkIsDigitDirect(digitRunes []rune, input []rune, fromIndex int) bool {
	for k := 0; k < len(digitRunes); k++ {
		if digitRunes[k] != input[fromIndex+k] {
			return false
		}
	}
	return true
}

func (s *Service) checkIsDigitReverse(digitRunes []rune, input []rune, fromIndex int) bool {
	for k := len(digitRunes) - 1; k >= 0; k-- {
		if digitRunes[k] != input[fromIndex+k-len(digitRunes)+1] {
			return false
		}
	}
	return true
}

func (s *Service) runeToDigit(r rune) int {
	switch r {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return 100500
}

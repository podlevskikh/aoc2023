package src

import (
	"os"
	"strconv"
	"strings"
)

func (s *Service) CalculateDay12FirstPart() int {
	fileName := "./data/day12/input1"
	file, _ := os.ReadFile(fileName)

	lines := strings.Split(string(file), "\n")

	res := 0
	for _, l := range lines {
		par := strings.Split(l, " ")
		str := par[0]
		nums := strings.Split(par[1], ",")
		vars := map[string]int{str: 1}
		allNum := 0
		for _, n := range nums {
			num, _ := strconv.Atoi(n)
			allNum += num
		}
		minLen := allNum + len(nums) - 1
		for _, n := range nums {
			num, _ := strconv.Atoi(n)
			allNum -= num
			minLen -= num + 1
			nextVars := map[string]int{}
			for next, prevs := range vars {
				for _, newVar := range s.getAllVars(num, next) {
					_, ne := s.cutByIndex(1, newVar)
					if s.checkPossibility(allNum, minLen, ne) {
						if _, ok := nextVars[ne]; !ok {
							nextVars[ne] = 0
						}
						nextVars[ne] += prevs
					}
				}
			}
			vars = nextVars
		}
		for next, prevs := range vars {
			isOk := true
			for _, r := range strings.Split(next, "") {
				if r == "#" {
					isOk = false
				}
			}
			if isOk {
				res += prevs

			}
		}
	}

	return res
}

func (s *Service) checkPossibility(all, minLen int, variant string) bool {
	varChars := strings.Split(variant, "")
	if len(varChars) < minLen {
		return false
	}
	allPossible := 0
	for _, c := range varChars {
		if c == "." {
			continue
		}
		allPossible++
	}
	return allPossible >= all
}

func (s *Service) getAllVars(num int, variant string) []string {
	variants := []string{variant}
	finVars := []string{}
	for {
		newVars := []string{}
		for _, v := range variants {
			for _, varCandidate := range s.getNewPossibilities(v) {
				if s.isFin(num, varCandidate) {
					finVars = append(finVars, varCandidate)
				} else if s.isFit(num, varCandidate) {
					newVars = append(newVars, varCandidate)
				}
			}
		}
		if len(newVars) == 0 {
			return finVars
		}
		variants = newVars
	}
}

func (s *Service) cutByIndex(i int, v string) (string, string) {
	k := 0
	prev := []string{}
	next := []string{}
	for _, el := range strings.Split(v, ".") {
		if k >= i {
			next = append(next, el)
		} else {
			prev = append(prev, el)
			if el != "" {
				k++
			}
		}
	}
	resPrev := strings.Join(prev, ".")
	if len(prev) > 0 && len(next) > 0 {
		resPrev += "."
	}
	return resPrev, strings.Join(next, ".")
}

func (s *Service) getNewPossibilities(v string) []string {
	for i, r := range []rune(v) {
		if r == '?' {
			res1 := []rune(v)
			res2 := []rune(v)
			res1[i] = '.'
			res2[i] = '#'
			return []string{string(res1), string(res2)}
		}
	}
	return []string{v}
}

func (s *Service) isFit(num int, variant string) bool {
	isIn := false
	countCurrent := 0

	for _, e := range strings.Split(variant, "") {
		if e == "?" {
			return true
		}
		if e == "." {
			if isIn {
				return countCurrent == num
			} else {
				continue
			}
		}
		if e == "#" {
			isIn = true
			countCurrent++
		}
	}
	return countCurrent == num
}

func (s *Service) isFin(num int, variant string) bool {
	isIn := false
	countCurrent := 0
	for _, e := range strings.Split(variant, "") {
		if e == "?" {
			return false
		}
		if e == "." {
			if isIn {
				return countCurrent == num
			} else {
				continue
			}
		}
		if e == "#" {
			isIn = true
			countCurrent++
		}
	}
	return countCurrent == num
}

func (s *Service) CalculateDay12SecondPart() int {
	fileName := "./data/day12/input1"
	file, _ := os.ReadFile(fileName)

	lines := strings.Split(string(file), "\n")

	res := 0
	for _, l := range lines {
		par := strings.Split(l, " ")
		str := par[0]
		nums := strings.Split(par[1]+","+par[1]+","+par[1]+","+par[1]+","+par[1], ",")
		vars := map[string]int{str + "?" + str + "?" + str + "?" + str + "?" + str: 1}
		allNum := 0
		for _, n := range nums {
			num, _ := strconv.Atoi(n)
			allNum += num
		}
		minLen := allNum + len(nums) - 1
		for _, n := range nums {
			num, _ := strconv.Atoi(n)
			allNum -= num
			minLen -= num + 1
			nextVars := map[string]int{}
			for next, prevs := range vars {
				for _, newVar := range s.getAllVars(num, next) {
					_, ne := s.cutByIndex(1, newVar)
					if s.checkPossibility(allNum, minLen, ne) {
						if _, ok := nextVars[ne]; !ok {
							nextVars[ne] = 0
						}
						nextVars[ne] += prevs
					}
				}
			}
			vars = nextVars
		}
		for next, prevs := range vars {
			isOk := true
			for _, r := range strings.Split(next, "") {
				if r == "#" {
					isOk = false
				}
			}
			if isOk {
				res += prevs

			}
		}
	}

	return res
}

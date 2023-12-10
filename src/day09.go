package src

import (
	"os"
	"strconv"
	"strings"
)

func (s *Service) CalculateDay09FirstPart() int {
	fileName := "./data/day09/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	res := 0
	for _, l := range lines {
		c := strings.Split(l, " ")
		seq := make([]int, len(c))
		for i, k := range c {
			seq[i], _ = strconv.Atoi(k)
		}

		res += s.findNext(seq)
	}

	return res
}

func (s *Service) findNext(seq []int) int {
	allZero := true
	for _, se := range seq {
		if se != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}
	nextSeq := make([]int, len(seq) - 1)
	for i := 0; i < len(seq) - 1; i++ {
		nextSeq[i] = seq[i+1]-seq[i]
	}

	return s.findNext(nextSeq) + seq[len(seq) - 1]
}

func (s *Service) findPrev(seq []int) int {
	allZero := true
	for _, se := range seq {
		if se != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}
	nextSeq := make([]int, len(seq) - 1)
	for i := 0; i < len(seq) - 1; i++ {
		nextSeq[i] = seq[i+1]-seq[i]
	}

	return seq[0] - s.findPrev(nextSeq)
}

func (s *Service) CalculateDay09SecondPart() int {
	fileName := "./data/day09/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	res := 0
	for _, l := range lines {
		c := strings.Split(l, " ")
		seq := make([]int, len(c))
		for i, k := range c {
			seq[i], _ = strconv.Atoi(k)
		}

		res += s.findPrev(seq)
	}

	return res
}

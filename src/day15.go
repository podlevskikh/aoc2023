package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (s *Service) CalculateDay15FirstPart() int {
	fileName := "./data/day15/input1"
	file, _ := os.ReadFile(fileName)

	inputs := strings.Split(string(file), ",")
	res := 0
	for _, input := range inputs {
		inputRes := 0
		for _, i := range []rune(input) {
			inputRes += int(i)
			inputRes *= 17
			inputRes = inputRes % 256
		}
		res += inputRes
	}

	return res
}



func (s *Service) CalculateDay15SecondPart() int {
	fileName := "./data/day15/input1"
	file, _ := os.ReadFile(fileName)

	inputs := strings.Split(string(file), ",")
	boxes := make([][][]string, 256)
	res := 0
	for _, input := range inputs {
		boxId := s.determineBox(input)

		if input[len(input)-1:] == "-" {
			for i, boxInput := range boxes[boxId] {
				if len(boxInput) == 0 {
					continue
				}
				if boxInput[0] == input[:len(input)-1] {
					boxes[boxId][i] = []string{}
				}
			}
		} else {
			inputArr := strings.Split(input, "=")
			if inputArr[0] == "sth" {
				fmt.Println("start", input, inputArr, boxes[boxId], boxId)
			}
			contains := false
			for i, boxInput := range boxes[boxId] {
				if len(boxInput) == 0 {
					continue
				}
				if boxInput[0] == inputArr[0] {
					boxes[boxId][i] = inputArr
					contains = true
				}
			}
			if inputArr[0] == "sth" {
				fmt.Println("finish", input, inputArr, boxes[boxId], boxId)
			}

			if !contains {
				boxes[boxId] = append(boxes[boxId], inputArr)
			}
		}
	}

	for i, box := range boxes {
		boxSlot := 1
		for _, lens := range box {
			if len(lens) == 0 {
				continue
			}
			n, _ := strconv.Atoi(lens[1])
			res += (i+1)*n*boxSlot
			boxSlot++
		}
	}

	return res
}

func (s *Service) determineBox(input string) int {
	inputRes := 0
	for _, i := range []rune(input) {
		if i == '=' || i == '-' {
			break
		}
		inputRes += int(i)
		inputRes *= 17
		inputRes = inputRes % 256
	}
	return inputRes
}

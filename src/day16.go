package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Beam struct {
	X int
	Y int
	Direction string
	Status string
}

func (b Beam) Key() string {
	return fmt.Sprintf("%v_%v_%s", strconv.Itoa(b.X), strconv.Itoa(b.Y), b.Direction)
}

func (s *Service) CalculateDay16FirstPart() int {
	fileName := "./data/day16/input1"
	file, _ := os.ReadFile(fileName)

	inputs := strings.Split(string(file), "\n")
	space := make([][]rune, len(inputs))
	for j, i := range inputs {
		space[j] = []rune(i)
	}

	startBeam := &Beam{
		X:         0,
		Y:         0,
		Direction: "right",
		Status:    "ok",
	}
	return s.calcEnergy(startBeam, space)
}

func (s *Service) calcEnergy(startBeam *Beam, space [][]rune) int {
	res := make([][]int, len(space))
	for j, i := range space {
		res[j] = make([]int, len(i))
	}
	beams := map[string]*Beam{startBeam.Key():startBeam}

	for {
		notAllTerminated := false
		for _, b := range beams {
			if b.Status == "terminated" {
				continue
			}
			if b.X < 0 || b.X >= len(space[0]) || b.Y < 0 || b.Y >= len(space) {
				b.Status = "terminated"
				continue
			}
			notAllTerminated = true
			res[b.X][b.Y] = 1
			if space[b.Y][b.X] == '.' || (space[b.Y][b.X] == '-' && (b.Direction == "right" || b.Direction == "left")) ||
				(space[b.Y][b.X] == '|' && (b.Direction == "top" || b.Direction == "bottom")){
				switch b.Direction {
				case "right":
					b.X++
					break
				case "left":
					b.X--
					break
				case "top":
					b.Y--
					break
				case "bottom":
					b.Y++
					break
				}
			} else if space[b.Y][b.X] == '\\' && b.Direction == "right" || space[b.Y][b.X] == '/' && b.Direction == "left" {
				b.Status = "terminated"
				newBeam := Beam{
					X:         b.X,
					Y:         b.Y+1,
					Direction: "bottom",
					Status:    "ok",
				}
				if _, ok := beams[newBeam.Key()]; !ok {
					beams[newBeam.Key()] = &newBeam
				}
			} else if space[b.Y][b.X] == '/' && b.Direction == "right" || space[b.Y][b.X] == '\\' && b.Direction == "left" {
				b.Status = "terminated"
				newBeam := Beam{
					X:         b.X,
					Y:         b.Y-1,
					Direction: "top",
					Status:    "ok",
				}
				if _, ok := beams[newBeam.Key()]; !ok {
					beams[newBeam.Key()] = &newBeam
				}
			} else if space[b.Y][b.X] == '/' && b.Direction == "top" || space[b.Y][b.X] == '\\' && b.Direction == "bottom" {
				b.Status = "terminated"
				newBeam := Beam{
					X:         b.X+1,
					Y:         b.Y,
					Direction: "right",
					Status:    "ok",
				}
				if _, ok := beams[newBeam.Key()]; !ok {
					beams[newBeam.Key()] = &newBeam
				}
			} else if space[b.Y][b.X] == '/' && b.Direction == "bottom" || space[b.Y][b.X] == '\\' && b.Direction == "top" {
				b.Status = "terminated"
				newBeam := Beam{
					X:         b.X-1,
					Y:         b.Y,
					Direction: "left",
					Status:    "ok",
				}
				if _, ok := beams[newBeam.Key()]; !ok {
					beams[newBeam.Key()] = &newBeam
				}
			} else if space[b.Y][b.X] == '-' && (b.Direction == "bottom" || b.Direction == "top") {
				b.Status = "terminated"
				newBeamRight := Beam{
					X:         b.X+1,
					Y:         b.Y,
					Direction: "right",
					Status:    "ok",
				}
				if _, ok := beams[newBeamRight.Key()]; !ok {
					beams[newBeamRight.Key()] = &newBeamRight
				}
				newBeamLeft := Beam{
					X:         b.X-1,
					Y:         b.Y,
					Direction: "left",
					Status:    "ok",
				}
				if _, ok := beams[newBeamLeft.Key()]; !ok {
					beams[newBeamLeft.Key()] = &newBeamLeft
				}
			} else if space[b.Y][b.X] == '|' && (b.Direction == "left" || b.Direction == "right") {
				b.Status = "terminated"
				newBeamTop := Beam{
					X:         b.X,
					Y:         b.Y-1,
					Direction: "top",
					Status:    "ok",
				}
				if _, ok := beams[newBeamTop.Key()]; !ok {
					beams[newBeamTop.Key()] = &newBeamTop
				}
				newBeamBottom := Beam{
					X:         b.X,
					Y:         b.Y+1,
					Direction: "bottom",
					Status:    "ok",
				}
				if _, ok := beams[newBeamBottom.Key()]; !ok {
					beams[newBeamBottom.Key()] = &newBeamBottom
				}
			}
		}
		if !notAllTerminated {
			break
		}
	}

	resSum := 0
	for _, l := range res {
		for _, c := range l {
			resSum+=c
		}
	}

	return resSum
}

func (s *Service) CalculateDay16SecondPart() int {
	fileName := "./data/day16/input1"
	file, _ := os.ReadFile(fileName)

	inputs := strings.Split(string(file), "\n")
	space := make([][]rune, len(inputs))
	for j, i := range inputs {
		space[j] = []rune(i)
	}

	maxRes := 0
	for i := 0; i < len(space[0]); i++ {
		startBeamTop := &Beam{
			X:         i,
			Y:         0,
			Direction: "bottom",
			Status:    "ok",
		}
		res := s.calcEnergy(startBeamTop, space)
		if res > maxRes {
			maxRes = res
		}
		startBeamBottom := &Beam{
			X:         i,
			Y:         len(space) - 1,
			Direction: "top",
			Status:    "ok",
		}
		res = s.calcEnergy(startBeamBottom, space)
		if res > maxRes {
			maxRes = res
		}
	}
	for i := 0; i < len(space); i++ {
		startBeamLeft := &Beam{
			X:         0,
			Y:         i,
			Direction: "right",
			Status:    "ok",
		}
		res := s.calcEnergy(startBeamLeft, space)
		if res > maxRes {
			maxRes = res
		}
		startBeamRight := &Beam{
			X:         len(space[0]) - 1,
			Y:         i,
			Direction: "left",
			Status:    "ok",
		}
		res = s.calcEnergy(startBeamRight, space)
		if res > maxRes {
			maxRes = res
		}
	}
	return maxRes
}

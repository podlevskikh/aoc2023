package src

import (
	"os"
	"strconv"
	"strings"
)

type MapPoint struct {
	X int
	Y int
}

func (s *Service) CalculateDay10FirstPart() int {
	fileName := "./data/day10/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	animalMap := map[int]map[int]string{}
	for i, l := range lines {
		animalMapLine := map[int]string{}
		for j, p := range strings.Split(l, "") {
			animalMapLine[j] = p
		}
		animalMap[i] = animalMapLine
	}
	startPoint := MapPoint{X: 95, Y: 74,}
	currentPoint := MapPoint{X: 94, Y: 74,}
	prevPoint := startPoint
	way := []MapPoint{startPoint, currentPoint}
	for {
		if animalMap[currentPoint.X][currentPoint.Y] == "S" {
			break
		}
		var nextPoint MapPoint
		switch animalMap[currentPoint.X][currentPoint.Y] {
		case "|":
			nextPoint = MapPoint{X: 2*currentPoint.X - prevPoint.X, Y: currentPoint.Y}
			break
		case "-":
			nextPoint = MapPoint{X: currentPoint.X, Y: 2*currentPoint.Y-prevPoint.Y}
			break
		case "L":
			nextPoint = MapPoint{X: currentPoint.X - prevPoint.Y + currentPoint.Y, Y: currentPoint.Y + currentPoint.X - prevPoint.X}
			break
		case "J":
			nextPoint = MapPoint{X: currentPoint.X + prevPoint.Y - currentPoint.Y, Y: currentPoint.Y - currentPoint.X + prevPoint.X}
			break
		case "7":
			nextPoint = MapPoint{X: currentPoint.X - prevPoint.Y + currentPoint.Y, Y: currentPoint.Y + currentPoint.X - prevPoint.X}
			break
		case "F":
			nextPoint = MapPoint{X: currentPoint.X + prevPoint.Y - currentPoint.Y, Y: currentPoint.Y - currentPoint.X + prevPoint.X}
			break
		}

		prevPoint = currentPoint
		currentPoint = nextPoint
		way = append(way, currentPoint)
	}

	return len(way) / 2
}

func (s *Service) CalculateDay10SecondPart() int {
	fileName := "./data/day10/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	animalMap := [][]string{}
	for _, l := range lines {
		animalMap = append(animalMap, strings.Split(l, ""))
	}
	startPoint := MapPoint{X: 95, Y: 74,}
	currentPoint := MapPoint{X: 94, Y: 74,}
	prevPoint := startPoint
	wayMap := map[string]struct{}{}
	wayMap[strconv.Itoa(startPoint.X)+"_"+strconv.Itoa(startPoint.Y)] = struct{}{}
	wayMap[strconv.Itoa(currentPoint.X)+"_"+strconv.Itoa(currentPoint.Y)] = struct{}{}
	for {
		if animalMap[currentPoint.X][currentPoint.Y] == "S" {
			break
		}
		var nextPoint MapPoint
		switch animalMap[currentPoint.X][currentPoint.Y] {
		case "|":
			nextPoint = MapPoint{X: 2*currentPoint.X - prevPoint.X, Y: currentPoint.Y}
			break
		case "-":
			nextPoint = MapPoint{X: currentPoint.X, Y: 2*currentPoint.Y-prevPoint.Y}
			break
		case "L":
			nextPoint = MapPoint{X: currentPoint.X - prevPoint.Y + currentPoint.Y, Y: currentPoint.Y + currentPoint.X - prevPoint.X}
			break
		case "J":
			nextPoint = MapPoint{X: currentPoint.X + prevPoint.Y - currentPoint.Y, Y: currentPoint.Y - currentPoint.X + prevPoint.X}
			break
		case "7":
			nextPoint = MapPoint{X: currentPoint.X - prevPoint.Y + currentPoint.Y, Y: currentPoint.Y + currentPoint.X - prevPoint.X}
			break
		case "F":
			nextPoint = MapPoint{X: currentPoint.X + prevPoint.Y - currentPoint.Y, Y: currentPoint.Y - currentPoint.X + prevPoint.X}
			break
		}

		prevPoint = currentPoint
		currentPoint = nextPoint
		wayMap[strconv.Itoa(currentPoint.X)+"_"+strconv.Itoa(currentPoint.Y)] = struct{}{}
	}

	animalMap[currentPoint.X][currentPoint.Y] = "L"
	in := false
	borderFromTop := false

	res := 0
	for i, l := range animalMap {
		for j, p := range l {
			if _, ok := wayMap[strconv.Itoa(i)+"_"+strconv.Itoa(j)]; !ok {
				if in {
					res++
				}
			} else {
				if p == "|" {
					in = !in
				} else if p == "L" {
					in = !in
					borderFromTop = true
				} else if p == "F" {
					in = !in
					borderFromTop = false
				} else if p == "7" {
					if !borderFromTop {
						in = !in
					}
				} else if p == "J" {
					if borderFromTop {
						in = !in
					}
				}
			}
		}
	}

	return res
}

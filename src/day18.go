package src

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func (s *Service) CalculateDay18FirstPart() int {
	fileName := "./data/day18/input1"
	file, _ := os.ReadFile(fileName)

	digMap := make([][]rune, 10000)
	for j := 0; j < 10000; j++ {
		digRow := make([]rune, 10000)
		for i := 0; i < 10000; i++ {
			digRow[i] = '.'
		}
		digMap[j] = digRow
	}

	inputs := strings.Split(string(file), "\n")
	curPosX := 5000
	curPosY := 5000
	digMap[curPosY][curPosX] = '<'
	for _, l := range inputs {
		parsedL := strings.Split(l, " ")
		steps, _ := strconv.Atoi(parsedL[1])
		var f func()
		var symbol rune
		switch parsedL[0] {
		case "U":
			symbol = '^'
			if digMap[curPosY][curPosX] == '>' {
				digMap[curPosY][curPosX] = 'J'
			} else if digMap[curPosY][curPosX] == '<' {
				digMap[curPosY][curPosX] = 'L'
			}
			f = func() {
				curPosY--
			}
		case "D":
			symbol = 'v'
			if digMap[curPosY][curPosX] == '>' {
				digMap[curPosY][curPosX] = '7'
			} else if digMap[curPosY][curPosX] == '<' {
				digMap[curPosY][curPosX] = 'F'
			}
			f = func() {
				curPosY++
			}
		case "L":
			symbol = '<'
			if digMap[curPosY][curPosX] == 'v' {
				digMap[curPosY][curPosX] = 'J'
			} else if digMap[curPosY][curPosX] == '^' {
				digMap[curPosY][curPosX] = '7'
			}

			f = func() {
				curPosX--
			}
		case "R":
			symbol = '>'
			if digMap[curPosY][curPosX] == 'v' {
				digMap[curPosY][curPosX] = 'L'
			} else if digMap[curPosY][curPosX] == '^' {
				digMap[curPosY][curPosX] = 'F'
			}
			f = func() {
				curPosX++
			}
		}
		for i := 0; i < steps; i++ {
			f()
			digMap[curPosY][curPosX] = symbol
		}

	}

	digMap[curPosY][curPosX] = '7'
	res := 0
	for i, l := range digMap {
		in := false
		fromTop := false
		for j, c := range l {
			if c == 'v' || c == '^' {
				in = !in
			}

			if c == 'L' {
				fromTop = true
			}
			if c == 'F' {
				fromTop = false
			}
			if c == '7' && fromTop || c == 'J' && !fromTop {
				in = !in
			}

			if in && c == '.' {
				digMap[i][j] = '#'
			}

			if digMap[i][j] != '.' {
				res++
			}
		}
	}

	return res
}

type BorderX struct {
	X     float64
	FromY float64
	ToY   float64
}

type BorderY struct {
	Y     float64
	FromX float64
	ToX   float64
}

type Square struct {
	FromX int64
	ToX   int64
	FromY int64
	ToY   int64
}

func (s *Service) CalculateDay18SecondPart() int64 {
	fileName := "./data/day18/input1"
	file, _ := os.ReadFile(fileName)

	curPosX := float64(0)
	curPosY := float64(0)

	xList := []float64{curPosX + 0.5, curPosX - 0.5}
	yList := []float64{curPosY + 0.5, curPosY - 0.5}
	bordersX := []BorderX{}
	bordersY := []BorderY{}

	biiiigNum := float64(1005001000)
	borderSumU := int64(0)
	borderSumD := int64(0)
	borderSumL := int64(0)
	borderSumR := int64(0)
	inputs := strings.Split(string(file), "\n")
	for _, l := range inputs {
		parsedL := strings.Split(l, " ")

		steps, _ := strconv.ParseInt(strings.Trim(parsedL[2], "(#)")[0:len(parsedL[2])-4], 16, 64)
		switch strings.Trim(parsedL[2], "(#)")[len(parsedL[2])-4] {
		case '3': //'U':
			borderSumU += int64(steps)
			bordersX = append(bordersX, BorderX{
				X:     curPosX,
				FromY: curPosY,
				ToY:   curPosY - float64(steps),
			})
			curPosY -= float64(steps)
			yList = append(yList, curPosY+0.5, curPosY-0.5)
		case '1': //'D':
			borderSumD += int64(steps)
			bordersX = append(bordersX, BorderX{
				X:     curPosX,
				FromY: curPosY,
				ToY:   curPosY + float64(steps),
			})
			curPosY += float64(steps)
			yList = append(yList, curPosY+0.5, curPosY-0.5)
		case '2': //"L":
			borderSumL += int64(steps)
			bordersY = append(bordersY, BorderY{
				Y:     curPosY,
				FromX: curPosX,
				ToX:   curPosX - float64(steps),
			})
			curPosX -= float64(steps)
			xList = append(xList, curPosX+0.5, curPosX-0.5)
		case '0': //"R":
			borderSumR += int64(steps)
			bordersY = append(bordersY, BorderY{
				Y:     curPosY,
				FromX: curPosX,
				ToX:   curPosX + float64(steps),
			})
			curPosX += float64(steps)
			xList = append(xList, curPosX+0.5, curPosX-0.5)
		}
	}

	leftSum := int64(0)
	rightSum := int64(0)

	sort.Slice(xList, func(i, j int) bool {
		return xList[i] < xList[j]
	})
	sort.Slice(yList, func(i, j int) bool {
		return yList[i] < yList[j]
	})

	prevX := float64(-biiiigNum)
	prevY := float64(-biiiigNum)
	for _, x := range xList {
		if prevX == float64(-biiiigNum) || prevX == x {
			prevX = x
			continue
		}
		for _, y := range yList {
			if prevY == float64(-biiiigNum) || prevY == y {
				prevY = y
				continue
			}
			nearestLeftBorder := BorderX{X: float64(-biiiigNum)}
			onBorder := false
			for _, b := range bordersX {
				if b.FromY > y && b.ToY > y || b.FromY < prevY && b.ToY < prevY {
					continue
				}
				if b.X > prevX && b.X < x {
					onBorder = true
					break
				}
				if nearestLeftBorder.X < b.X && b.X < prevX {
					nearestLeftBorder = b
				}
			}
			for _, b := range bordersY {
				if b.FromX > x && b.ToX > x || b.FromX < prevX && b.ToX < prevX {
					continue
				}
				if b.Y > prevY && b.Y < y {
					onBorder = true
					break
				}
			}
			if onBorder || nearestLeftBorder.X == float64(-biiiigNum) {
				prevY = y
				continue
			}
			if nearestLeftBorder.ToY < nearestLeftBorder.FromY {
				leftSum += int64(x-prevX) * int64(y-prevY)
			} else {
				rightSum += int64(x-prevX) * int64(y-prevY)
			}
			prevY = y
		}
		prevX = x
	}

	return leftSum + borderSumU + borderSumD + borderSumL + borderSumR
}

package src

import (
	"math"
	"os"
	"strings"
)

func (s *Service) CalculateDay11FirstPart() int {
	fileName := "./data/day11/input1"
	file, _ := os.ReadFile(fileName)

	lines := strings.Split(string(file), "\n")

	originSpace := [][]string{}
	for _, l := range lines {
		spaceRow := []string{}
		for _, p := range strings.Split(l, "") {
			spaceRow = append(spaceRow, p)
		}
		originSpace = append(originSpace, spaceRow)
	}

	emptyColls := map[int]struct{}{}
	emptyRows := map[int]struct{}{}

	for i := range originSpace {
		empty := true
		for j := range originSpace[0] {
			if originSpace[i][j] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyRows[i] = struct{}{}
		}
	}

	for j := range originSpace[0] {
		empty := true
		for i := range originSpace {
			if originSpace[i][j] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyColls[j] = struct{}{}
		}
	}

	wideSpace := [][]string{}
	for i := range originSpace {
		row := []string{}
		for j, c := range originSpace[i] {
			row = append(row, c)
			if _, ok := emptyColls[j]; ok {
				row = append(row, c)
			}
		}
		wideSpace = append(wideSpace, row)
		if _, ok := emptyRows[i]; ok {
			wideSpace = append(wideSpace, row)
		}
	}

	rows := []float64{}
	cols := []float64{}
	for i := range wideSpace {
		for j, c := range wideSpace[i] {
			if c == "#" {
				rows = append(rows, float64(i))
				cols = append(cols, float64(j))
			}
		}
	}

	res := float64(0)
	for _, r1 := range rows {
		for _, r2 := range rows {
			res += math.Abs(r1 - r2)
		}
	}
	for _, c1 := range cols {
		for _, c2 := range cols {
			res += math.Abs(c1 - c2)
		}
	}

	return int(res / 2)
}

func (s *Service) CalculateDay11SecondPart() int {
	fileName := "./data/day11/input1"
	file, _ := os.ReadFile(fileName)

	lines := strings.Split(string(file), "\n")

	originSpace := [][]string{}
	for _, l := range lines {
		spaceRow := []string{}
		for _, p := range strings.Split(l, "") {
			spaceRow = append(spaceRow, p)
		}
		originSpace = append(originSpace, spaceRow)
	}

	emptyColls := map[int]struct{}{}
	emptyRows := map[int]struct{}{}

	for i := range originSpace {
		empty := true
		for j := range originSpace[0] {
			if originSpace[i][j] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyRows[i] = struct{}{}
		}
	}

	for j := range originSpace[0] {
		empty := true
		for i := range originSpace {
			if originSpace[i][j] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyColls[j] = struct{}{}
		}
	}

	rows := []float64{}
	cols := []float64{}
	for i := range originSpace {
		for j, c := range originSpace[i] {
			if c == "#" {
				rows = append(rows, float64(i))
				cols = append(cols, float64(j))
			}
		}
	}

	for i, c := range cols {
		for ec := range emptyColls {
			if int(c) > ec {
				cols[i] += 1000000 - 1
			}
		}
	}

	for i, r := range rows {
		for er := range emptyRows {
			if int(r) > er {
				rows[i] += 1000000 - 1
			}
		}
	}

	res := float64(0)
	for _, r1 := range rows {
		for _, r2 := range rows {
			res += math.Abs(r1 - r2)
		}
	}
	for _, c1 := range cols {
		for _, c2 := range cols {
			res += math.Abs(c1 - c2)
		}
	}

	return int(res / 2)
}

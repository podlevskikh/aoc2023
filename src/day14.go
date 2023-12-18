package src

import (
	"os"
	"strings"
)

func (s *Service) CalculateDay14FirstPart() int {
	fileName := "./data/day14/input1"
	file, _ := os.ReadFile(fileName)

	field := [][]rune{}
	for _, l := range strings.Split(string(file), "\n") {
		field = append(field, []rune(l))
	}
	size := len(field)
	res := 0
	for i, l := range field {
		for j, c := range l {
			if c == '.' || c == '#' {
				continue
			}
			k := i
			for {
				if k == 0 || field[k-1][j] == 'O' || field[k-1][j] == '#' {
					res += size - k
					break
				}
				field[k][j] = field[k-1][j]
				field[k-1][j] = 'O'
				k = k -1
			}
		}
	}

	return res
}

func (s *Service) rollField(field [][]rune) [][]rune {
	//north
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] == '.' || field[i][j] == '#' {
				continue
			}
			k := i
			for {
				if k == 0 || field[k-1][j] == 'O' || field[k-1][j] == '#' {
					break
				}
				field[k][j] = field[k-1][j]
				field[k-1][j] = 'O'
				k--
			}
		}
	}

	//west
	for j := 0; j < len(field[0]); j++ {
		for i := 0; i < len(field); i++ {
			if field[i][j] == '.' || field[i][j] == '#' {
				continue
			}
			k := j
			for {
				if k == 0 || field[i][k-1] == 'O' || field[i][k-1] == '#' {
					break
				}
				field[i][k] = field[i][k-1]
				field[i][k-1] = 'O'
				k--
			}
		}
	}

	//south
	for i := len(field) - 1; i >= 0; i-- {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] == '.' || field[i][j] == '#' {
				continue
			}
			k := i
			for {
				if k == len(field) - 1 || field[k+1][j] == 'O' || field[k+1][j] == '#' {
					break
				}
				field[k][j] = field[k+1][j]
				field[k+1][j] = 'O'
				k++
			}
		}
	}

	//east
	for j := len(field[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(field); i++ {
			if field[i][j] == '.' || field[i][j] == '#' {
				continue
			}
			k := j
			for {
				if k == len(field[0]) - 1 || field[i][k+1] == 'O' || field[i][k+1] == '#' {
					break
				}
				field[i][k] = field[i][k+1]
				field[i][k+1] = 'O'
				k++
			}
		}
	}
	return field
}

func (s *Service) calcField(field [][]rune) int {
	size := len(field)
	res := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] == 'O' {
				res += size - i
			}
		}
	}
	return res
}

func (s *Service) CalculateDay14SecondPart() int {
	fileName := "./data/day14/input1"
	file, _ := os.ReadFile(fileName)

	field := [][]rune{}
	for _, l := range strings.Split(string(file), "\n") {
		field = append(field, []rune(l))
	}

	fieldHashMap := map[string][]int{}

	step := 0
	begin := 0
	res := 0
	for i := 0; i<1000000000; i++ {
		sum := ""
		for _, l := range field {
			sum += string(l)
		}
		if j, ok := fieldHashMap[sum]; ok {
			step = i - j[0]
			begin = j[0]
			need := (1000000000 - begin) % step + begin

			for _, q := range fieldHashMap {
				if q[0] == need {
					res = q[1]
					break
				}
			}

			break
		}

		fieldHashMap[sum] = []int{i, s.calcField(field)}
		field = s.rollField(field)
	}


	return res
}

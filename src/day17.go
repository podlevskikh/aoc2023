package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LavaStep struct {
	X int
	Y int
	Lost int
	Direction rune
	StepsInOneDirection int
}

func (ls LavaStep) Hash() string {
	return fmt.Sprintf("%v_%v_%s_%v", ls.X, ls.Y, string(ls.Direction), ls.StepsInOneDirection)
	//return fmt.Sprintf("%v_%v", ls.X, ls.Y)
}

func (s *Service) CalculateDay17FirstPart() int {
	fileName := "./data/day17/input1"
	file, _ := os.ReadFile(fileName)
//file := "2413432311323\n3215453535623\n3255245654254\n3446585845452\n4546657867536\n1438598798454\n4457876987766\n3637877979653\n4654967986887\n4564679986453\n1224686865563\n2546548887735\n4322674655533"
	inputs := strings.Split(string(file), "\n")
	space := make([][]int, len(inputs))
	for j, l := range inputs {
		space[j] = make([]int, len(l))
		for i, c := range strings.Split(l, "") {
			space[j][i], _ = strconv.Atoi(c)
		}
	}

	waysMap := map[string]int{}
	activeSteps := []LavaStep{{
		X:                   1,
		Y:                   0,
		Lost:                space[0][1],
		Direction:           '>',
		StepsInOneDirection: 1,
	},{
		X:                   0,
		Y:                   1,
		Lost:                space[1][0],
		Direction:           'v',
		StepsInOneDirection: 1,
	}}

	for ;len(activeSteps) > 0; {
		newActiveSteps := make([]LavaStep, 0, len(activeSteps)*3)
		for _, as := range activeSteps {
			for _, d := range []rune{'>', '<', 'v', '^'} {
				if as.X == 0 && d == '<' ||
					as.Y == 0 && d == '^' ||
					as.Y == len(space)-1 && d == 'v' ||
					as.X == len(space[0])-1 && d == '>' ||
					d == as.Direction && as.StepsInOneDirection == 3 ||
					as.Direction == '>' && d == '<' ||
					as.Direction == '<' && d == '>' ||
					as.Direction == 'v' && d == '^' ||
					as.Direction == '^' && d == 'v' {
					continue
				}
				nas := LavaStep{
					X:                   as.X,
					Y:                   as.Y,
					Lost:                as.Lost,
					Direction:           d,
					StepsInOneDirection: as.StepsInOneDirection+1,
				}
				switch d {
				case '>':
					nas.X++
				case '<':
					nas.X--
				case 'v':
					nas.Y++
				case '^':
					nas.Y--
				}
				if d != as.Direction {
					nas.StepsInOneDirection=1
				}
				nas.Lost+=space[nas.Y][nas.X]

				hash := nas.Hash()
				if nas.Y == len(space)-1 && nas.X == len(space[0])-1 {
					hash = "fin"
				}
				if lost, ok := waysMap[hash]; !ok || lost > nas.Lost {
					waysMap[hash] = nas.Lost
					newActiveSteps = append(newActiveSteps, nas)
				}
			}
		}
		activeSteps = newActiveSteps
	}

	return waysMap["fin"]
}

func (s *Service) CalculateDay17SecondPart() int {
	fileName := "./data/day17/input1"
	file, _ := os.ReadFile(fileName)
	//file := "111111111111\n999999999991\n999999999991\n999999999991\n999999999991"
	inputs := strings.Split(string(file), "\n")
	space := make([][]int, len(inputs))
	for j, l := range inputs {
		space[j] = make([]int, len(l))
		for i, c := range strings.Split(l, "") {
			space[j][i], _ = strconv.Atoi(c)
		}
	}

	waysMap := map[string]int{}
	activeSteps := []LavaStep{{
		X:                   1,
		Y:                   0,
		Lost:                space[0][1],
		Direction:           '>',
		StepsInOneDirection: 1,
	},{
		X:                   0,
		Y:                   1,
		Lost:                space[1][0],
		Direction:           'v',
		StepsInOneDirection: 1,
	}}

	for ;len(activeSteps) > 0; {
		newActiveSteps := make([]LavaStep, 0, len(activeSteps)*3)
		for _, as := range activeSteps {
			for _, d := range []rune{'>', '<', 'v', '^'} {
				if as.X == 0 && d == '<' ||
					as.Y == 0 && d == '^' ||
					as.Y == len(space)-1 && d == 'v' ||
					as.X == len(space[0])-1 && d == '>' ||
					d == as.Direction && as.StepsInOneDirection == 10 ||
					d != as.Direction && as.StepsInOneDirection < 4 ||
					as.Direction == '>' && d == '<' ||
					as.Direction == '<' && d == '>' ||
					as.Direction == 'v' && d == '^' ||
					as.Direction == '^' && d == 'v' {
					continue
				}
				nas := LavaStep{
					X:                   as.X,
					Y:                   as.Y,
					Lost:                as.Lost,
					Direction:           d,
					StepsInOneDirection: as.StepsInOneDirection+1,
				}
				switch d {
				case '>':
					nas.X++
				case '<':
					nas.X--
				case 'v':
					nas.Y++
				case '^':
					nas.Y--
				}
				if d != as.Direction {
					nas.StepsInOneDirection=1
				}
				nas.Lost+=space[nas.Y][nas.X]

				hash := nas.Hash()
				if nas.Y == len(space)-1 && nas.X == len(space[0])-1 {
					hash = "fin"
				}
				if hash == "fin" && nas.StepsInOneDirection < 4{
					continue
				}
				if lost, ok := waysMap[hash]; !ok || lost > nas.Lost {
					waysMap[hash] = nas.Lost
					newActiveSteps = append(newActiveSteps, nas)
				}
			}
		}
		activeSteps = newActiveSteps
	}

	return waysMap["fin"]
}

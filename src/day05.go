package src

import (
	"os"
	"strconv"
	"strings"
)

func (s *Service) CalculateDay05FirstPart() int {
	seedsLine := "4043382508 113348245 3817519559 177922221 3613573568 7600537 773371046 400582097 2054637767 162982133 2246524522 153824596 1662955672 121419555 2473628355 846370595 1830497666 190544464 230006436 483872831"
	seeds := make([]int, 0)
	for _, seed := range strings.Split(seedsLine, " ") {
		s, _ := strconv.Atoi(seed)
		seeds = append(seeds, s)
	}

	fileName := "./data/day05/input1"
	file, _ := os.ReadFile(fileName)
	maps := strings.Split(string(file), "\n\n")

	for _, m := range maps {
		mapsLines := strings.Split(m, "\n")
		mapper := map[int][]int{}
		for i := 1; i < len(mapsLines); i++ {
			line := strings.Split(mapsLines[i], " ")
			to, _ := strconv.Atoi(line[0])
			from, _ := strconv.Atoi(line[1])
			dur, _ := strconv.Atoi(line[2])
			mapper[from] = []int{to, dur}
		}

		newSeeds := []int{}
		for _, s := range seeds {
			find := false
			for from, todur := range mapper {
				if s >= from && s < from+todur[1] {
					newSeeds = append(newSeeds, todur[0]+s-from)
					find = true
					break
				}
			}
			if !find {
				newSeeds = append(newSeeds, s)
			}
		}
		seeds = newSeeds
	}

	res := -1
	for _, s := range seeds {
		if res == -1 || res > s {
			res = s
		}
	}

	return res
}

func (s *Service) CalculateDay05SecondPart() int {
	seedsLine := "4043382508 113348245 3817519559 177922221 3613573568 7600537 773371046 400582097 2054637767 162982133 2246524522 153824596 1662955672 121419555 2473628355 846370595 1830497666 190544464 230006436 483872831"
	seeds := make([][]int, 0)
	seedssplit := strings.Split(seedsLine, " ")
	for i := 0; i < len(seedssplit); i += 2 {
		seed, _ := strconv.Atoi(seedssplit[i])
		dur, _ := strconv.Atoi(seedssplit[i+1])
		seeds = append(seeds, []int{seed, dur})
	}

	fileName := "./data/day05/input1"
	file, _ := os.ReadFile(fileName)
	maps := strings.Split(string(file), "\n\n")

	for _, m := range maps {
		mapsLines := strings.Split(m, "\n")
		mapper := map[int][]int{}
		for i := 1; i < len(mapsLines); i++ {
			line := strings.Split(mapsLines[i], " ")
			to, _ := strconv.Atoi(line[0])
			from, _ := strconv.Atoi(line[1])
			dur, _ := strconv.Atoi(line[2])
			mapper[from] = []int{to, dur}
		}

		newSeeds := [][]int{}
		for _, originSeedDur := range seeds {
			seedDursToFind := [][]int{originSeedDur}
			for from, todur := range mapper {
				newSeedDursToFind := [][]int{}
				for _, seedDur := range seedDursToFind {
					if seedDur[0]+seedDur[1]-1 < from || seedDur[0] > from+todur[1]-1 {
						newSeedDursToFind = append(newSeedDursToFind, seedDur)
					} else if seedDur[0] < from && seedDur[0]+seedDur[1]-1 >= from && seedDur[0]+seedDur[1]-1 <= from+todur[1]-1 {
						newSeeds = append(newSeeds, []int{todur[0], seedDur[0]+seedDur[1]-from})
						newSeedDursToFind = append(newSeedDursToFind, []int{seedDur[0], from-seedDur[0]})
					} else if seedDur[0] < from && seedDur[0]+seedDur[1]-1 > from+todur[1]-1 {
						newSeeds = append(newSeeds, []int{todur[0], todur[1]})
						newSeedDursToFind = append(newSeedDursToFind, []int{seedDur[0], from-seedDur[0]})
						newSeedDursToFind = append(newSeedDursToFind, []int{from+todur[1], seedDur[0]+seedDur[1]-from-todur[1]})
					} else if seedDur[0] >= from && seedDur[0]+seedDur[1]-1 <= from+todur[1]-1 {
						newSeeds = append(newSeeds, []int{todur[0]+seedDur[0]-from, seedDur[1]})
					} else if seedDur[0] >= from && seedDur[0] <= from+todur[1]-1 && seedDur[0]+seedDur[1]-1 > from+todur[1]-1 {
						newSeeds = append(newSeeds, []int{todur[0]+seedDur[0]-from, from+todur[1]-seedDur[0]})
						newSeedDursToFind = append(newSeedDursToFind, []int{from+todur[1], seedDur[0]+seedDur[1]-from-todur[1]})
					}
				}
				seedDursToFind = newSeedDursToFind
				if len(seedDursToFind) == 0 {
					break
				}
			}

			for _, sf := range seedDursToFind {
				newSeeds = append(newSeeds, sf)
			}
		}
		seeds = newSeeds
	}

	res := -1
	for _, s := range seeds {
		if res == -1 || res > s[0] {
			res = s[0]
		}
	}

	return res
}

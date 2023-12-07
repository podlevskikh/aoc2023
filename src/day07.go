package src

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

var CardWeight = map[string]int {
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

const (
	ComboFive = "five"
	ComboFour = "four"
	ComboFull = "full"
	ComboThree = "three"
	ComboTwo = "two"
	ComboOne = "one"
	ComboHigh = "high"
)

var ComboWeight = map[string]int{
	ComboFive: 7,
	ComboFour: 6,
	ComboFull: 5,
	ComboThree: 4,
	ComboTwo: 3,
	ComboOne: 2,
	ComboHigh: 1,
}

type PokerStruct struct {
	Cards []string
	Combo string
	Rank int
}

func (s *Service) CalculateDay07FirstPart() int64 {
	fileName := "./data/day07/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	combos := []PokerStruct{}
	for _, l := range lines {
		hand := strings.Split(l, " ")
		cards := strings.Split(hand[0], "")
		rank, _ := strconv.Atoi(hand[1])
		combo := PokerStruct{
			Combo: s.getCombo(cards),
			Cards: cards,
			Rank: rank,
		}
		combos = append(combos, combo)
	}

	sort.Slice(combos, func(i, j int) bool {
		if ComboWeight[combos[i].Combo] != ComboWeight[combos[j].Combo] {
			return ComboWeight[combos[i].Combo] < ComboWeight[combos[j].Combo]
		}
		for k := 0; k < 5; k++ {
			if CardWeight[combos[i].Cards[k]] != CardWeight[combos[j].Cards[k]] {
				return CardWeight[combos[i].Cards[k]] < CardWeight[combos[j].Cards[k]]
			}
		}
		return false
	})

	res := int64(0)
	for i, combo := range combos {
		res += int64((i+1) * combo.Rank)
	}

	return res
}

func (s *Service) getCombo(cards []string) string {
	cardCounts := map[string]int{}
	for _, card := range cards {
		if _, ok := cardCounts[card]; !ok {
			cardCounts[card] = 0
		}
		cardCounts[card]++
	}
	if len(cardCounts) == 5 {
		return ComboHigh
	}
	if len(cardCounts) == 1 {
		return ComboFive
	}
	if len(cardCounts) == 4 {
		return ComboOne
	}
	maxCount := 0
	for _, count := range cardCounts {
		if count > maxCount {
			maxCount = count
		}
	}
	if len(cardCounts) == 2 && maxCount == 4 {
		return ComboFour
	}
	if len(cardCounts) == 2{
		return ComboFull
	}
	if len(cardCounts) == 3 && maxCount == 3 {
		return ComboThree
	}
	if len(cardCounts) == 3 {
		return ComboTwo
	}
	return ""
}

func (s *Service) CalculateDay07SecondPart() int {
	fileName := "./data/day07/input1"
	file, _ := os.ReadFile(fileName)
	lines := strings.Split(string(file), "\n")

	combos := []PokerStruct{}
	for _, l := range lines {
		hand := strings.Split(l, " ")
		cards := strings.Split(hand[0], "")
		rank, _ := strconv.Atoi(hand[1])
		combo := PokerStruct{
			Combo: s.getComboWithJ(cards),
			Cards: cards,
			Rank: rank,
		}
		combos = append(combos, combo)
	}

	sort.Slice(combos, func(i, j int) bool {
		if ComboWeight[combos[i].Combo] != ComboWeight[combos[j].Combo] {
			return ComboWeight[combos[i].Combo] < ComboWeight[combos[j].Combo]
		}
		for k := 0; k < 5; k++ {
			iWeight := CardWeight[combos[i].Cards[k]]
			jWeight := CardWeight[combos[j].Cards[k]]
			if combos[j].Cards[k] == "J" {
				jWeight = 0
			}
			if combos[i].Cards[k] == "J" {
				iWeight = 0
			}
			if iWeight != jWeight {
				return iWeight < jWeight
			}
		}
		return false
	})

	res := 0
	for i, combo := range combos {
		res += (i+1) * combo.Rank
	}

	return res
}

func (s *Service) getComboWithJ(cards []string) string {
	cardCounts := map[string]int{}
	for _, card := range cards {
		if card == "J" {
			continue
		}
		if _, ok := cardCounts[card]; !ok {
			cardCounts[card] = 0
		}
		cardCounts[card]++
	}
	if len(cardCounts) == 5 {
		return ComboHigh
	}
	if len(cardCounts) == 0 {
		return ComboFive
	}
	if len(cardCounts) == 1 {
		return ComboFive
	}
	if len(cardCounts) == 4 {
		return ComboOne
	}
	minCount := 10
	maxCount := 0
	summ := 0
	for _, count := range cardCounts {
		if count < minCount {
			minCount = count
		}
		if count > maxCount {
			maxCount = count
		}
		summ+=count
	}
	if len(cardCounts) == 2 && minCount == 1 {
		return ComboFour
	}
	if len(cardCounts) == 2 {
		return ComboFull
	}
	if len(cardCounts) == 3 && summ == 5 && maxCount == 2 {
		return ComboTwo
	}
	if len(cardCounts) == 3 {
		return ComboThree
	}
	return ""
}
package main

import (
	"fmt"

	"aoc2023/src"
)

func main() {
	sAOC := src.NewService()
	res := sAOC.CalculateDay01SecondPart()
	fmt.Println(res)
}

package main

import (
	"fmt"

	"aoc2023/src"
)

func main() {
	sAOC := src.NewService()
	//res := sAOC.CalculateDay01FirstPart()
	//res := sAOC.CalculateDay01SecondPart()
	//res := sAOC.CalculateDay02FirstPart()
	//res := sAOC.CalculateDay02SecondPart()
	//res := sAOC.CalculateDay03FirstPart()
	res := sAOC.CalculateDay03SecondPart()
	fmt.Println(res)
}

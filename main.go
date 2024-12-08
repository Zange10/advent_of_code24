package main

import (
	"aoc/sol_01"
	"aoc/sol_02"
	"aoc/sol_03"
	"aoc/sol_04"
	"aoc/sol_05"
	"aoc/sol_06"
)

func main() {
	day := 6

	switch day {
	case 1:
		sol_01.RunSol01()
		break
	case 2:
		sol_02.RunSol02()
		break
	case 3:
		sol_03.RunSol03()
		break
	case 4:
		sol_04.RunSol04()
		break
	case 5:
		sol_05.RunSol05()
		break
	case 6:
		sol_06.RunSol06()
		break
	}
}

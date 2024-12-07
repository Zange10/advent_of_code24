package main

import (
	"aoc/sol_01"
	"aoc/sol_02"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	day := 2

	switch day {
	case 1:
		sol_01.RunSol01()
	case 2:
		sol_02.RunSol02()
	}
}

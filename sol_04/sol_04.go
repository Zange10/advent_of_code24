package sol_04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
)

func RunSol04() {

	var data [][]int

	// Parse input file --------------

	// open file
	f, err := os.Open("./sol_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var line_i []int
		for _,val := range line {
			switch val {
			case 'X':
				line_i = append(line_i, 1)
				break
			case 'M':
				line_i = append(line_i, 2)
				break
			case 'A':
				line_i = append(line_i, 3)
				break
			case 'S':
				line_i = append(line_i, 4)
				break
			}
		}
		data = append(data, line_i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Solution --------------
	count := 0

	for row,row_vals := range data {
		for col,val := range row_vals {
			if val == 1 {
				count += count_xmas_from_starting_cell(data, row, col)
			}
		}
	}

	fmt.Println(count)



	count = 0

	for row,row_vals := range data {
		for col,val := range row_vals {
			if val == 3 && row >= 1 && col >= 1 && row+1 < len(data) && col+1 < len(data[row]) {
				if has_mas_cross_from_center_cell(data, row, col) {count++}
				if has_mas_cross_from_center_cell(data, row, col) {
					fmt.Println("")
					fmt.Println(data[row-1][col-1:col+2])
					fmt.Println(data[row  ][col-1:col+2])
					fmt.Println(data[row+1][col-1:col+2])
				}
			}
		}
	}

	fmt.Println(count)
}


func count_xmas_from_starting_cell(data [][]int, row int, col int) int {
	count := 0

	// left-right row
	if col+3 < len(data[row]) {
		if data[row][col+1] == 2 && data[row][col+2] == 3 && data[row][col+3] == 4 {count++}
	}

	// right-left row
	if col >= 3 {
		if data[row][col-1] == 2 && data[row][col-2] == 3 && data[row][col-3] == 4 {count++}
	}

	// top-down column
	if row+3 < len(data) {
		if data[row+1][col] == 2 && data[row+2][col] == 3 && data[row+3][col] == 4 {count++}
	}

	// bottom-up column
	if row >= 3 {
		if data[row-1][col] == 2 && data[row-2][col] == 3 && data[row-3][col] == 4 {count++}
	}

	// topleft-bottomright diagonal
	if col+3 < len(data[row]) && row+3 < len(data) {
		if data[row+1][col+1] == 2 && data[row+2][col+2] == 3 && data[row+3][col+3] == 4 {count++}
	}

	// topright-bottomleft diagonal
	if col >= 3 && row+3 < len(data) {
		if data[row+1][col-1] == 2 && data[row+2][col-2] == 3 && data[row+3][col-3] == 4 {count++}
	}

	// bottomleft-topright diagonal
	if col+3 < len(data[row]) && row >= 3 {
		if data[row-1][col+1] == 2 && data[row-2][col+2] == 3 && data[row-3][col+3] == 4 {count++}
	}

	// bottomright-topleft diagonal
	if col >= 3 && row >= 3 {
		if data[row-1][col-1] == 2 && data[row-2][col-2] == 3 && data[row-3][col-3] == 4 {count++}
	}

	return count
}


func has_mas_cross_from_center_cell(data [][]int, row int, col int) bool {
	tl := data[row-1][col-1]
	tr := data[row-1][col+1]
	bl := data[row+1][col-1]
	br := data[row+1][col+1]

	if !(tl == 2 || tl == 4) {return false}
	if !(tr == 2 || tr == 4) {return false}
	if !(bl == 2 || bl == 4) {return false}
	if !(br == 2 || br == 4) {return false}

	if tl == br {return false}
	if tr == bl {return false}

	return true
}

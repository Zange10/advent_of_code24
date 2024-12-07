package sol_02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunSol02() {

	var data [][]int

	// open file
	f, err := os.Open("./sol_02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)



	for scanner.Scan() {
		var report []int

		// do something with a line
		temp := strings.Fields(scanner.Text())

		for _, val_s := range temp {
			val_i, _ := strconv.Atoi(val_s)
			report = append(report, val_i)
		}
		
		data = append(data, report)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	amt := 0
	amt_damp := 0
	fail := 0

	for _, report := range data {
		if is_increasing_correctly(report) || is_decreasing_correctly(report) {
			amt += 1
		}
		if is_increasing_correctly_damp(report) || is_decreasing_correctly_damp(report) {
			amt_damp += 1
		} else {
			fail += 1
		}
	}

	fmt.Printf("%d  %d  %d\n", amt, amt_damp, fail)
}


func is_increasing_correctly(report []int) bool {
	last := report[0]
	for _, val := range report[1:] {
		if val <= last || val > last+3 {
			return false
		}
		last = val
	}

	return true
}

func is_decreasing_correctly(report []int) bool {
	last := report[0]
	for _, val := range report[1:] {
		if val >= last || val < last-3 {
			return false
		}
		last = val
	}

	return true
}


func is_increasing_correctly_damp(report []int) bool {
	damp := false
	last := report[0]
	for _, val := range report[1:] {
		if val <= last || val > last+3 {
			if !damp {
				damp = true
				continue
			}
			return is_increasing_correctly(report[1:])
		}
		last = val
	}

	return true
}

func is_decreasing_correctly_damp(report []int) bool {
	damp := false
	last := report[0]
	for _, val := range report[1:] {
		if val >= last || val < last-3 {
			if !damp {
				damp = true
				continue
			}
			return is_decreasing_correctly(report[1:])
		}
		last = val
	}

	return true
}

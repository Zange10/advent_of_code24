package sol_03

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "strings"
)

func RunSol03() {

	sum1 := 0
	sum2 := 0


	// Parse input file --------------

	// open file
	f, err := os.Open("./sol_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	count1 := 0
	count := 0

	do := true

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line)-6; i++ {
			s := line[i:]

			if s[:4] == "do()" {do = true}
			if s[:7] == "don't()" {do = false}

			if s[:4] != "mul(" {continue}

			comma_ind := 5
			for s[comma_ind] >= '0' && s[comma_ind] <= '9' {comma_ind++}
			if s[comma_ind] != ',' {continue}

			bracket_ind := comma_ind+1
			for s[bracket_ind] >= '0' && s[bracket_ind] <= '9' {bracket_ind++}
			if s[bracket_ind] != ')' {continue}

			x,_ := strconv.Atoi(s[4:comma_ind])
			y,_ := strconv.Atoi(s[comma_ind+1:bracket_ind])

			sum1 += x*y
			if do {sum2 += x*y; count1++}
			if !do {count++}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	

	// Solution --------------
	fmt.Printf("%d  %d %d %d\n", sum1, sum2, count, count1)
}

package sol_01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func RunSol01() {

	var left []int
	var right []int

	// open file
	f, err := os.Open("./sol_01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		temp := strings.Fields(scanner.Text())

		templeft, _ := strconv.Atoi(temp[0])
		tempright, _ := strconv.Atoi(temp[1])

		left = append(left, templeft)
		right = append(right, tempright)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)

	sort.Ints(right)

	sum := 0

	for i, _ := range left {
		tempdiff := left[i] - right[i]
		if tempdiff < 0 {
			tempdiff = -tempdiff
		}
		sum += tempdiff
	}
	fmt.Printf("%d\n", sum)

	sum = 0

	for _, val_left := range left {
		for _, val_right := range right {
			if val_left == val_right {
				sum += val_left
			}
		}
	}

	fmt.Printf("%d\n", sum)

}

package sol_05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func RunSol05() {

	var rules [][]int
	var pages [][]int

	// Parse input file --------------

	// open file
	f, err := os.Open("./sol_05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	is_rule := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {is_rule = false; continue}
		
		if is_rule {
			var newrule []int

			x,_ := strconv.Atoi(line[0:2])
			y,_ := strconv.Atoi(line[3:5])

			newrule = append(newrule, x)
			newrule = append(newrule, y)

			rules = append(rules, newrule)
		} else {
			new_pages_s := strings.Split(line,",")
			var new_pages []int
			for _,val := range new_pages_s {
				x,_ := strconv.Atoi(val)
				new_pages = append(new_pages, x)
			}
			pages = append(pages, new_pages)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Solution --------------
	sum := 0
	for _,p := range pages {
		sum += find_valid_middle(rules, p)
	}
	fmt.Println(sum)

	
	sum = 0
	for _,p := range pages {
		was_ordered := false
		p,was_ordered = order_pages(rules, p)
		if was_ordered {sum += find_valid_middle(rules, p)}
	}
	fmt.Println(sum)
}


func find_valid_middle(rules [][]int, pages []int) int {
	for i,page1 := range pages {
		for _,page2 := range pages[:i] {
			if !rule_is_valid(rules, page1, page2) {return 0}
		}
	}

	middle := (len(pages)-1)/2
	return pages[middle]
}

func rule_is_valid(rules [][]int, page1 int, page2 int) bool {
	for _,rule := range rules {
		if rule[0] == page1 && rule[1] == page2 {return false}
	}
	return true
}

func order_pages(rules [][]int, pages []int) ([]int, bool) {
	was_ordered := false
	for i:=0; i<len(pages);i++ {
		for j:=0; j<i; j++ {
			if !rule_is_valid(rules, pages[i], pages[j]) {
				temp := pages[i]
				pages[i] = pages[j]
				pages[j] = temp
				i = 0
				was_ordered = true;
				break;
			}
		}
	}

	return pages, was_ordered
}
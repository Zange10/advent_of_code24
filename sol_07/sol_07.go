package sol_07

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunSol07() {
	eq_comps := parse_input_file()
	sum1 := 0
	sum2 := 0
	for _, eq := range eq_comps {
		new_sum1 := get_eq_result(eq, false)
		sum1 += new_sum1
		if new_sum1 == 0 {
			sum2 += get_eq_result(eq, true)
		}
		// fmt.Println(get_eq_result(eq, true))
	}

	fmt.Println(sum1)
	fmt.Println(sum1 + sum2)
}

func parse_input_file() [][]int {
	var input_array [][]int

	f, err := os.Open("./sol_07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var line_array []int
		line := scanner.Text()
		colon_idx := 0
		for i, c := range line {
			if c == ':' {
				colon_idx = i
				break
			}
		}

		res_val, _ := strconv.Atoi(line[:colon_idx])
		line_array = append(line_array, res_val)

		temp := strings.Fields(line[colon_idx+2:])

		for _, val_s := range temp {
			val_i, _ := strconv.Atoi(val_s)
			line_array = append(line_array, val_i)
		}
		input_array = append(input_array, line_array)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input_array
}

func get_eq_result(eq_vals []int, do_concat bool) int {
	res := eq_vals[0]
	eq_vals = eq_vals[1:]
	num_ops := len(eq_vals) - 1
	operators := 2
	if do_concat {
		operators++
	}

	for ops_i := 0; ops_i < int(math.Pow(float64(operators), float64(num_ops))); ops_i++ {
		ops := int2ops(ops_i, num_ops, operators)
		if res == calc_equation(eq_vals, ops) {
			return res
		}
	}

	return 0
}

func int2ops(num_i int, size int, operators int) []int {
	num_b := make([]int, size)
	for i := size - 1; i >= 0; i-- {
		num_b[i] = num_i % operators
		num_i /= operators
	}
	return num_b
}

func calc_equation(vals []int, ops []int) int {
	res := vals[0]
	for i, op := range ops {
		if op == 0 {
			res *= vals[i+1]
		} else if op == 1 {
			res += vals[i+1]
		} else {
			res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, vals[i+1]))
		}
	}
	return res
}

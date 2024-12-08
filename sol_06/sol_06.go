package sol_06

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
)

const (
	UP int = 0
	DOWN   = 1
	RIGHT  = 2
	LEFT   = 3
)

func RunSol06() {
	var cells [][]int
	var dir int

	// Parse input file --------------

	// open file
	f, err := os.Open("./sol_06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var new_row []int
		for _,c := range line {
			switch c {
			case '.':
				new_row = append(new_row, 0)
				break
			case '#':
				new_row = append(new_row, 1)
				break
			case '^':
				new_row = append(new_row, 2)
				dir = UP
				break
			}
		}
		cells = append(cells, new_row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Solution --------------
	
	no1 := false
	no2 := true

	if no1 {fmt.Println(count_steps(cells, dir))}
	if no2 {fmt.Println(count_loops(cells, dir))}
}

func get_curr_position(cells [][]int) (int,int) {
	for row_id, row := range cells {
		for col_id, val := range row {
			if val == 2 {
				return row_id, col_id
			}
		}
	}
	return -1, -1
}

func count_steps(cells [][]int, dir int) int {
	curr_row, curr_col := get_curr_position(cells)

	has_not_exited := true
	for has_not_exited {

		if dir == UP {
			if curr_row == 0 {
				has_not_exited = false
				break
			}
			if cells[curr_row-1][curr_col] == 1 {
				dir = RIGHT
				continue
			}
			curr_row--
		}

		if dir == LEFT {
			if curr_col == 0 {
				has_not_exited = false
				break
			}
			if cells[curr_row][curr_col-1] == 1 {
				dir = UP
				continue
			}
			curr_col--
		}

		if dir == DOWN {
			if curr_row == len(cells)-1 {
				has_not_exited = false
				break
			}
			if cells[curr_row+1][curr_col] == 1 {
				dir = LEFT
				continue
			}
			curr_row++
		}

		if dir == RIGHT {
			if curr_col == len(cells[0])-1 {
				has_not_exited = false
				break
			}
			if cells[curr_row][curr_col+1] == 1 {
				dir = DOWN
				continue
			}
			curr_col++
		}

		cells[curr_row][curr_col] = 2
	}

	steps := 0
	for _, row := range cells {
		for _, val := range row {
			if val == 2 {steps++}
		}
	}
	return steps
}

func count_loops(cells [][]int, dir int) int {
	curr_row, curr_col := get_curr_position(cells)

	var new_obstacles [][]int

	has_not_exited := true
	for has_not_exited {

		if dir == UP {
			if curr_row == 0 {
				has_not_exited = false
				break
			}
			if cells[curr_row-1][curr_col] == 1 {
				dir = RIGHT
				continue
			}
			new_obstacle := []int{curr_row-1, curr_col}
			if cells[new_obstacle[0]][new_obstacle[1]] != 2 && is_new_obstacle(new_obstacle, new_obstacles) {
				var new_cells [][]int
				for i := range cells {
					var new_row []int
					for j := range cells[i] {
						new_row = append(new_row, cells[i][j])
					}
					new_cells = append(new_cells, new_row)
				}
				new_cells[new_obstacle[0]][new_obstacle[1]] = 1
				curr_pos := []int{curr_row, curr_col}
				if is_loop(new_cells, curr_pos, dir) {
					new_obstacles = append(new_obstacles, new_obstacle)
				}
			}

			curr_row--
		}

		if dir == LEFT {
			if curr_col == 0 {
				has_not_exited = false
				break
			}
			if cells[curr_row][curr_col-1] == 1 {
				dir = UP
				continue
			}
			new_obstacle := []int{curr_row, curr_col-1}
			if cells[new_obstacle[0]][new_obstacle[1]] != 2 && is_new_obstacle(new_obstacle, new_obstacles) {
				var new_cells [][]int
				for i := range cells {
					var new_row []int
					for j := range cells[i] {
						new_row = append(new_row, cells[i][j])
					}
					new_cells = append(new_cells, new_row)
				}
				new_cells[new_obstacle[0]][new_obstacle[1]] = 1
				curr_pos := []int{curr_row, curr_col}
				if is_loop(new_cells, curr_pos, dir) {
					new_obstacles = append(new_obstacles, new_obstacle)
				}
			}

			curr_col--
		}

		if dir == DOWN {
			if curr_row == len(cells)-1 {
				has_not_exited = false
				break
			}
			if cells[curr_row+1][curr_col] == 1 {
				dir = LEFT
				continue
			}
			new_obstacle := []int{curr_row+1, curr_col}
			if cells[new_obstacle[0]][new_obstacle[1]] != 2 && is_new_obstacle(new_obstacle, new_obstacles) {
				var new_cells [][]int
				for i := range cells {
					var new_row []int
					for j := range cells[i] {
						new_row = append(new_row, cells[i][j])
					}
					new_cells = append(new_cells, new_row)
				}
				new_cells[new_obstacle[0]][new_obstacle[1]] = 1
				curr_pos := []int{curr_row, curr_col}
				if is_loop(new_cells, curr_pos, dir) {
					new_obstacles = append(new_obstacles, new_obstacle)
				}
			}

			curr_row++
		}

		if dir == RIGHT {
			if curr_col == len(cells[0])-1 {
				has_not_exited = false
				break
			}
			if cells[curr_row][curr_col+1] == 1 {
				dir = DOWN
				continue
			}
			new_obstacle := []int{curr_row, curr_col+1}
			if cells[new_obstacle[0]][new_obstacle[1]] != 2 && is_new_obstacle(new_obstacle, new_obstacles) {
				var new_cells [][]int
				for i := range cells {
					var new_row []int
					for j := range cells[i] {
						new_row = append(new_row, cells[i][j])
					}
					new_cells = append(new_cells, new_row)
				}
				new_cells[new_obstacle[0]][new_obstacle[1]] = 1
				curr_pos := []int{curr_row, curr_col}
				if is_loop(new_cells, curr_pos, dir) {
					new_obstacles = append(new_obstacles, new_obstacle)
				}
			}

			curr_col++
		}

		cells[curr_row][curr_col] = 2
	}


	return len(new_obstacles)
}

func is_new_obstacle(new_obstacle []int, obstacles [][]int) bool {
	for _,o := range obstacles {
		if o[0] == new_obstacle[0] && o[1] == new_obstacle[1] {return false}
	}
	return true
}


func is_loop(cells [][]int, initial_pos []int, dir int) bool {	
	curr_row := initial_pos[0]
	curr_col := initial_pos[1]
	initial_dir := dir
	has_not_exited := true

	counter := 0

	for has_not_exited {
		if dir == UP {
			if curr_row == 0 {
				return false
			}
			if cells[curr_row-1][curr_col] == 1 {
				dir = RIGHT
				continue
			}
			curr_row--
		}

		if dir == LEFT {
			if curr_col == 0 {
				return false
			}
			if cells[curr_row][curr_col-1] == 1 {
				dir = UP
				continue
			}
			curr_col--
		}

		if dir == DOWN {
			if curr_row == len(cells)-1 {
				return false
			}
			if cells[curr_row+1][curr_col] == 1 {
				dir = LEFT
				continue
			}
			curr_row++
		}

		if dir == RIGHT {
			if curr_col == len(cells[0])-1 {
				return false
			}
			if cells[curr_row][curr_col+1] == 1 {
				dir = DOWN
				continue
			}
			curr_col++
		}

		counter++

		if counter > 10000 {return true}

		if curr_row == initial_pos[0] && curr_col == initial_pos[1] && dir == initial_dir {return true}
	}

	return false
}

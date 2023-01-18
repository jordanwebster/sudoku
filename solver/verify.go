package main

import (
	"math"
)

func verify(solution [][]int, initial [][]int) bool {
	for j := 0; j < len(solution); j++ {
		seen := make(map[int]bool)
		for i := 0; i < len(solution); i++ {
			number := solution[j][i]
			_, was_seen := seen[number]
			if was_seen {
				return false
			}
			seen[number] = true
		}
	}

	for i := 0; i < len(solution); i++ {
		seen := make(map[int]bool)
		for j := 0; j < len(solution); j++ {
			number := solution[j][i]
			_, was_seen := seen[number]
			if was_seen {
				return false
			}
			seen[number] = true
		}
	}

	n := int(math.Sqrt(float64(len(solution))))
	for j := 0; j < len(solution); j += n {
		for i := 0; i < len(solution); i += n {
			seen := make(map[int]bool)
			for j_ := j; j_ < j+n; j_++ {
				for i_ := i; i_ < i+n; i_++ {
					number := solution[j_][i_]
					_, was_seen := seen[number]
					if was_seen {
						return false
					}
					seen[number] = true
				}
			}
		}
	}


    // Check that the proposed solution is actually a valid solution of the
    // given input sudoku.
    for j := 0; j < len(solution); j++ {
        for i := 0; i < len(solution); i++ {
            if initial[j][i] == UNSOLVED_CELL {
                continue
            }

            if initial[j][i] != solution[j][i] {
                return false
            }
        }
    }

	return true
}

func verify_partial(grid [][]int) bool {
	for j := 0; j < len(grid); j++ {
		seen := make(map[int]bool)
		for i := 0; i < len(grid); i++ {
			number := grid[j][i]
			if number == UNSOLVED_CELL {
				continue
			}

			_, was_seen := seen[number]
			if was_seen {
				return false
			}
			seen[number] = true
		}
	}

	for i := 0; i < len(grid); i++ {
		seen := make(map[int]bool)
		for j := 0; j < len(grid); j++ {
			number := grid[j][i]
			if number == UNSOLVED_CELL {
				continue
			}

			_, was_seen := seen[number]
			if was_seen {
				return false
			}
			seen[number] = true
		}
	}

	n := int(math.Sqrt(float64(len(grid))))
	for j := 0; j < len(grid); j += n {
		for i := 0; i < len(grid); i += n {
			seen := make(map[int]bool)
			for j_ := j; j_ < j+n; j_++ {
				for i_ := i; i_ < i+n; i_++ {
					number := grid[j_][i_]
					if number == UNSOLVED_CELL {
						continue
					}

					_, was_seen := seen[number]
					if was_seen {
						return false
					}
					seen[number] = true
				}
			}
		}
	}

	return true
}

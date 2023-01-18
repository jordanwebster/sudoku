package main

import (
	"errors"
)

const UNSOLVED_CELL = 0

func solve(grid [][]int) ([][]int, error) {
	result := make([][]int, len(grid))
	for j := 0; j < len(grid); j++ {
		result[j] = make([]int, len(grid[j]))
		copy(result[j], grid[j])
	}

	is_ok := solve_helper(result, 0, 0)
	if !is_ok {
		return nil, errors.New("Failed to find a valid solution")
	}

	return result, nil
}

func solve_helper(grid [][]int, i int, j int) bool {
	lower_bound := 1
	upper_bound := len(grid)
	is_clue := false

	if grid[j][i] != UNSOLVED_CELL {
		// This is a clue meaning we can jump to the next cell. We do this by
		// limiting the number of possible guesses to only the clue.
		is_clue = true
		lower_bound = grid[j][i]
		upper_bound = grid[j][i]
	}

	for candidate := lower_bound; candidate <= upper_bound; candidate++ {
		grid[j][i] = candidate
		if !verify_partial(grid) {
			continue
		}

		is_ok := false
		switch {
		case i < len(grid[j])-1:
			is_ok = solve_helper(grid, i+1, j)
		case j < len(grid)-1:
			is_ok = solve_helper(grid, 0, j+1)
		default:
			// Got to end of the grid and we have a verified solution
			is_ok = true
		}

		if is_ok {
			return true
		}
	}

	// Tried every possible candidate but did not find a solution so we need
	// to backtrack.
	if !is_clue {
		// Undo the write we made to this cell before we backtrack unless it's
		// a clue and can't be modified.
		grid[j][i] = UNSOLVED_CELL
	}
	return false
}

package main

import (
	"testing"
)

func Test2x2Grid(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 4},
		{0, 4, 0, 0},
		{2, 0, 4, 0},
		{0, 0, 0, 3},
	}

	result, _ := solve(grid)
	if !verify(result, grid) {
		t.Fatalf("%v is not a valid solution", result)
	}
}

func Test3x3Grid(t *testing.T) {
	grid := [][]int{
		{6, 0, 0, 0, 2, 0, 0, 0, 0},
		{4, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 3, 0, 6, 0, 7, 0, 0, 9},
		{0, 2, 0, 0, 9, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 0, 5},
		{3, 0, 0, 2, 0, 5, 0, 4, 0},
		{0, 6, 0, 5, 0, 3, 0, 0, 7},
		{0, 0, 3, 0, 0, 8, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 9, 0},
	}

	result, _ := solve(grid)
	if !verify(result, grid) {
		t.Fatalf("%v is not a valid solution", result)
	}
}

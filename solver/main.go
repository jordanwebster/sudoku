package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile(os.Args[1])
	var puzzle [][]int
    if err := json.Unmarshal(bytes, &puzzle); err != nil {
        panic(err)
    }

    solution, err := solve(puzzle)
    if err != nil {
        panic(err)
    }

    fmt.Printf("The solution is: %v\n", solution)
}

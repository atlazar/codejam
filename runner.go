package main

import (
	"fmt"
	"os"
)

type RunError struct {
	text string
}

func (this RunError) Error() string {
	return this.text
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return RunError{"Expected exactly one program argument (problem name)"}
	}

	solvers := map[string]func() (string, error){
		"problema": SolveA,
	}

	problem := os.Args[1]
	solver, exist := solvers[problem];
	if !exist {
		return RunError{fmt.Sprintf("Unknown problem \"%s\". Solver for this problem not found\n", problem)}
	}

	var cases int
	if _, err := fmt.Scan(&cases); err != nil {
		return RunError{fmt.Sprintf("Unable to read number of test cases. Error", err)}
	}

	for i := 0; i < cases; i++ {
		result, err := solver();
		if err != nil {
			return err
		}
		fmt.Printf("Case #%d: %s\n", i+1, result)
	}
	return nil
}

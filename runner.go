package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return errors.New("expected exactly one program argument (problem name)")
	}

	solvers := map[string]func() (string, error){
		"problema": SolveA,
		"problemb": SolveB,
	}

	problem := os.Args[1]
	solver, exist := solvers[problem];
	if !exist {
		return fmt.Errorf("unknown problem '%v'. Solver for this problem not found", problem)
	}

	var cases int
	if _, err := fmt.Scan(&cases); err != nil {
		return fmt.Errorf("unable to read number of test cases. Error: '%v'", err)
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

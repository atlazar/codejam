package main

import (
	"fmt"
	"math"
	"sort"
)

type Pancake struct {
	R, H        int
	Stop, Sside float64
}

func SolveA() (string, error) {
	stackSize, cakes, err := readCase()
	if err != nil {
		return "", err
	}
	sort.Slice(cakes, func(i, j int) bool {

		if cakes[i].R == cakes[j].R {
			return cakes[i].H < cakes[j].H
		}

		return cakes[i].R > cakes[j].R
	})
	return fmt.Sprintf("%f", solve(cakes, stackSize)), nil
}

func readCase() (stackSize int, cakes []Pancake, err error) {
	var totalSize int
	if _, e := fmt.Scan(&totalSize, &stackSize); e != nil {
		err = RunError{fmt.Sprint("Unable to read total/stack sizes. Error:", e)}
		return
	}

	cakes = make([]Pancake, totalSize)
	for i := 0; i < totalSize; i++ {
		var r, h int
		if _, e := fmt.Scan(&r, &h); e != nil {
			err = RunError{fmt.Sprint("Unable to read pancake. Error:", e)}
			return
		}
		cakes[i] = newPancake(r, h)
	}
	return
}

func newPancake(r int, h int) Pancake {
	return Pancake{
		R:     r,
		H:     h,
		Stop:  math.Pi * float64(r*r),
		Sside: math.Pi * float64(2*r*h)}
}

func sideSquare(cakes []Pancake, stackSize int) float64 {
	if len(cakes) < stackSize {
		panic("Number of cakes is less then expected stack size")
	}
	sortedBySside := make([]Pancake, len(cakes))
	copy(sortedBySside, cakes)
	sort.Slice(sortedBySside, func(i, j int) bool {
		return sortedBySside[i].Sside < sortedBySside[j].Sside
	})

	var square float64
	for i := len(sortedBySside) - stackSize; i < len(sortedBySside); i++ {
		square += sortedBySside[i].Sside
	}
	return square
}

func solve(sortedCakes []Pancake, stackSize int) float64 {
	if len(sortedCakes) < stackSize {
		panic("Number of cakes is less then expected stack size")
	}
	if stackSize <= 0 {
		panic("Non-positive stack size")
	}

	//Include bottom in stack
	square1 := sortedCakes[0].Stop + sortedCakes[0].Sside
	if stackSize > 1 {
		square1 += sideSquare(sortedCakes[1:], stackSize-1)
	}
	//fmt.Printf("Square1: %f Cakes: %+v\n", square1, sortedCakes)

	//Do not include bottom in stack if it possible
	if len(sortedCakes) > stackSize {
		square2 := solve(sortedCakes[1:], stackSize)
		return math.Max(square1, square2)
	}

	return square1
}

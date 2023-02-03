package main

import (
	"fmt"
	"os"
	"strings"

	"0036515065/optimization"
)

const (
	eDefault     = 1e-6
	alphaDefault = 1.3
)

func main() {
	zad1()
}

func zad1() {
	println("zad 1")
	var sb strings.Builder

	// banana
	firstFunc := optimization.FunctionOne()

	firstImplicit := optimization.OneOne
	secondImplicit := optimization.OneTwo
	implicits := []func(point *optimization.Point) float64{firstImplicit, secondImplicit}

	min := -100.0
	max := 100.0

	firstMin := optimization.Box(firstFunc.Begin, alphaDefault, eDefault, firstFunc, min, max, implicits, true)
	sb.WriteString(fmt.Sprintf("[banana] minimum: %s, evaluations: %d, value = %f\n", firstMin.ToString(), firstFunc.NumCalled, firstFunc.Call(firstMin)))

	secondFunc := optimization.FunctionTwo()
	secondMin := optimization.Box(secondFunc.Begin, alphaDefault, eDefault, secondFunc, min, max, implicits, true)
	sb.WriteString(fmt.Sprintf("[second] minimum: %s, evaluations: %d, value = %f\n", secondMin.ToString(), secondFunc.NumCalled, secondFunc.Call(secondMin)))

	print(sb.String())

	writeToFile(sb.String(), "solutions/zad1_solution.txt")
}

func writeToFile(str string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(str))
	if err != nil {
		panic(err)
	}
}

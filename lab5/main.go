package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"0036515065/internal"
	"0036515065/optimization"
)

func main() {
	zad1()
	zad2()
	zad3()
	zad4()
}

const printEvery = 100

func zad1() {
	println("zad 1")
	T := 0.01
	tMax := 10.0

	A := internal.NewMatrixFromFile("examples/zad1A.txt")
	begin := internal.NewMatrixFromFile("examples/zad1Begin.txt")

	correctXs := make([]float64, 0)
	correctYs := make([]float64, 0)
	for i := 0.0; i <= tMax; i += T {
		newX := 1*math.Cos(i) + 1*math.Sin(i)
		newY := 1*math.Cos(i) - 1*math.Sin(i)

		correctXs = append(correctXs, newX)
		correctYs = append(correctYs, newY)
	}

	sb := &strings.Builder{}
	sb.WriteString("Time\n")
	times := make([]float64, 0)
	for i := 0.0; i <= tMax; i += T {
		times = append(times, i)
	}
	writeListFloats(times, sb)
	sb.WriteString("\n")

	sb.WriteString("Correct\n")
	writeListFloats(correctXs, sb)
	writeListFloats(correctYs, sb)
	sb.WriteString("\n")

	rungeXs, rungeYs := optimization.Runge(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Runge\n")
	writeListFloats(rungeXs, sb)
	writeListFloats(rungeYs, sb)
	sb.WriteString("\n")

	trapezoidXs, trapezoidYs := optimization.Trapezoid(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Trapezoid\n")
	writeListFloats(trapezoidXs, sb)
	writeListFloats(trapezoidYs, sb)
	sb.WriteString("\n")

	eulerXs, eulerYs := optimization.Euler(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Euler\n")
	writeListFloats(eulerXs, sb)
	writeListFloats(eulerYs, sb)
	sb.WriteString("\n")

	reverseXs, reverseYs := optimization.Reverse(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Reverse Euler\n")
	writeListFloats(reverseXs, sb)
	writeListFloats(reverseYs, sb)
	sb.WriteString("\n")

	peceXs, peceYs := optimization.Pece(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Pece\n")
	writeListFloats(peceXs, sb)
	writeListFloats(peceYs, sb)
	sb.WriteString("\n")

	rungeSum := 0.0
	trapezoidSum := 0.0
	eulerSum := 0.0
	reverseSum := 0.0
	peceSum := 0.0
	for i := range correctXs {
		corrX := correctXs[i]
		corrY := correctYs[i]

		rungeSum += math.Sqrt(math.Pow(corrX-rungeXs[i], 2) + math.Pow(corrY-rungeYs[i], 2))
		trapezoidSum += math.Sqrt(math.Pow(corrX-trapezoidXs[i], 2) + math.Pow(corrY-trapezoidYs[i], 2))
		eulerSum += math.Sqrt(math.Pow(corrX-eulerXs[i], 2) + math.Pow(corrY-eulerYs[i], 2))
		reverseSum += math.Sqrt(math.Pow(corrX-reverseXs[i], 2) + math.Pow(corrY-reverseYs[i], 2))
		peceSum += math.Sqrt(math.Pow(corrX-peceXs[i], 2) + math.Pow(corrY-peceYs[i], 2))
	}

	fmt.Println("Runge Err: ", rungeSum)
	fmt.Println("Trapezoid Err: ", trapezoidSum)
	fmt.Println("Euler Err: ", eulerSum)
	fmt.Println("Reverse Err: ", reverseSum)
	fmt.Println("Pece Err: ", peceSum)

	writeToFile(sb.String(), "solutions/zad1_solution.txt")
}

func zad2() {
	println("zad 2")
	T := 0.01
	tMax := 10.0

	A := internal.NewMatrixFromFile("examples/zad2A.txt")
	begin := internal.NewMatrixFromFile("examples/zad2Begin.txt")

	sb := &strings.Builder{}
	sb.WriteString("Time\n")
	times := make([]float64, 0)
	for i := 0.0; i <= tMax; i += T {
		times = append(times, i)
	}
	writeListFloats(times, sb)
	sb.WriteString("\n")

	rungeXs, rungeYs := optimization.Runge(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Runge\n")
	writeListFloats(rungeXs, sb)
	writeListFloats(rungeYs, sb)
	sb.WriteString("\n")

	trapezoidXs, trapezoidYs := optimization.Trapezoid(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Trapezoid\n")
	writeListFloats(trapezoidXs, sb)
	writeListFloats(trapezoidYs, sb)
	sb.WriteString("\n")

	eulerXs, eulerYs := optimization.Euler(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Euler\n")
	writeListFloats(eulerXs, sb)
	writeListFloats(eulerYs, sb)
	sb.WriteString("\n")

	reverseXs, reverseYs := optimization.Reverse(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Reverse Euler\n")
	writeListFloats(reverseXs, sb)
	writeListFloats(reverseYs, sb)
	sb.WriteString("\n")

	peceXs, peceYs := optimization.Pece(A, nil, begin, nil, tMax, T, printEvery)
	sb.WriteString("Pece\n")
	writeListFloats(peceXs, sb)
	writeListFloats(peceYs, sb)
	sb.WriteString("\n")

	writeToFile(sb.String(), "solutions/zad2_solution.txt")
}

func zad3() {
	println("zad 3")
	T := 0.01
	tMax := 10.0

	A := internal.NewMatrixFromFile("examples/zad3A.txt")
	B := internal.NewMatrixFromFile("examples/zad3B.txt")
	begin := internal.NewMatrixFromFile("examples/zad3Begin.txt")

	f := func(t float64) *internal.Matrix {
		return internal.NewMatrixFromArray([][]float64{{1}, {1}})
	}

	sb := &strings.Builder{}
	sb.WriteString("Time\n")
	times := make([]float64, 0)
	for i := 0.0; i <= tMax; i += T {
		times = append(times, i)
	}
	writeListFloats(times, sb)
	sb.WriteString("\n")

	rungeXs, rungeYs := optimization.Runge(A, B, begin, f, tMax, T, printEvery)
	sb.WriteString("Runge\n")
	writeListFloats(rungeXs, sb)
	writeListFloats(rungeYs, sb)
	sb.WriteString("\n")

	trapezoidXs, trapezoidYs := optimization.Trapezoid(A, B, begin, f, tMax, T, printEvery)
	sb.WriteString("Trapezoid\n")
	writeListFloats(trapezoidXs, sb)
	writeListFloats(trapezoidYs, sb)
	sb.WriteString("\n")

	eulerXs, eulerYs := optimization.Euler(A, B, begin, f, tMax, T, printEvery)
	sb.WriteString("Euler\n")
	writeListFloats(eulerXs, sb)
	writeListFloats(eulerYs, sb)
	sb.WriteString("\n")

	reverseXs, reverseYs := optimization.Reverse(A, B, begin, f, tMax, T, printEvery)
	sb.WriteString("Reverse Euler\n")
	writeListFloats(reverseXs, sb)
	writeListFloats(reverseYs, sb)
	sb.WriteString("\n")

	peceXs, peceYs := optimization.Pece(A, B, begin, f, tMax, T, printEvery)
	sb.WriteString("Pece\n")
	writeListFloats(peceXs, sb)
	writeListFloats(peceYs, sb)
	sb.WriteString("\n")

	writeToFile(sb.String(), "solutions/zad3_solution.txt")
}

func zad4() {
	println("zad 4")
	T := 0.01
	tMax := 1.0

	A := internal.NewMatrixFromFile("examples/zad4A.txt")
	B := internal.NewMatrixFromFile("examples/zad4B.txt")
	begin := internal.NewMatrixFromFile("examples/zad4Begin.txt")

	f := func(t float64) *internal.Matrix {
		return internal.NewMatrixFromArray([][]float64{{t}, {t}})
	}

	sb := &strings.Builder{}
	sb.WriteString("Time\n")
	times := make([]float64, 0)
	for i := 0.0; i <= tMax; i += T {
		times = append(times, i)
	}
	writeListFloats(times, sb)
	sb.WriteString("\n")

	rungeXs, rungeYs := optimization.Runge(A, B, begin, f, tMax, T, printEvery/10)
	sb.WriteString("Runge\n")
	writeListFloats(rungeXs, sb)
	writeListFloats(rungeYs, sb)
	sb.WriteString("\n")

	trapezoidXs, trapezoidYs := optimization.Trapezoid(A, B, begin, f, tMax, T, printEvery/10)
	sb.WriteString("Trapezoid\n")
	writeListFloats(trapezoidXs, sb)
	writeListFloats(trapezoidYs, sb)
	sb.WriteString("\n")

	eulerXs, eulerYs := optimization.Euler(A, B, begin, f, tMax, T, printEvery/10)
	sb.WriteString("Euler\n")
	writeListFloats(eulerXs, sb)
	writeListFloats(eulerYs, sb)
	sb.WriteString("\n")

	reverseXs, reverseYs := optimization.Reverse(A, B, begin, f, tMax, T, printEvery/10)
	sb.WriteString("Reverse Euler\n")
	writeListFloats(reverseXs, sb)
	writeListFloats(reverseYs, sb)
	sb.WriteString("\n")

	peceXs, peceYs := optimization.Pece(A, B, begin, f, tMax, T, printEvery/10)
	sb.WriteString("Pece\n")
	writeListFloats(peceXs, sb)
	writeListFloats(peceYs, sb)
	sb.WriteString("\n")

	writeToFile(sb.String(), "solutions/zad4_solution.txt")
}

func writeListFloats(arr []float64, sb *strings.Builder) *strings.Builder {
	for _, num := range arr {
		n := strconv.FormatFloat(num, 'e', 10, 64)
		sb.WriteString(n + ", ")
	}
	sb.WriteString("\n")
	return sb
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

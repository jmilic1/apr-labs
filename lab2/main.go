package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"0036515065/optimization"
)

const (
	goldenConfig      = "config/golden.txt"
	coordinateConfig  = "config/coordinate.txt"
	simplexConfig     = "config/simplex.txt"
	hookeJeevesConfig = "config/hookeJeeves.txt"
	eDefault          = 1e-6
	stepDefault       = 1
	alphaDefault      = 1
	betaDefault       = 0.5
	gammaDefault      = 2
	sigmaDefault      = 0.5
	dxDefault         = 0.5
)

func main() {
	zad1()
	zad2()
	zad3()
	zad4()
	zad5()
}

func zad1() {
	println("zad 1")

	file, err := os.Open("examples/1zad.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var x float64
	for scanner.Scan() {
		row := scanner.Text()
		split := " = "
		if !strings.Contains(row, split) {
			continue
		}
		strArr := strings.Split(row, split)

		switch strArr[0] {
		case "x":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}
			x = value
		}
	}

	xP := optimization.CreatePoint([]float64{x})
	var sb strings.Builder

	functionGolden := optimization.FunctionFirstTask()
	println("golden ratio")
	goldenMinimum := doGolden(functionGolden, xP, true)
	sb.WriteString(fmt.Sprintf("[Golden]: minimum: %f, evaluations: %d, value = %f\n", goldenMinimum, functionGolden.NumCalled, functionGolden.Call(optimization.CreatePoint([]float64{goldenMinimum}))))

	xP = optimization.CreatePoint([]float64{x})
	functionCoordinate := optimization.FunctionFirstTask()
	println("coordinate")
	coordinateMinimum := doCoordinate(functionCoordinate, xP, true)
	sb.WriteString(fmt.Sprintf("[Coordinate]: minimum: %f, evaluations: %d, value = %f\n", coordinateMinimum.GetVal(0), functionCoordinate.NumCalled, functionCoordinate.Call(coordinateMinimum)))

	xP = optimization.CreatePoint([]float64{x})
	functionSimplex := optimization.FunctionFirstTask()
	println("simplex")
	simplexMinimum := doSimplex(functionSimplex, xP, true)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %f, evaluations: %d, value = %f\n", simplexMinimum.GetVal(0), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	xP = optimization.CreatePoint([]float64{x})
	functionHookeJeeeves := optimization.FunctionFirstTask()
	hookeJeevesMinimum := doHookeJeeves(functionHookeJeeeves, xP, true)
	sb.WriteString(fmt.Sprintf("[HookeJeeves]: minimum: %f, evaluations: %d, value = %f\n", hookeJeevesMinimum.GetVal(0), functionHookeJeeeves.NumCalled, functionHookeJeeeves.Call(hookeJeevesMinimum)))

	print(sb.String())

	writeToFile(sb.String(), "solutions/zad1_solution.txt")
}

func zad2() {
	println("zad 2")

	var sb strings.Builder

	x := getPoint("examples/2zad1.txt")

	println("first function\n")
	sb.WriteString("first function\n")
	secondTaskDoFirst(x, &sb)
	print(sb.String())

	sb.Reset()
	x = getPoint("examples/2zad2.txt")
	println("second function\n")
	sb.WriteString("second function\n")
	secondTaskDoSecond(x, &sb)
	print(sb.String())

	sb.Reset()
	x = getPoint("examples/2zad3.txt")
	println("third function\n")
	sb.WriteString("third function\n")
	secondTaskDoThird(x, &sb)
	print(sb.String())

	sb.Reset()
	x = getPoint("examples/2zad4.txt")
	println("fourth function\n")
	sb.WriteString("fourth function\n")
	secondTaskDoFourth(x, &sb)
	print(sb.String())

	print(sb.String())

	writeToFile(sb.String(), "solutions/zad2_solution.txt")
}

func zad3() {
	println("zad 3")

	var sb strings.Builder

	x := getPoint("examples/3zad.txt")

	functionSimplex := optimization.FunctionFour()
	println("simplex")
	simplexMinimum := doSimplex(functionSimplex, x.Copy(), true)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	functionHookeJeeeves := optimization.FunctionFour()
	println("hookeJeeves")
	hookeJeevesMinimum := doHookeJeeves(functionHookeJeeeves, x.Copy(), true)
	sb.WriteString(fmt.Sprintf("[HookeJeeves]: minimum: %s, evaluations: %d, value = %f\n", hookeJeevesMinimum.ToString(), functionHookeJeeeves.NumCalled, functionHookeJeeeves.Call(hookeJeevesMinimum)))

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad3_solution.txt")
}

func zad4() {
	println("zad 4")

	var sb strings.Builder

	x := getPoint("examples/4zad1.txt")

	functionSimplex := optimization.FunctionOne()
	println("simplex")
	simplexMinimum := doSimplex(functionSimplex, x.Copy(), true)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	x = getPoint("examples/4zad2.txt")

	functionSimplex = optimization.FunctionOne()
	println("simplex")
	simplexMinimum = doSimplex(functionSimplex, x.Copy(), true)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad4_solution.txt")
}

func zad5() {
	println("zad 5")

	var sb strings.Builder

	for i := 0; i < 5; i++ {
		min := -50
		max := 50
		x1 := float64(rand.Intn(max-min) + min)
		x2 := float64(rand.Intn(max-min) + min)
		vals := optimization.CreatePoint([]float64{x1, x2})

		functionSimplex := optimization.FunctionFive()

		println("simplex")
		simplexMinimum := doSimplex(functionSimplex, vals.Copy(), true)
		sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))
	}

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad4_solution.txt")
}

func getPoint(fileName string) *optimization.Point {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		split := " = "
		if !strings.Contains(row, split) {
			continue
		}
		strArr := strings.Split(row, split)

		switch strArr[0] {
		case "x":
			return optimization.CreatePoint(readMultipleValues(strArr[1], ", "))
		}
	}
	panic("x not found")
}

func doGolden(function *optimization.Function, x *optimization.Point, print bool) float64 {
	file, err := os.Open(goldenConfig)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var a, b, e float64
	eFlag := false
	for scanner.Scan() {
		row := scanner.Text()
		split := " = "
		if !strings.Contains(row, split) {
			continue
		}
		strArr := strings.Split(row, split)
		value, err := strconv.ParseFloat(strArr[1], 64)
		if err != nil {
			panic(err)
		}
		switch strArr[0] {
		case "a":
			a = value
		case "b":
			b = value
		case "e":
			e = value
			eFlag = true
		}
	}

	if !eFlag {
		e = eDefault
	}
	if x == nil {
		return optimization.Golden(optimization.CreatePoint([]float64{a}), optimization.CreatePoint([]float64{b}), function, e, print)
	} else {
		return optimization.GoldenDoUnimodal(1, x, function, e, print)
	}
}

func doCoordinate(function *optimization.Function, x *optimization.Point, print bool) *optimization.Point {
	file, err := os.Open(coordinateConfig)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var e float64
	eFlag := false
	for scanner.Scan() {
		row := scanner.Text()
		split := " = "
		if !strings.Contains(row, split) {
			continue
		}
		strArr := strings.Split(row, split)

		switch strArr[0] {
		case "x":
			x = optimization.CreatePoint(readMultipleValues(strArr[1], ", "))
		case "e":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			e = value
			eFlag = true
		}
	}

	if !eFlag {
		e = eDefault
	}
	return optimization.Coordinate(x, e, 1, function, print)
}

func doSimplex(function *optimization.Function, x *optimization.Point, print bool) *optimization.Point {
	file, err := os.Open(simplexConfig)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var e float64
	var step float64
	var alpha, beta, gamma, sigma float64

	var eFlag, stepFlag, alphaFlag, betaFlag, gammaFlag, sigmaFlag bool
	for scanner.Scan() {
		row := scanner.Text()
		split := " = "
		if !strings.Contains(row, split) {
			continue
		}
		strArr := strings.Split(row, split)

		switch strArr[0] {
		case "e":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			e = value
			eFlag = true
		case "step":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			step = value
			stepFlag = true
		case "alpha":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			alpha = value
			alphaFlag = true
		case "beta":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			beta = value
			betaFlag = true
		case "gamma":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			gamma = value
			gammaFlag = true
		case "sigma":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			sigma = value
			sigmaFlag = true
		}
	}

	if !eFlag {
		e = eDefault
	}
	if !stepFlag {
		step = stepDefault
	}
	if !alphaFlag {
		alpha = alphaDefault
	}
	if !betaFlag {
		beta = betaDefault
	}
	if !gammaFlag {
		gamma = gammaDefault
	}
	if !sigmaFlag {
		sigma = sigmaDefault
	}
	return optimization.Simplex(x, step, alpha, beta, gamma, sigma, e, function, print)
}

func doHookeJeeves(function *optimization.Function, x *optimization.Point, print bool) *optimization.Point {
	file, err := os.Open(hookeJeevesConfig)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var e float64
	var dx float64
	dxFlag := false
	eFlag := false
	for scanner.Scan() {
		row := scanner.Text()
		split := " = "
		if !strings.Contains(row, split) {
			continue
		}
		strArr := strings.Split(row, split)

		switch strArr[0] {
		case "dx":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			dx = value
			dxFlag = true
		case "e":
			value, err := strconv.ParseFloat(strArr[1], 64)
			if err != nil {
				panic(err)
			}

			e = value
			eFlag = true
		}
	}

	if !eFlag {
		e = eDefault
	}
	if !dxFlag {
		dx = dxDefault
	}
	return optimization.HookeJeeves(x, dx, function, e, print)
}

func readMultipleValues(str, delimiter string) []float64 {
	strArr := strings.Split(str, delimiter)
	values := make([]float64, len(strArr))

	for i, strElem := range strArr {
		value, err := strconv.ParseFloat(strElem, 64)
		if err != nil {
			panic(err)
		}
		values[i] = value
	}

	return values
}

func secondTaskDoFirst(x *optimization.Point, sb *strings.Builder) {
	functionCoordinate := optimization.FunctionOne()
	println("coordinate")
	coordinateMinimum := doCoordinate(functionCoordinate, x.Copy(), false)

	sb.WriteString(fmt.Sprintf("[Coordinate]: minimum: %s, evaluations: %d, value = %f\n", coordinateMinimum.ToString(), functionCoordinate.NumCalled, functionCoordinate.Call(coordinateMinimum)))

	functionSimplex := optimization.FunctionOne()
	println("simplex")
	simplexMinimum := doSimplex(functionSimplex, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	functionHookeJeeeves := optimization.FunctionOne()
	println("hookeJeeves")
	hookeJeevesMinimum := doHookeJeeves(functionHookeJeeeves, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[HookeJeeves]: minimum: %s, evaluations: %d, value = %f\n", hookeJeevesMinimum.ToString(), functionHookeJeeeves.NumCalled, functionHookeJeeeves.Call(hookeJeevesMinimum)))
}

func secondTaskDoSecond(x *optimization.Point, sb *strings.Builder) {
	functionCoordinate := optimization.FunctionTwo()
	println("coordinate")
	coordinateMinimum := doCoordinate(functionCoordinate, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[Coordinate]: minimum: %s, evaluations: %d, value = %f\n", coordinateMinimum.ToString(), functionCoordinate.NumCalled, functionCoordinate.Call(coordinateMinimum)))

	functionSimplex := optimization.FunctionTwo()
	println("simplex")
	simplexMinimum := doSimplex(functionSimplex, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	functionHookeJeeeves := optimization.FunctionTwo()
	println("hookeJeeves")
	hookeJeevesMinimum := doHookeJeeves(functionHookeJeeeves, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[HookeJeeves]: minimum: %s, evaluations: %d, value = %f\n", hookeJeevesMinimum.ToString(), functionHookeJeeeves.NumCalled, functionHookeJeeeves.Call(hookeJeevesMinimum)))
}

func secondTaskDoThird(x *optimization.Point, sb *strings.Builder) {
	functionCoordinate := optimization.FunctionThree()
	println("coordinate")
	coordinateMinimum := doCoordinate(functionCoordinate, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[Coordinate]: minimum: %s, evaluations: %d, value = %f\n", coordinateMinimum.ToString(), functionCoordinate.NumCalled, functionCoordinate.Call(coordinateMinimum)))

	functionSimplex := optimization.FunctionThree()
	println("simplex")
	simplexMinimum := doSimplex(functionSimplex, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	functionHookeJeeeves := optimization.FunctionThree()
	println("hookeJeeves")
	hookeJeevesMinimum := doHookeJeeves(functionHookeJeeeves, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[HookeJeeves]: minimum: %s, evaluations: %d, value = %f\n", hookeJeevesMinimum.ToString(), functionHookeJeeeves.NumCalled, functionHookeJeeeves.Call(hookeJeevesMinimum)))
}

func secondTaskDoFourth(x *optimization.Point, sb *strings.Builder) {
	functionCoordinate := optimization.FunctionFour()
	println("coordinate")
	coordinateMinimum := doCoordinate(functionCoordinate, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[Coordinate]: minimum: %s, evaluations: %d, value = %f\n", coordinateMinimum.ToString(), functionCoordinate.NumCalled, functionCoordinate.Call(coordinateMinimum)))

	functionSimplex := optimization.FunctionFour()
	println("simplex")
	simplexMinimum := doSimplex(functionSimplex, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[Simplex]: minimum: %s, evaluations: %d, value = %f\n", simplexMinimum.ToString(), functionSimplex.NumCalled, functionSimplex.Call(simplexMinimum)))

	functionHookeJeeeves := optimization.FunctionFour()
	println("hookeJeeves")
	hookeJeevesMinimum := doHookeJeeves(functionHookeJeeeves, x.Copy(), false)
	sb.WriteString(fmt.Sprintf("[HookeJeeves]: minimum: %s, evaluations: %d, value = %f\n", hookeJeevesMinimum.ToString(), functionHookeJeeeves.NumCalled, functionHookeJeeeves.Call(hookeJeevesMinimum)))
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

package main

import (
	"fmt"
	"os"
	"strings"

	"0036515065/optimization"
)

const eDefault = 1e-6

func main() {
	zad1()
	zad2()
	zad3()
	zad4()
	zad5()
	zad6()
}

func zad1() {
	println("zad 1")
	var sb strings.Builder

	firstFunc := optimization.FunctionThree()
	println("no golden ratio")
	firstMin := optimization.Gradient(firstFunc.Begin, eDefault, firstFunc, false, true)
	sb.WriteString(fmt.Sprintf("[No Golden]: minimum: %f, evaluations: %d, value = %f\n", firstMin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	secondFunc := optimization.FunctionThree()
	println("golden ratio")
	secondMin := optimization.Gradient(secondFunc.Begin, eDefault, secondFunc, true, true)
	sb.WriteString(fmt.Sprintf("[Golden]: minimum: %f, evaluations: %d, value = %f\n", secondMin, secondFunc.NumCalled, secondFunc.Call(secondMin)))

	print(sb.String())

	writeToFile(sb.String(), "solutions/zad1_solution.txt")
}

func zad2() {
	println("zad 2")

	var sb strings.Builder

	// firstFunc := FunctionOne()
	println("f1 Gradient")
	// firstMin := internal2.Gradient(firstFunc.Begin, eDefault, firstFunc, true, true)
	// sb.WriteString(fmt.Sprintf("[f1 Gradient]: minimum: %f, evaluations: %d, value = %f\n", firstMin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	secondFunc := optimization.FunctionOne()
	println("f1 Newton")
	secondMin := optimization.Newton(secondFunc.Begin, eDefault, secondFunc, true, true)
	sb.WriteString(fmt.Sprintf("[f1 Newton-Raphson]: minimum: %f, evaluations: %d, value = %f\n", secondMin, secondFunc.NumCalled, secondFunc.Call(secondMin)))

	// thirdFunc := FunctionTwo()
	println("f2 Gradient")
	// thirdMin := internal2.Gradient(thirdFunc.Begin, eDefault, thirdFunc, true, true)
	// sb.WriteString(fmt.Sprintf("[f2 Gradient]: minimum: %f, evaluations: %d, value = %f\n", thirdMin, thirdFunc.NumCalled, thirdFunc.Call(thirdMin)))

	fourthFunc := optimization.FunctionOne()
	println("f2 Newton")
	fourthMin := optimization.Newton(fourthFunc.Begin, eDefault, fourthFunc, true, true)
	sb.WriteString(fmt.Sprintf("[f2 Newton-Raphson]: minimum: %f, evaluations: %d, value = %f\n", fourthMin, fourthFunc.NumCalled, fourthFunc.Call(fourthMin)))

	print(sb.String())

	writeToFile(sb.String(), "solutions/zad2_solution.txt")
}

func zad3() {
	println("zad 3")

	firstPoint := optimization.CreatePoint([]float64{3.0, 3.0})
	secondPoint := optimization.CreatePoint([]float64{1.0, 2.0})

	var sb strings.Builder

	firstFunc := optimization.FunctionFour()
	println("(3, 3) no golden")
	firstMin := optimization.Newton(firstPoint, eDefault, firstFunc, false, true)
	sb.WriteString(fmt.Sprintf("[(3, 3) no golden]: minimum: %f, evaluations: %d, value = %f\n", firstMin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	secondFunc := optimization.FunctionFour()
	println("(3, 3) golden")
	secondMin := optimization.Newton(firstPoint, eDefault, secondFunc, true, true)
	sb.WriteString(fmt.Sprintf("[(3, 3) golden]: minimum: %f, evaluations: %d, value = %f\n", secondMin, secondFunc.NumCalled, secondFunc.Call(secondMin)))

	thirdFunc := optimization.FunctionFour()
	println("(1, 2) no golden")
	thirdMin := optimization.Newton(secondPoint, eDefault, thirdFunc, false, true)
	sb.WriteString(fmt.Sprintf("[(1, 2) no golden]: minimum: %f, evaluations: %d, value = %f\n", thirdMin, thirdFunc.NumCalled, thirdFunc.Call(thirdMin)))

	fourthFunc := optimization.FunctionFour()
	println("(1, 2) golden")
	fourthMin := optimization.Newton(secondPoint, eDefault, fourthFunc, true, true)
	sb.WriteString(fmt.Sprintf("[(1, 2) golden]: minimum: %f, evaluations: %d, value = %f\n", fourthMin, fourthFunc.NumCalled, fourthFunc.Call(fourthMin)))

	print(sb.String())

	writeToFile(sb.String(), "solutions/zad3_solution.txt")
}

func zad4() {
	println("zad 4")

	var sb strings.Builder

	firstFunc := optimization.FunctionOne()
	println("Gauss")
	firstMin := optimization.Gauss(firstFunc.Begin, eDefault, firstFunc, true)
	sb.WriteString(fmt.Sprintf("[Gauss]: minimum: %f, evaluations: %d, value = %f\n", firstMin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad4_solution.txt")
}

func zad5() {
	println("zad 5")

	firstPoint := optimization.CreatePoint([]float64{-2.0, 2.0})
	secondPoint := optimization.CreatePoint([]float64{2.0, 2.0})
	thirdPoint := optimization.CreatePoint([]float64{2.0, -2.0})

	var sb strings.Builder

	firstFunc := optimization.CustomFunc()
	println("Gauss (-2, 2)")
	firstMin := optimization.Gauss(firstPoint, eDefault, firstFunc, true)
	sb.WriteString(fmt.Sprintf("[Gauss (-2, 2)]: minimum: %f, evaluations: %d, value = %f\n", firstMin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	firstFunc = optimization.CustomFunc()
	println("Gauss (2, 2)")
	firstMin = optimization.Gauss(secondPoint, eDefault, firstFunc, true)
	sb.WriteString(fmt.Sprintf("[Gauss (2, 2)]: minimum: %f, evaluations: %d, value = %f\n", firstMin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	firstFunc = optimization.CustomFunc()
	println("Gauss (2, -2)")
	firstMin = optimization.Gauss(thirdPoint, eDefault, firstFunc, true)
	sb.WriteString(fmt.Sprintf("[Gauss (2, 2)]: minimum: %f, evaluations: %d, value = %f\n", firstMin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad5_solution.txt")
}

func zad6() {
	println("zad 6")

	begin := optimization.CreatePoint([]float64{1.0, 1.0, 1.0})

	var sb strings.Builder

	firstFunc := optimization.SecondCustomFunc()
	println("Gauss (1, 1, 1)")
	firstMin := optimization.Gauss(begin, eDefault, firstFunc, true)
	sb.WriteString(fmt.Sprintf("[Gauss (1, 1, 1)]: minimum: %f, evaluations: %d, value = %f\n", begin, firstFunc.NumCalled, firstFunc.Call(firstMin)))

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad5_solution.txt")
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

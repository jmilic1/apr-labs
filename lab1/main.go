package main

import (
	"os"
	"strconv"
	"strings"

	"0036515065/internal"
)

func main() {
	zad1()
	zad2()
	zad3()
	zad4()
	zad5()
	zad6()
	zad7()
	zad8()
	zad9()
	zad10()
}

func zad1() {
	println("zad 1")
	A := internal.NewMatrixFromFile("examples/zad1A.txt")
	ACopy := A.Copy()

	A.MulScalar(2.5)
	A.MulScalar(0.4)

	ACopy.MulScalar(2.5)
	ACopy.MulScalar(0.4)

	var sb strings.Builder
	sb.WriteString("A:\n" + A.String())
	sb.WriteString("New A:\n" + ACopy.String())

	var str string
	if A.Equal(ACopy) {
		str = "true"
	} else {
		str = "false"
	}
	sb.WriteString("Is A equal to ACopy? " + str + "\n")

	print(sb.String())

	writeToFile(sb.String(), "solutions/zad1_solution.txt")
}

func zad2() {
	println("zad 2")
	A := internal.NewMatrixFromFile("examples/zad2A.txt")
	B := internal.NewMatrixFromFile("examples/zad2B.txt")

	X := internal.SolveEquationLUP(A, B)

	var sb strings.Builder
	sb.WriteString("Solved with LUP:\nA\n" + A.String() + "b\n" + B.String() + "x\n" + X.String())

	A = internal.NewMatrixFromFile("examples/zad2A.txt")
	B = internal.NewMatrixFromFile("examples/zad2B.txt")

	X = internal.SolveEquationLU(A, B)

	sb.WriteString("Solved with LU:\nA\n" + A.String() + "b\n" + B.String() + "x\n" + X.String())
	print(sb.String())

	writeToFile(sb.String(), "solutions/zad2_solution.txt")
}

func zad3() {
	println("zad3")
	A := internal.NewMatrixFromFile("examples/zad3A.txt")

	LU := A.Copy()
	LU.LUDecomposition()

	LUP := A.Copy()
	LUP.LUPDecomposition()

	var sb strings.Builder
	sb.WriteString("LU:\n" + LU.String() + "LUP:\n" + LUP.String() + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad3_solution.txt")
}

func zad4() {
	println("zad4")

	A1 := internal.NewMatrixFromFile("examples/zad4A.txt")
	B1 := internal.NewMatrixFromFile("examples/zad4B.txt")

	A2 := A1.Copy()
	B2 := B1.Copy()

	X1 := internal.SolveEquationLU(A1, B1)
	X2 := internal.SolveEquationLUP(A2, B2)

	var sb strings.Builder
	sb.WriteString("LU:\n" + X1.String() + "LUP:\n" + X2.String() + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad4_solution.txt")
}

func zad5() {
	println("zad5")

	A := internal.NewMatrixFromFile("examples/zad5A.txt")
	B := internal.NewMatrixFromFile("examples/zad5B.txt")

	X := internal.SolveEquationLUP(A, B)
	var sb strings.Builder
	sb.WriteString("LUP:\n" + X.String() + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad5_solution.txt")
}

func zad6() {
	println("zad6")

	A := internal.NewMatrixFromFile("examples/zad6A.txt")
	B := internal.NewMatrixFromFile("examples/zad6B.txt")

	X := internal.SolveEquationLUP(A, B)
	var sb strings.Builder
	sb.WriteString("LUP:\n" + X.String() + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad6_solution.txt")
}

func zad7() {
	println("zad7")

	A := internal.NewMatrixFromFile("examples/zad7A.txt")
	det := A.Determinant()

	var sb strings.Builder
	sb.WriteString("A:\n" + A.String() + "\ndet(A): " + strconv.FormatFloat(det, 'e', 10, 64) + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad7_solution.txt")
}

func zad8() {
	println("zad8")

	A := internal.NewMatrixFromFile("examples/zad8A.txt")
	inverse := internal.Inverse(A)

	var sb strings.Builder
	sb.WriteString("A:\n" + A.String() + "\n")
	sb.WriteString("A inverse:\n" + inverse.String() + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad8_solution.txt")
}

func zad9() {
	println("zad9")

	A := internal.NewMatrixFromFile("examples/zad9A.txt")
	det := A.Determinant()

	var sb strings.Builder
	sb.WriteString("A determinant:\n" + strconv.FormatFloat(det, 'e', 10, 64) + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad9_solution.txt")
}

func zad10() {
	println("zad10")

	A := internal.NewMatrixFromFile("examples/zad9A.txt")
	det := A.Determinant()

	var sb strings.Builder
	sb.WriteString("A determinant:\n" + strconv.FormatFloat(det, 'e', 10, 64) + "\n")

	print(sb.String())
	writeToFile(sb.String(), "solutions/zad10_solution.txt")
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

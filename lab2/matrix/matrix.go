package matrix

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operation int64

const (
	plus Operation = iota
	minus
)

type Matrix struct {
	elements [][]float64
	rows     int
	columns  int

	// permutationNumber int
	// permutationMatrix [][]float64
	Perms []int
}

// region Constructors

func NewMatrix(rows, columns int) *Matrix {
	elements := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		elements[i] = make([]float64, columns)
	}
	return &Matrix{
		elements: elements,
		rows:     rows,
		columns:  columns,
	}
}

func NewMatrixFromArray(arr [][]float64) *Matrix {
	matrix := NewMatrix(len(arr), len(arr[0]))

	for i := 0; i < matrix.rows; i++ {
		for j := 0; j < matrix.columns; j++ {
			matrix.Set(i, j, arr[i][j])
		}
	}

	return matrix
}

func NewMatrixFromFile(fileName string) *Matrix {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rows := make([]string, 0)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	columnsDim := -1
	rowsDim := len(rows)
	elements := make([][]float64, rowsDim)
	for i, row := range rows {
		stringNums := strings.Split(row, " ")

		columns := len(stringNums)
		if columnsDim == -1 {
			columnsDim = columns
		}
		if columnsDim != columns {
			panic("invalid columns")
		}

		rowNums := make([]float64, columns)

		for j, numStr := range stringNums {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				panic(err)
			}

			rowNums[j] = num
		}
		elements[i] = rowNums
	}

	return &Matrix{
		elements: elements,
		rows:     rowsDim,
		columns:  columnsDim,
	}
}

// endregion Constructors

// region print

func (m *Matrix) String() string {
	var sb strings.Builder

	for _, row := range m.elements {
		for i, element := range row {
			sb.WriteString(fmt.Sprintf("%f", element))

			if i != len(row)-1 {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (m *Matrix) Print() {
	for _, row := range m.elements {
		for _, element := range row {
			print(element, " ")
		}
		println()
	}
}

func (m *Matrix) WriteToFile(fileName string) {
	str := m.String()

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(str))
	if err != nil {
		panic(err)
	}
}

// endregion print

// region aliases

func (m *Matrix) Get(row, column int) float64 {
	return m.elements[row][column]
}

func (m *Matrix) Set(row, column int, element float64) {
	m.elements[row][column] = element
}

func (m *Matrix) Plus(m2 *Matrix) *Matrix {
	return Plus(m, m2)
}

func (m *Matrix) Mul(m2 *Matrix) *Matrix {
	return Mul(m, m2)
}

func (m *Matrix) Transpose() *Matrix {
	return Transpose(m)
}

func Plus(m1 *Matrix, m2 *Matrix) *Matrix {
	return elementWise(plus, m1, m2)
}

func (m *Matrix) Equal(m2 *Matrix) bool {
	return Equal(m, m2)
}

func (m *Matrix) Forward(b *Matrix) *Matrix {
	return Forward(m, b)
}

func (m *Matrix) Backward(b *Matrix) *Matrix {
	return Backward(m, b)
}

func (m *Matrix) LUDecomposition() {
	LUDecomposition(m)
}

func (m *Matrix) LUPDecomposition() {
	LUPDecomposition(m)
}

// endregion aliases

func (m *Matrix) Row(row int) *Matrix {
	rowElements := m.elements[row]
	elements := make([][]float64, 1)
	elements[0] = rowElements
	return NewMatrixFromArray(elements)
}

func Mul(m1 *Matrix, m2 *Matrix) *Matrix {
	if m1.columns != m2.rows {
		panic("invalid rows")
	}

	rows := m1.rows
	columns := m2.columns
	common := m1.columns

	newMatrix := NewMatrix(rows, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			element := 0.0
			for k := 0; k < common; k++ {
				element += m1.Get(i, k) * m2.Get(k, j)
			}

			newMatrix.Set(i, j, element)
		}
	}

	return newMatrix
}

func Transpose(m *Matrix) *Matrix {
	rows := m.rows
	columns := m.columns

	newRows := columns
	newColumns := rows

	newMatrix := NewMatrix(newRows, newColumns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			element := m.Get(i, j)
			newMatrix.Set(j, i, element)
		}
	}

	return newMatrix
}

func (m *Matrix) PlusE(m2 *Matrix) {
	newMatrix := elementWise(plus, m, m2)
	*m = *newMatrix
}

func (m *Matrix) MinusE(m2 *Matrix) {
	newMatrix := elementWise(minus, m, m2)
	*m = *newMatrix
}

func Equal(m1 *Matrix, m2 *Matrix) bool {
	if m1.rows != m2.rows {
		return false
	}
	if m1.columns != m2.columns {
		return false
	}

	rows := m1.rows
	columns := m2.columns
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			firstElement := m1.Get(i, j)
			secondElement := m2.Get(i, j)

			if firstElement != secondElement {
				return false
			}
		}
	}

	return true
}

func Forward(m *Matrix, b *Matrix) *Matrix {
	rows := m.rows
	bElems := b.GetVector()

	for i := 0; i < rows-1; i++ {
		for j := i + 1; j < rows; j++ {
			bElems[j] -= m.Get(j, i) * bElems[i]
		}
	}

	return NewMatrixFromArray([][]float64{bElems})
}

func (m *Matrix) GetVector() []float64 {
	if m.rows != 1 && m.columns != 1 {
		panic("matrix is not a vector")
	}

	if m.rows == 1 {
		return m.elements[0]
	}

	return m.Transpose().elements[0]
}

func Backward(m *Matrix, b *Matrix) *Matrix {
	rows := m.rows
	bElems := b.GetVector()

	for i := rows - 1; i >= 0; i-- {
		bElems[i] /= m.Get(i, i)

		for j := 0; j <= i-1; j++ {
			bElems[j] -= m.Get(j, i) * bElems[i]
		}
	}

	return NewMatrixFromArray([][]float64{bElems})
}

func LUDecomposition(m *Matrix) {
	rows := m.rows
	for i := 0; i < rows-1; i++ {
		for j := i + 1; j < rows; j++ {
			old := m.Get(j, i)
			newElement := old / m.Get(i, i)
			m.Set(j, i, newElement)

			for k := i + 1; k < rows; k++ {
				old = m.Get(j, k)
				newElement = old - m.Get(j, i)*m.Get(i, k)
				m.Set(j, k, newElement)
			}
		}
	}
}

// LUPDecomposition returns permutation
func LUPDecomposition(m *Matrix) int {
	n := m.rows
	permutationNumber := 0
	P := make([]int, n)
	for i := range P {
		P[i] = i
	}

	for i := 0; i < n-1; i++ {
		pivot := i

		for j := i + 1; j < n; j++ {
			if math.Abs(m.Get(P[j], i)) > math.Abs(m.Get(P[pivot], i)) {
				pivot = j
			}
		}

		if m.Get(P[pivot], i) == 0 {
			return 0
		}

		// zamijeni(P[i], P[pivot]);
		temp := P[i]
		P[i] = P[pivot]
		P[pivot] = temp
		permutationNumber++

		for j := i + 1; j < n; j++ {
			old := m.Get(P[j], i)
			newElement := old / m.Get(P[i], i)
			m.Set(P[j], i, newElement)

			for k := i + 1; k < n; k++ {
				old = m.Get(P[j], k)
				newElement = old - m.Get(P[j], i)*m.Get(P[i], k)
				m.Set(P[j], k, newElement)
			}
		}
	}

	elements := make([][]float64, len(m.elements))
	for toIndex, fromIndex := range P {
		elements[toIndex] = m.elements[fromIndex]
	}
	m.elements = elements
	m.Perms = P

	return permutationNumber
}

func SolveEquationLUP(A *Matrix, b *Matrix) *Matrix {
	if b.rows != 1 && b.columns != 1 {
		panic("b is not a vector")
	}

	A.LUPDecomposition()
	L := A.ExtractL()
	U := A.ExtractU()

	newElements := make([][]float64, len(A.Perms))
	for targetIndex, index := range A.Perms {
		newElements[targetIndex] = b.elements[index]
	}
	B := NewMatrixFromArray(newElements)

	y := Forward(L, B)
	x := Backward(U, y)

	return x
}

func SolveEquationLU(A *Matrix, B *Matrix) *Matrix {
	A.LUDecomposition()

	L := A.ExtractL()
	U := A.ExtractU()

	y := L.Forward(B)
	x := U.Backward(y)

	return x
}

func IdentityMatrix(n int) *Matrix {
	values := make([][]float64, n)

	for i := 0; i < n; i++ {
		values[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				values[i][j] = 1
			}
		}
	}

	return NewMatrixFromArray(values)
}

func Inverse(m *Matrix) *Matrix {
	n := m.rows

	identity := IdentityMatrix(n)

	LU := m.Copy()
	LUPDecomposition(LU)

	L := LU.ExtractL()
	U := LU.ExtractU()

	inverseValues := make([][]float64, n)
	for idx := range inverseValues {
		inverseValues[idx] = make([]float64, n)
	}

	perms := LU.Perms

	for i := 0; i < n; i++ {
		Y := L.Forward(identity.Row(i))
		X := U.Backward(Y)
		XVec := X.GetVector()

		for j := 0; j < n; j++ {
			inverseValues[j][perms[i]] = XVec[j]
		}
	}

	return NewMatrixFromArray(inverseValues)
}

func (m *Matrix) ExtractL() *Matrix {
	rows := m.rows
	columns := m.columns

	values := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		values[i] = make([]float64, columns)
		for j := 0; j < rows; j++ {
			if j == i {
				values[i][j] = 1
			}
			if j > i {
				values[i][j] = 0
			}
			if j < i {
				values[i][j] = m.Get(i, j)
			}
		}
	}

	return NewMatrixFromArray(values)
}

func (m *Matrix) ExtractU() *Matrix {
	rows := m.rows
	columns := m.columns

	values := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		values[i] = make([]float64, columns)
		for j := 0; j < columns; j++ {
			if j == i {
				values[i][j] = m.Get(i, j)
			}
			if j < i {
				values[i][j] = 0
			}
			if j > i {
				values[i][j] = m.Get(i, j)
			}
		}
	}

	return NewMatrixFromArray(values)
}

func (m *Matrix) Determinant() float64 {
	a := m.Copy()
	p := LUPDecomposition(a)

	L := a.ExtractL()
	U := a.ExtractU()

	detL := L.DiagonalValue()
	detU := U.DiagonalValue()

	determinant := math.Pow(-1, float64(p)) * detL * detU
	if determinant < 1e-6 {
		return 0
	}

	return determinant
}

func (m *Matrix) DiagonalValue() float64 {
	value := 1.0

	for i := 0; i < m.rows; i++ {
		value *= m.Get(i, i)
	}

	return value
}

func (m *Matrix) Copy() *Matrix {
	rows := m.rows
	columns := m.columns

	values := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		values[i] = make([]float64, columns)
		for j := 0; j < columns; j++ {
			values[i][j] = m.Get(i, j)
		}
	}

	return NewMatrixFromArray(values)
}

func elementWise(operation Operation, m1 *Matrix, m2 *Matrix) *Matrix {
	if m1.rows != m2.rows {
		panic("invalid rows")
	}
	if m1.columns != m2.columns {
		panic("invalid columns")
	}
	rows := m1.rows
	columns := m1.columns

	newMatrix := NewMatrix(rows, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			firstElement := m1.Get(i, j)
			secondElement := m2.Get(i, j)

			var newElement float64
			switch operation {
			case plus:
				newElement = firstElement + secondElement
			case minus:
				newElement = firstElement - secondElement
			}

			newMatrix.Set(i, j, newElement)
		}
	}

	return newMatrix
}

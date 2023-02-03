package optimization

import (
	"math"

	"0036515065/matrix"
)

type Function struct {
	NumCalled int
	SetDim    int
	Begin     *Point
	Min       *Point
	MinVal    float64
	function  func(point *Point) float64
	gradient  func(point *Point) *Point
	Hess      func(point *Point) *matrix.Matrix
	Jacob     func(point *Point) *matrix.Matrix
	NonLinear func(point *Point) *matrix.Matrix
}

func NewFunction(function func(*Point) float64, begin, min *Point, gradient func(point *Point) *Point) *Function {
	var setDim int
	var minVal float64
	if min != nil {
		setDim = min.Dims()
		minVal = function(min)
	}
	return &Function{
		NumCalled: 0,
		SetDim:    setDim,
		Begin:     begin,
		Min:       min,
		MinVal:    minVal,
		function:  function,
		gradient:  gradient,
	}
}

func (f *Function) Call(point *Point) float64 {
	f.NumCalled++
	return f.function(point)
}

func (f *Function) Gradient(point *Point) *Point {
	return f.gradient(point)
}

// FunctionOne is banana
func FunctionOne() *Function {
	f := func(point *Point) float64 {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		first := math.Pow(x2-math.Pow(x1, 2), 2)
		return 100*first + math.Pow(1-x1, 2)
	}

	vals := []float64{-1.9, 2}
	begin := CreatePoint(vals)

	vals = []float64{1, 1}
	min := CreatePoint(vals)

	gradient := func(point *Point) *Point {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		firstDer := 400*math.Pow(x1, 3) + (2-400*x2)*x1 - 2
		secondDer := 200 * (x2 - math.Pow(x1, 2))
		return CreatePoint([]float64{firstDer, secondDer})
	}

	fu := NewFunction(f, begin, min, gradient)

	fu.Hess = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		oneOne := 1200*math.Pow(x1, 2) - 400*x2 + 2
		oneTwo := -400 * x1
		twoTwo := 200.0

		mat := [][]float64{{oneOne, oneTwo}, {oneTwo, twoTwo}}
		return matrix.NewMatrixFromArray(mat)
	}

	return fu
}

func FunctionTwo() *Function {
	f := func(point *Point) float64 {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		return math.Pow(x1-4, 2) + 4*math.Pow(x2-2, 2)
	}

	vals := []float64{0.1, 0.3}
	begin := CreatePoint(vals)

	vals = []float64{4, 2}
	min := CreatePoint(vals)

	gradient := func(point *Point) *Point {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		firstDer := 2 * (x1 - 4)
		secondDer := 8 * (x2 - 2)
		return CreatePoint([]float64{firstDer, secondDer})
	}

	fu := NewFunction(f, begin, min, gradient)

	fu.Hess = func(point *Point) *matrix.Matrix {
		oneOne := 2.0
		oneTwo := 0.0
		twoTwo := 8.0

		mat := [][]float64{{oneOne, oneTwo}, {oneTwo, twoTwo}}
		return matrix.NewMatrixFromArray(mat)
	}

	return fu
}

func FunctionThree() *Function {
	f := func(point *Point) float64 {
		sum := 0.0
		for i := 0; i < point.Dims(); i++ {
			sum += math.Pow(point.GetVal(i)-float64(i), 2)
		}

		return sum
	}

	vals := []float64{0, 0, 0, 0, 0}
	begin := CreatePoint(vals)

	vals = []float64{1, 2, 3, 4, 5}
	min := CreatePoint(vals)

	return NewFunction(f, begin, min, nil)
}

func FunctionFour() *Function {
	f := func(point *Point) float64 {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		return math.Abs((x1-x2)*(x1+x2)) + math.Sqrt(math.Pow(x1, 2)+math.Pow(x2, 2))
	}

	vals := []float64{5.1, 1.1}
	begin := CreatePoint(vals)

	vals = []float64{0, 0}
	min := CreatePoint(vals)

	return NewFunction(f, begin, min, nil)
}

func FunctionFive() *Function {
	f := func(point *Point) float64 {
		sum := 0.0
		for i := 0; i < point.Dims(); i++ {
			element := point.GetVal(i)
			sum += math.Pow(element, 2)
		}
		return 0.5 + math.Sin(math.Sqrt(sum)) - 0.5/math.Pow(1+0.001*sum, 2)
	}

	vals := []float64{0, 0, 0, 0}
	min := CreatePoint(vals)

	return NewFunction(f, nil, min, nil)
}

func FunctionFirstTask() *Function {
	f := func(point *Point) float64 {
		sum := 0.0
		for i := 0; i < point.Dims(); i++ {
			sum += math.Pow(point.GetVal(i)-3.0, 2)
		}

		return sum
	}

	vals := []float64{10}
	begin := CreatePoint(vals)

	vals = []float64{3}
	min := CreatePoint(vals)

	return NewFunction(f, begin, min, nil)
}

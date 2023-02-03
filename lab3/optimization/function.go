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

	fu.Jacob = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		oneOne := 400 * x1 * (math.Pow(x1, 2) - x2)
		oneTwo := 2*x1 - 2
		twoOne := 200 * (x2 - math.Pow(x1, 2))
		twoTwo := 0.0

		return matrix.NewMatrixFromArray([][]float64{{oneOne, oneTwo}, {twoOne, twoTwo}})
	}

	fu.NonLinear = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		fu.NumCalled++

		first := math.Pow(x2-math.Pow(x1, 2), 2)
		return matrix.NewMatrixFromArray([][]float64{{100 * first}, {math.Pow(1-x1, 2)}})
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
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		return math.Pow(x1-2, 2) + 4*math.Pow(x2+3, 2)
	}

	vals := []float64{0, 0}
	begin := CreatePoint(vals)

	vals = []float64{2, -3}
	min := CreatePoint(vals)

	gradient := func(point *Point) *Point {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		firstDer := 2 * (x1 - 2)
		secondDer := 2 * (x2 + 3)
		return CreatePoint([]float64{firstDer, secondDer})
	}

	fu := NewFunction(f, begin, min, gradient)

	fu.Hess = func(point *Point) *matrix.Matrix {
		oneOne := 2.0
		oneTwo := 0.0
		twoTwo := 2.0

		mat := [][]float64{{oneOne, oneTwo}, {oneTwo, twoTwo}}
		return matrix.NewMatrixFromArray(mat)
	}

	return fu
}

func FunctionFour() *Function {
	f := func(point *Point) float64 {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		first := math.Pow(x1, 4) / 4
		second := math.Pow(x1, 2)
		third := 2 * x1
		fourth := math.Pow(x2-1, 2)

		return first - second + third + fourth
	}

	vals := []float64{0, 0}
	begin := CreatePoint(vals)

	gradient := func(point *Point) *Point {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		firstDer := math.Pow(x1, 3) - 2*x1 + 2
		secondDer := 2 * (x2 - 1)
		return CreatePoint([]float64{firstDer, secondDer})
	}

	fu := NewFunction(f, begin, nil, gradient)

	fu.Hess = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]

		oneOne := 3*math.Pow(x1, 2) - 2
		oneTwo := 0.0
		twoTwo := 2.0

		mat := [][]float64{{oneOne, oneTwo}, {oneTwo, twoTwo}}
		return matrix.NewMatrixFromArray(mat)
	}

	return fu
}

func CustomFunc() *Function {
	var f func(point *Point) float64

	vals := []float64{0, 0}
	begin := CreatePoint(vals)

	fu := NewFunction(f, begin, nil, nil)

	fu.Jacob = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		oneOne := 2 * x1
		oneTwo := -2 * x1
		twoOne := 2 * x2
		twoTwo := 1.0

		return matrix.NewMatrixFromArray([][]float64{{oneOne, oneTwo}, {twoOne, twoTwo}})
	}

	fu.NonLinear = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		fu.NumCalled++

		first := math.Pow(x1, 2) + math.Pow(x2, 2) - 1
		second := x2 - math.Pow(x1, 2)
		return matrix.NewMatrixFromArray([][]float64{{first, second}})
	}

	return fu
}

func SecondCustomFunc() *Function {
	var f func(point *Point) float64

	vals := []float64{0, 0}
	begin := CreatePoint(vals)

	fu := NewFunction(f, begin, nil, nil)

	fu.Jacob = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]

		mapp := make(map[float64]float64, 0)
		mapp[1.0] = 3.0
		mapp[2.0] = 4.0
		mapp[3.0] = 4.0
		mapp[5.0] = 5.0
		mapp[6.0] = 6.0
		mapp[7.0] = 8.0

		jacob := make([][]float64, 0)
		for x := range mapp {
			smth := make([]float64, 0)

			oneOne := math.Pow(math.E, x2*x)
			oneTwo := x * x1 * math.Pow(math.E, x2*x)
			oneThree := 1.0

			smth = append(smth, oneOne)
			smth = append(smth, oneTwo)
			smth = append(smth, oneThree)

			jacob = append(jacob, smth)
		}

		return matrix.NewMatrixFromArray(jacob)
	}

	fu.NonLinear = func(point *Point) *matrix.Matrix {
		x1 := point.DimensionVals[0]
		x2 := point.DimensionVals[1]
		x3 := point.DimensionVals[2]

		mapp := make(map[float64]float64, 0)
		mapp[1.0] = 3.0
		mapp[2.0] = 4.0
		mapp[3.0] = 4.0
		mapp[5.0] = 5.0
		mapp[6.0] = 6.0
		mapp[7.0] = 8.0

		jacob := make([][]float64, 0)
		for x, y := range mapp {
			smth := make([]float64, 0)

			smth = append(smth, x1*math.Pow(math.E, x2*x)+x3-y)

			jacob = append(jacob, smth)
		}

		return matrix.NewMatrixFromArray(jacob)
	}

	return fu
}

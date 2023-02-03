package optimization

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	DimensionVals []float64
}

func CreatePoint(vals []float64) *Point {
	return &Point{DimensionVals: vals}
}

func (p *Point) GetVal(i int) float64 {
	return p.DimensionVals[i]
}

func (p *Point) SetVal(i int, x float64) {
	p.DimensionVals[i] = x
}

func (p *Point) Dims() int {
	return len(p.DimensionVals)
}

func Distance(x1, x2 *Point) float64 {
	sum := 0.0

	for i := 0; i < x1.Dims(); i++ {
		sum += math.Pow(x1.GetVal(i)-x2.GetVal(i), 2)
	}

	return math.Sqrt(sum)
}

func (p *Point) Sub(x2 *Point) *Point {
	return Sub(p, x2)
}

func (p *Point) Add(x2 *Point) *Point {
	return Add(p, x2)
}

func (p *Point) AddElementWise(scalar float64) *Point {
	return AddElementWise(p, scalar)
}

func (p *Point) SubElementWise(scalar float64) *Point {
	return SubElementWise(p, scalar)
}

func (p *Point) MulScalar(scalar float64) *Point {
	return MulScalar(p, scalar)
}

func Sub(x1, x2 *Point) *Point {
	newVals := make([]float64, x1.Dims())
	for i := 0; i < x1.Dims(); i++ {
		newVals[i] = x1.GetVal(i) - x2.GetVal(i)
	}

	return CreatePoint(newVals)
}

func Add(x1, x2 *Point) *Point {
	newVals := make([]float64, x1.Dims())
	for i := 0; i < x1.Dims(); i++ {
		newVals[i] = x1.GetVal(i) + x2.GetVal(i)
	}

	return CreatePoint(newVals)
}

func MulScalar(x *Point, scalar float64) *Point {
	newVals := make([]float64, x.Dims())
	for i := 0; i < x.Dims(); i++ {
		newVals[i] = x.GetVal(i) * scalar
	}

	return CreatePoint(newVals)
}

func AddElementWise(x *Point, scalar float64) *Point {
	newVals := make([]float64, x.Dims())
	for i := 0; i < x.Dims(); i++ {
		newVals[i] = x.GetVal(i) + scalar
	}

	return CreatePoint(newVals)
}

func SubElementWise(x *Point, scalar float64) *Point {
	newVals := make([]float64, x.Dims())
	for i := 0; i < x.Dims(); i++ {
		newVals[i] = x.GetVal(i) - scalar
	}

	return CreatePoint(newVals)
}

func (p *Point) EuclidNorm() float64 {
	return EuclidNorm(p)
}

func EuclidNorm(x *Point) float64 {
	sum := 0.0
	for i := 0; i < x.Dims(); i++ {
		sum += math.Pow(x.GetVal(i), 2)
	}

	return math.Sqrt(sum)
}

func (p *Point) Equals(x *Point) bool {
	if p == nil {
		if x == nil {
			return true
		}

		return false
	}
	cnt := 0
	for i := 0; i < p.Dims(); i++ {
		if math.Abs(p.GetVal(i)-x.GetVal(i)) < 1e-5 {
			cnt++
		}
	}

	return cnt == p.Dims()
}

func (p *Point) ToString() string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, val := range p.DimensionVals {
		sb.WriteString(strconv.FormatFloat(val, 'f', 3, 64) + ", ")
	}
	sb.WriteString("]")
	return sb.String()
}

func (p *Point) Copy() *Point {
	newVals := make([]float64, len(p.DimensionVals))
	copy(newVals, p.DimensionVals)
	return CreatePoint(newVals)
}

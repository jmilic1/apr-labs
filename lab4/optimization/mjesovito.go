package optimization

import (
	"math"
)

func InnerPoint(constraints []func(point *Point) float64, X0 *Point) *Point {
	x := X0
	// for {
	//
	// }

	return x
}

func KaznaBarijera(f Function, constraints []func(point *Point) float64, equalities []func(point *Point) float64, t0 float64, x0 *Point) *Point {
	x := x0
	// t := t0
	// F := f.Call(x)
	// xP := x0
	// xB := x0
	//
	// for {
	// 	xS := x
	// 	var minX *internal.Point
	// 	minF := 0.0
	//
	// 	for
	//
	// 	x = minX
	// 	t = 10 * t
	// }
	return x
}

func F(f Function, constraints []func(point *Point) float64, equalities []func(point *Point) float64, t float64, x *Point) float64 {
	F := f.Call(x)

	for _, constraint := range constraints {
		if constraint(x) <= 0 {
			return math.Inf(1)
		} else {
			F = F - 1/t*math.Log(constraint(x))
		}
	}

	for _, equality := range equalities {
		F = F - t*math.Pow(equality(x), 2)
	}

	return F
}

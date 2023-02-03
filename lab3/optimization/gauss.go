package optimization

import (
	"fmt"

	"0036515065/matrix"
)

func Gauss(x0 *Point, e float64, function *Function, print bool) *Point {
	x := x0

	for {
		jacob := function.Jacob(x)
		G := function.NonLinear(x)

		A := jacob.Transpose().Mul(jacob)
		g := jacob.Transpose().Mul(G)

		delta := matrix.SolveEquationLU(A, g.MulScalar(-1)).MulScalar(-1)

		x = x.Add(CreatePoint([]float64{delta.Get(0, 0), delta.Get(0, 1)}))

		if print {
			fmt.Printf("x = [%f, %f]\n", x.GetVal(0), x.GetVal(1))
		}

		if function.Gradient(x).EuclidNorm() < e {
			break
		}
	}

	return x
}

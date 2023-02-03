package optimization

import (
	"fmt"

	"0036515065/matrix"
)

func Newton(x0 *Point, e float64, function *Function, golden bool, print bool) *Point {
	x := x0

	for {
		if print {
			fmt.Printf("x = [%f, %f]\n", x.GetVal(0), x.GetVal(1))
		}

		p := function.Gradient(x)
		arr := [][]float64{{p.GetVal(0)}, {p.GetVal(1)}}
		mat := matrix.NewMatrixFromArray(arr)

		hess := function.Hess(x)
		inv := matrix.Inverse(hess)
		deltaMat := inv.Mul(mat)

		delta := CreatePoint([]float64{deltaMat.Get(0, 0), deltaMat.Get(1, 0)})

		x = x.Add(delta)

		if function.Gradient(x).EuclidNorm() < e {
			break
		}
	}

	return x
}

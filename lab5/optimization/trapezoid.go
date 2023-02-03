package optimization

import (
	"fmt"

	"0036515065/internal"
)

func Trapezoid(A *internal.Matrix, B *internal.Matrix, beginX *internal.Matrix, funcForT func(t float64) *internal.Matrix, maxT float64, T float64, printEvery int) ([]float64, []float64) {
	x := beginX
	pr := 0

	correctXs := []float64{x.Get(0, 0)}
	correctYs := []float64{x.Get(1, 0)}
	U := internal.IdentityMatrix(A.SizeSquare())

	for i := 0.0; i <= maxT; i += T {
		first := internal.Inverse(U.Sub(A.MulScalar(T / 2)))
		second := U.Plus(A.MulScalar(T / 2))
		R := first.Mul(second)

		x = R.Mul(x)
		if B != nil {
			S := first.MulScalar(T / 2).Mul(B)
			x = x.Plus(S.Mul(funcForT(i).Plus(funcForT(i + T))))
		}

		correctXs = append(correctXs, x.Get(0, 0))
		correctYs = append(correctYs, x.Get(1, 0))

		if pr%printEvery == 0 {
			fmt.Printf("iteration %d: %.2f, %.2f\n", pr, x.Get(0, 0), x.Get(1, 0))
		}
		pr++
	}

	return correctXs, correctYs
}

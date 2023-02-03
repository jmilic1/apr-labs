package optimization

import (
	"fmt"

	"0036515065/internal"
)

func Reverse(A *internal.Matrix, B *internal.Matrix, beginX *internal.Matrix, funcForT func(t float64) *internal.Matrix, maxT float64, T float64, printEvery int) ([]float64, []float64) {
	x := beginX
	pr := 0

	correctXs := []float64{x.Get(0, 0)}
	correctYs := []float64{x.Get(1, 0)}
	U := internal.IdentityMatrix(A.SizeSquare())

	for i := 0.0; i <= maxT; i += T {
		P := internal.Inverse(U.Sub(A.MulScalar(T)))

		x = P.Mul(x)
		if B != nil {
			Q := P.MulScalar(T).Mul(B)
			x = x.Plus(Q.Mul(funcForT(i + T)))
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

package optimization

import (
	"fmt"

	"0036515065/internal"
)

func Runge(A *internal.Matrix, B *internal.Matrix, beginX *internal.Matrix, funcForT func(t float64) *internal.Matrix, maxT float64, T float64, printEvery int) ([]float64, []float64) {
	x := beginX
	pr := 0

	correctXs := []float64{x.Get(0, 0)}
	correctYs := []float64{x.Get(1, 0)}
	for i := 0.0; i <= maxT; i += T {
		m1 := A.Mul(x)
		m2 := A.Mul(x.Plus(m1.MulScalar(T / 2)))
		m3 := A.Mul(x.Plus(m2.MulScalar(T / 2)))
		m4 := A.Mul(x.Plus(m2.MulScalar(T)))

		x = x.Plus(m1.Plus(m2.MulScalar(2)).Plus(m3.MulScalar(2)).Plus(m4).MulScalar(T / 6))

		correctXs = append(correctXs, x.Get(0, 0))
		correctYs = append(correctYs, x.Get(1, 0))

		if pr%printEvery == 0 {
			fmt.Printf("iteration %d: %.2f, %.2f\n", pr, x.Get(0, 0), x.Get(1, 0))
		}
		pr++
	}

	return correctXs, correctYs
}

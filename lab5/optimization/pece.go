package optimization

import (
	"fmt"

	"0036515065/internal"
)

func Pece(A *internal.Matrix, B *internal.Matrix, beginX *internal.Matrix, funcForT func(t float64) *internal.Matrix, maxT float64, T float64, printEvery int) ([]float64, []float64) {
	x := beginX
	pr := 0

	correctXs := []float64{x.Get(0, 0)}
	correctYs := []float64{x.Get(1, 0)}

	for i := 0.0; i <= maxT; i += T {
		temp := A.Mul(x)
		if B != nil {
			temp = temp.Plus(B.Mul(funcForT(i)))
		}
		xCap := x.Plus(temp.MulScalar(T))

		temp = A.Mul(x)
		if B != nil {
			temp = temp.Plus(B.Mul(funcForT(i)))
		}

		temp2 := A.Mul(xCap)
		if B != nil {
			temp2 = temp2.Plus(B.Mul(funcForT(i + T)))
		}

		x = x.Plus(temp.Plus(temp2).MulScalar(T / 2))

		correctXs = append(correctXs, x.Get(0, 0))
		correctYs = append(correctYs, x.Get(1, 0))

		if pr%printEvery == 0 {
			fmt.Printf("iteration %d: %.2f, %.2f\n", pr, x.Get(0, 0), x.Get(1, 0))
		}
		pr++
	}

	return correctXs, correctYs
}

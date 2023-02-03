package optimization

import (
	"fmt"
)

func Gradient(x0 *Point, e float64, function *Function, golden bool, print bool) *Point {
	x := x0

	previous := make([]*Point, 10)
	previous[0] = x
	cntPrevious := 0
	i := 0

	for {
		if print {
			fmt.Printf("x = [%f, %f]\n", x.GetVal(0), x.GetVal(1))
		}

		x = x.Sub(function.Gradient(x))

		if inPrevious(previous, x) {
			cntPrevious++
			if cntPrevious == 10 {
				break
			}
		} else {
			cntPrevious = 0
		}

		previous[i%10] = x

		if function.Gradient(x).EuclidNorm() < e {
			break
		}

		i++
	}

	return x
}

func inPrevious(previous []*Point, curr *Point) bool {
	for _, prev := range previous {
		if prev.Equals(curr) {
			return true
		}
	}
	return false
}

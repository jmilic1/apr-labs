package optimization

import (
	"fmt"
	"math"
	"math/rand"
)

func Box(X0 *Point, alpha, epsilon float64, function *Function, min float64, max float64, constraints []func(point *Point) float64, print bool) *Point {
	n := X0.Dims()
	for i := 0; i < n; i++ {
		if min > X0.GetVal(i) || X0.GetVal(i) > max {
			return nil
		}

		if OutOfBounds(X0, constraints) {
			return nil
		}
	}

	centroid := X0
	simplex := make([]*Point, 0)
	simplex = append(simplex, centroid)

	for t := 1; t < 2*n; t++ {
		vals := make([]float64, n)
		for i := 0; i < n; i++ {
			r := rand.Float64()

			vals[i] = min + r*(max-min)
		}
		newPoint := CreatePoint(vals)
		for OutOfBounds(newPoint, constraints) {
			newPoint = newPoint.Add(centroid).MulScalar(0.5)
		}

		simplex = append(simplex, newPoint)

		// h := -1
		// maxVal := 0.0

		// for i, point := range simplex {
		// 	newVal := function.Call(point)
		// 	if h == -1 || newVal > maxVal {
		// 		h = i
		// 		maxVal = newVal
		// 	}
		// }

		centroid = Centroid(simplex, -1)
	}

	first := true
	var lastCentroid *Point
	for first || Continue(simplex, centroid, function, epsilon) {
		first = false

		h := -1
		h2 := -1

		maxVal := 0.0
		secondMaxVal := 0.0

		for i := 0; i < 2*n; i++ {
			newVal := function.Call(simplex[i])
			if h == -1 || newVal > maxVal {
				if newVal > maxVal {
					h2 = h
					secondMaxVal = maxVal
				}
				h = i
				maxVal = newVal
			}
			if (newVal > secondMaxVal || h2 == -1) && newVal != maxVal {
				h2 = i
				secondMaxVal = newVal
			}
		}

		centroid = Centroid(simplex, h)
		if lastCentroid != nil && centroid.Equals(lastCentroid) {
			break
		}

		Xr := Reflexion(simplex[h], centroid, alpha)

		for i := 0; i < n; i++ {
			if Xr.GetVal(i) < min {
				Xr.DimensionVals[i] = min
			}
			if Xr.GetVal(i) > max {
				Xr.DimensionVals[i] = max
			}
		}

		for OutOfBounds(Xr, constraints) {
			Xr = Xr.Add(centroid).MulScalar(0.5)
		}

		if function.Call(Xr) > function.Call(simplex[h2]) {
			Xr = Xr.Add(centroid).MulScalar(0.5)
		}

		simplex[h] = Xr

		if print {
			fmt.Printf("Centroid: "+centroid.ToString()+", function(centroid) = %f\n", function.Call(centroid))
		}
		lastCentroid = centroid
	}

	return centroid
}

func Centroid(simplex []*Point, ignore int) *Point {
	vals := make([]float64, simplex[0].Dims())
	for i := 0; i < len(simplex); i++ {
		if ignore == i {
			continue
		}
		if simplex[i] == nil {
			continue
		}

		for dim := 0; dim < simplex[0].Dims(); dim++ {
			vals[dim] += simplex[i].GetVal(dim)
		}
	}

	for i := 0; i < simplex[0].Dims(); i++ {
		vals[i] /= float64(len(simplex))
	}

	return CreatePoint(vals)
}

func Reflexion(x, centroid *Point, alpha float64) *Point {
	return centroid.MulScalar(1 + alpha).Sub(x.MulScalar(alpha))
}

func Continue(simplex []*Point, centroid *Point, function *Function, e float64) bool {
	n := float64(centroid.Dims())
	sum := 0.0
	for i := 0; float64(i) < n; i++ {
		sum += math.Pow(function.Call(simplex[i])-function.Call(centroid), 2)
	}
	test := math.Sqrt(sum / n)
	return test > e
}

func OutOfBounds(x *Point, constraints []func(point *Point) float64) bool {
	for _, constraint := range constraints {
		if constraint(x) < 0 {
			return true
		}
	}
	return false
}

package optimization

import (
	"fmt"
	"math"
)

func Simplex(X0 *Point, step, alpha, beta, gamma, sigma, epsilon float64, function *Function, print bool) *Point {
	n := X0.Dims()
	simplex := make([]*Point, n+1)

	for i := 0; i < n; i++ {
		vals := make([]float64, n)
		copy(vals, X0.DimensionVals)
		vals[i] += step
		simplex[i] = CreatePoint(vals)
	}
	vals := make([]float64, n)
	copy(vals, X0.DimensionVals)
	simplex[n] = CreatePoint(vals)

	first := true
	var centroid *Point
	var lastCentroid *Point
	for first || Continue(simplex, centroid, function, epsilon) {
		first = false

		h := -1
		l := -1
		minVal := 0.0
		maxVal := 0.0

		for i := 0; i < n+1; i++ {
			newVal := function.Call(simplex[i])
			if h == -1 || newVal > maxVal {
				h = i
				maxVal = newVal
			}
			if l == -1 || newVal < minVal {
				l = i
				minVal = newVal
			}
		}

		centroid = Centroid(simplex, h)
		if lastCentroid != nil && function.Call(centroid) == function.Call(lastCentroid) {
			break
		}

		Xr := Reflexion(simplex[h], centroid, alpha)
		if function.Call(Xr) < function.Call(simplex[l]) {
			Xe := Expansion(Xr, centroid, gamma)
			if function.Call(Xe) < function.Call(simplex[l]) {
				simplex[h] = Xe
			} else {
				simplex[h] = Xr
			}
		} else {
			flag := true
			for j := 0; j < n+1; j++ {
				if j == h {
					continue
				}

				if function.Call(Xr) <= function.Call(simplex[j]) {
					flag = false
					break
				}
			}

			if flag {
				if function.Call(Xr) < function.Call(simplex[h]) {
					simplex[h] = Xr
				}

				Xk := Contraction(simplex[h], centroid, beta)
				if function.Call(Xk) < function.Call(simplex[h]) {
					simplex[h] = Xk
				} else {
					for j := 0; j < n+1; j++ {
						if j == l {
							continue
						}

						simplex[j] = simplex[l].Add(simplex[j].Sub(simplex[l]).MulScalar(sigma))
					}
				}
			} else {
				simplex[h] = Xr
			}
		}

		if print {
			fmt.Printf("Centroid: "+centroid.ToString()+", function(centroid) = %f\n", function.Call(centroid))
		}
		lastCentroid = centroid
	}

	return centroid
}

func Continue(simplex []*Point, centroid *Point, function *Function, e float64) bool {
	n := float64(centroid.Dims())
	sum := 0.0
	for i := 0; float64(i) < n; i++ {
		sum += n * math.Pow(function.Call(simplex[i])-function.Call(centroid), 2)
	}
	test := math.Sqrt(sum / n)
	return test > e
}

func Centroid(simplex []*Point, ignore int) *Point {
	vals := make([]float64, simplex[0].Dims())
	for i := 0; i < len(simplex); i++ {
		if ignore == i {
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

func Expansion(x, centroid *Point, gamma float64) *Point {
	return centroid.MulScalar(1 - gamma).Sub(x.MulScalar(gamma))
}

func Contraction(x, centroid *Point, beta float64) *Point {
	return centroid.MulScalar(1 - beta).Sub(x.MulScalar(beta))
}

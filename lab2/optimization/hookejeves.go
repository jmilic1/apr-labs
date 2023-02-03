package optimization

import "fmt"

func HookeJeeves(begin *Point, dx float64, function *Function, e float64, print bool) *Point {
	beginSearch := begin
	base := begin

	first := true
	for first || Distance(beginSearch, base) >= e {
		first = false

		searched := search(beginSearch, dx, function)

		if print {
			fmt.Printf("============\nbase: "+base.ToString()+" f(base) = %f\n"+
				"beginSearch: "+beginSearch.ToString()+" f(beginSearch) = %f\n"+
				"searched: "+searched.ToString()+" f(searched) = %f\n", function.Call(base), function.Call(beginSearch), function.Call(searched))
		}

		if function.Call(searched) < function.Call(base) {
			beginSearch = searched.MulScalar(2).Sub(base)
			base = searched
		} else {
			dx /= 2
			beginSearch = base
		}
	}

	return base
}

func search(beginSearch *Point, dx float64, function *Function) *Point {
	x := beginSearch
	n := beginSearch.Dims()

	for i := 0; i < n; i++ {
		P := function.Call(x)

		x = x.AddElementWise(dx)
		N := function.Call(x)
		if N > P {
			x = x.SubElementWise(2 * dx)
			N = function.Call(x)
			if N > P {
				x = x.AddElementWise(dx)
			}
		}
	}

	return x
}

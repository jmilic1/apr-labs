package optimization

func Coordinate(x0 *Point, e float64, h float64, function *Function, print bool) *Point {
	x := x0
	xs := x
	n := x0.Dims()

	first := true

	for first || Distance(x, xs) > e {
		first = false
		xs = x

		for i := 0; i < n; i++ {

			searchFunc := func(lambda *Point) float64 {
				values := make([]float64, x.Dims())
				for i2, val := range x.DimensionVals {
					if i2 == i {
						values[i2] = x.GetVal(i) + lambda.GetVal(0)
					} else {
						values[i2] = val + 1
					}
				}
				return function.Call(CreatePoint(values))
			}

			lambdaMin := GoldenDoUnimodal(h, x, NewFunction(searchFunc, nil, nil, nil), e, print)

			x.SetVal(i, x.GetVal(i)+lambdaMin)
		}
	}

	return x
}

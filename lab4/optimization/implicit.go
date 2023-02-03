package optimization

func OneOne(point *Point) float64 {
	x1 := point.DimensionVals[0]
	x2 := point.DimensionVals[1]

	return x2 - x1
}

func OneTwo(point *Point) float64 {
	x1 := point.DimensionVals[0]

	return 2 - x1
}

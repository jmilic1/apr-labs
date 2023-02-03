package optimization

import (
	"fmt"
	"math"
)

func Unimodal(h float64, x *Point, function *Function) (*Point, *Point) {
	l := x.SubElementWise(h)
	r := x.AddElementWise(h)

	m := x
	var fl, fm, fr float64
	step := 1

	fm = function.Call(x)
	fl = function.Call(l)
	fr = function.Call(r)

	if fm < fr && fm < fl {
		return l, r
	} else if fm > fr {
		for fm > fr {
			l = m
			m = r
			fm = fr
			step = step * 2
			r = x.AddElementWise(h * float64(step))
			fr = function.Call(r)
		}
	} else {
		for fm > fl {
			r = m
			m = l
			fm = fl
			step = step * 2
			l = x.SubElementWise(h * float64(step))
			fl = function.Call(l)
		}
	}

	return l, r
}

func GoldenDoUnimodal(h float64, x *Point, function *Function, e float64, print bool) float64 {
	a, b := Unimodal(h, x, function)

	return Golden(a, b, function, e, print)
}

func Golden(a *Point, b *Point, function *Function, e float64, print bool) float64 {
	k := K()

	c := b.Sub(b.Sub(a).MulScalar(k))
	d := a.Add(b.Sub(a).MulScalar(k))

	fc := function.Call(c)
	fd := function.Call(d)

	for b.Sub(a).GetVal(0) > e {
		if fc < fd {
			b = d
			d = c
			c = b.Sub(b.Sub(a).MulScalar(k))
			fd = fc
			fc = function.Call(c)
		} else {
			a = c
			c = d
			d = a.Add(b.Sub(a).MulScalar(k))
			fc = fd
			fd = function.Call(d)
		}

		if print {
			fmt.Printf("a = %f, b = %f, c = %f, d = %f\n", a.GetVal(0), b.GetVal(0), c.GetVal(0), d.GetVal(0))
		}
	}

	return a.Add(b).GetVal(0) / 2
}

func K() float64 {
	return 0.5 * (math.Sqrt(5) - 1)
}

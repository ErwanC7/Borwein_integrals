package main

import (
	"math"
	"os"
)

func f(n int, x float64) float64 {
	var result = 1.0
	var num = 0.0

	for k := 0.0; int(k) <= n; k++ {
		num = 2.0 * k + 1.0
		if num == 0  || x / num == 0 {
			os.Exit(84)
		}
		result *= (math.Sin(x / num)) / (x / num)
	}
	return result
}

func midpoint(n int) float64 {
	var result = 0.0
	var b = 0.0

	for a := 0.0; a < 5000; a += 0.5 {
		b = a + 0.5
		result += (b - a) * f(n, (a + b) / 2.0)
	}
	return result
}

func trapezoid(n int) float64 {
	var result = 0.0
	var b = 0.0

	for a := 0.000000000000001; a < 5000; a += 0.5 {
		b = a + 0.5
		result += ((b - a) / 2) * (f(n, a) + f(n, b))
	}
	return result
}

func simpson(n int) float64 {
	var result = 0.0
	var b = 0.0

	for a := 0.000000000000001; a < 5000; a += 0.5 {
		b = a + 0.5
		result += ((b - a) / 6) * (f(n, a) + 4 * f( n, (a + b) / 2) + f(n, b))
	}
	return result
}
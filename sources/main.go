/*
## EPITECH PROJECT, 2021
## main
## File description:
## main of 110borwein
 */

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func displayHelp() int {
	fmt.Println(
		"USAGE\n" +
		"    ./110borwein n\n\n" +
		"DESCRIPTION\n" +
		"   n    constant defining the integral to be computed")
	return 0
}

func absZero(nbr float64) float64 {
	if 0.0 > nbr && nbr > -0.00000000009 {
		nbr *= -1
	}
	return nbr
}

func printResult(n int, mid float64, trap float64, simpson float64) {
	var pi2 = math.Pi / 2
	var diffMid = absZero(pi2 - mid)
	var diffTrap = absZero(pi2 - trap)
	var diffSimpson = absZero(pi2 - simpson)

	fmt.Println("Midpoint:")
	fmt.Printf("I%d = %.10f\n", n, mid)
	fmt.Printf("diff = %.10f\n\n", diffMid)
	fmt.Println("Trapezoidal:")
	fmt.Printf("I%d = %.10f\n", n, trap)
	fmt.Printf("diff = %.10f\n\n", diffTrap)
	fmt.Println("Simpson:")
	fmt.Printf("I%d = %.10f\n", n, simpson)
	fmt.Printf("diff = %.10f\n", diffSimpson)
}

func main() {
	var n = 0
	var _ = 0

	if len(os.Args) != 2 {
		os.Exit(84)
	}
	if os.Args[1] == "-h" {
		os.Exit(displayHelp())
	}
	for i := 0; i < len(os.Args[1]); i++ {
		if '0' > os.Args[1][i] || os.Args[1][i] > '9' {
			os.Exit(84)
		}
	}
	n, _ = strconv.Atoi(os.Args[1])
	printResult(n, midpoint(n), trapezoid(n), simpson(n))
}

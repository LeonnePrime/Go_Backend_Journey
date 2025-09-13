package main

import (
	"fmt"
)

// slope calculates the gradient (m) of the line through two points
func slope(x1, y1, x2, y2 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

// yIntercept calculates the y-intercept (b) of the line
func yIntercept(y, m, x float64) float64 {
	return y - m*x
}

func main() {
	for {
		var x1, y1, x2, y2 float64

		// Ask user to enter two points
		fmt.Println("Enter coordinates of first point (x1 y1):")
		fmt.Scan(&x1, &y1)

		fmt.Println("Enter coordinates of second point (x2 y2):")
		fmt.Scan(&x2, &y2)

		// Calculate slope and y-intercept
		m := slope(x1, y1, x2, y2)
		b := yIntercept(y1, m, x1)

		// Print results with proper formatting
		fmt.Printf("\nSlope (gradient) = %.2f\n", m)
		fmt.Printf("Y-intercept = %.2f\n", b)

		// Print equation neatly with + or - for intercept
		if b >= 0 {
			fmt.Printf("Equation of the line: y = %.2fx + %.2f\n", m, b)
		} else {
			fmt.Printf("Equation of the line: y = %.2fx - %.2f\n", m, -b)
		}
	}
}

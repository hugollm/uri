package main

import (
	"fmt"
	"math"
)

const PI = 3.14159;

func main() {
	var area, radius float64;
	fmt.Scanf("%f", &radius)
	area = PI * math.Pow(radius, 2)
	fmt.Printf("A=%.4f\n", area)
}

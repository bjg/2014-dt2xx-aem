package main

import (
	"fmt"
	"math"
)

var results = []float64{
	70.21067, 95.23731, 48.56398, 45.80380, 84.57824, 57.56347, 63.98469, 76.35332,
    76.61905, 78.08247, 58.86418, 52.45476, 41.10990, 74.44590, 60.94644, 69.56891,
    82.40521, 94.88734, 78.25650, 73.42231, 64.45375, 83.20946, 73.53518, 58.06571,
    28.38046, 52.63720, 80.55414, 79.74410, 95.17229, 61.29581,
}

func main() {
	fmt.Println(isort(results))
	fmt.Println(mean(results), sd(results), median(results))
}

func mean(dist []float64) float64 {
	var sum float64
	for _, x := range dist {
		sum += x
	}
	return sum / float64(len(dist))
}

func sd(dist []float64) float64 {
	// 0. calculate x-bar (i.e. mean)
	xBar := mean(dist)

	// 1. sum of the square differences
	var sum float64
	for _, x := range dist {
		d := x - xBar
		sum += d * d
	}

	// 2. divide by n - 1
	sum /= float64(len(dist) - 1)

	// 3. square root of result 2
	return math.Sqrt(sum)

}

func isort(dist []float64) []float64 {
	for i := 1; i < len(dist); i++ {
		for j := i; j > 0 && dist[j-1] > dist[j]; j-- {
			dist[j-1], dist[j] = dist[j], dist[j-1]
		}
	}
	return dist
}

func median(dist []float64) float64 {
	sorted := isort(dist)
	return sorted[len(sorted) / 2]
}
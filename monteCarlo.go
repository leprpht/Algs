import (
	"fmt"
	"math/rand"
)

func monteCarloSqrt(n float64, samples int) float64 {
	if n < 0 {
		fmt.Println("n must be positive")
		return -1
	}

	side := n
	if side < 1 {
		side = 1
	}

	inside := 0

	for i := 0; i < samples; i++ {
		x := rand.Float64() * side
		y := rand.Float64() * side

		if y <= x*x && x*x <= n {
			inside++
		}
	}

	estimatedArea := (float64(inside) / float64(samples)) * (side * side)
	result := (3 * estimatedArea) / n
	fmt.Printf("Estimated sqrt(%f) = %f\n", n, result)

	return result
}
import (
	"fmt"
	"math/rand"
)

func monteCarlo(samples int) float64 {
	insideCircle := 0

	for i := 0; i < samples; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x*x+y*y <= 1 {
			insideCircle++
		}
	}

	return 4.0 * float64(insideCircle) / float64(samples)
}
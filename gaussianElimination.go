import "fmt"

func gaussianElimination(a [][]float64) []float64 {
	n := len(a)

	for i := 0; i < n; i++ {

		divider := a[i][i]
		for j := 0; j <= n; j++ {
			a[i][j] /= divider
		}

		for k := i + 1; k < n; k++ {
			factor := a[k][i]

			for j := 0; j <= n; j++ {
				a[k][j] -= factor * a[i][j]
			}
		}
	}

	x := make([]float64, n)

	for i := n - 1; i >= 0; i-- {
		x[i] = a[i][n]

		for j := i + 1; j < n; j++ {
			x[i] -= a[i][j] * x[j]
		}
	}

	for i, v := range x {
		fmt.Printf("x%d = %.2f\n", i+1, v)
	}

	return x
}
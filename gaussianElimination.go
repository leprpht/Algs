import "fmt"

func determinant(matrix [][]float64) float64 {
	n := len(matrix)
	det := 1.0

	for i := 0; i < n; i++ {

		pivot := matrix[i][i]

		if pivot == 0 {
			return 0
		}

		for k := i + 1; k < n; k++ {
			factor := matrix[k][i] / pivot

			for j := i; j < n; j++ {
				matrix[k][j] -= factor * matrix[i][j]
			}
		}

		det *= pivot
	}

	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%8.2f ", value)
		}
		fmt.Println()
	}

	return det
}
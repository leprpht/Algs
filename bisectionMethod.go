func f(x float64) float64 {
	return x*x*x - x - 2
}

func bisection(a, b float64, iterations int) float64 {
	if f(a)*f(b) >= 0 {
		fmt.Printf("invalid interval")
		return 0
	}

	var c float64

	for i := 0; i < iterations; i++ {
		c = (a + b) / 2

		fmt.Printf("Iteration %d: x = %.10f\n", i+1, c)

		if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
	}

	return c
}
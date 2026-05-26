func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 5; i <= limit; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func findPrimeQuadruplets(limit int) {
	for p := 2; p <= limit-8; p++ {
		if isPrime(p) &&
			isPrime(p+2) &&
			isPrime(p+6) &&
			isPrime(p+8) {

			fmt.Printf("(%d, %d, %d, %d)\n", p, p+2, p+6, p+8)
		}
	}
}
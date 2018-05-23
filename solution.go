package solution

var bankcoin = []float64{1000, 500, 100, 50, 20, 10, 5, 2, 1, 0.5, 0.25}

func Change(price float64, paid float64) map[float64]int {
	c := paid - price
	if c < 0 {
		return map[float64]int{}
	}

	r := map[float64]int{}

	for _, bc := range bankcoin {
		if c/bc >= 1 {
			r[bc] = int(c / bc)
			c -= (c / bc)
		}
	}

	return r
}

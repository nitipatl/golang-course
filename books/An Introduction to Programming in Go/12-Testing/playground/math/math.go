package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Max(xs []float64) (max float64) {
	max = -9999
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return
}

func Min(xs []float64) (min float64) {
	min = 9999
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return
}

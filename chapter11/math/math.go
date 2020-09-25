package math

// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return total
	} else {
		for _, x := range xs {
			total += x
		}
		return total / float64(len(xs))
	}

}

//Найти минимальное число в массиве чисел
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		var total float64 = 0
		return total
	} else {
		var minElement float64 = xs[1]
		for _, value := range xs {
			if minElement > value || minElement == value {
				minElement = value
			}
		}
		return minElement
	}
}

//Найти максиммальное число в массиве чисел
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		var total float64 = 0
		return total
	} else {
		var maxElement float64 = xs[1]
		for _, value := range xs {
			if maxElement < value || maxElement == value {
				maxElement = value
			}
		}
		return maxElement
	}
}

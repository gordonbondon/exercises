package user_growth

import (
	"math"
)

func GetBillionUsersDay(target float64, arr []float64) int {
	min := 1
	max := 1

	current := float64(0)

	for current < target {
		current = sum(arr, max)

		max = max * 2
	}

	var t int

	for {
		if max-min == 1 {
			// edge case, check if we need to go to max
			if current < target {
				t++
			}

			break
		}

		t = min + ((max - min) / 2)

		current = sum(arr, t)

		if current > target {
			max = t

			continue
		} else if current < target {
			min = t

			continue
		}

		break
	}

	return t
}

func sum(arr []float64, t int) float64 {
	sum := 0.0
	for _, g := range arr {
		sum = sum + math.Pow(g, float64(t))
	}

	return sum
}

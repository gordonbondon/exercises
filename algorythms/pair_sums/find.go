package pair_sums

func FindNumberOfWays(arr []int, k int) int {
	count := 0

	remainders := make(map[int]int)

	for _, n := range arr {
		if n > k {
			continue
		}

		if v, ok := remainders[n]; ok {
			count = count + v
		}

		remainder := k - n

		if _, ok := remainders[remainder]; ok {
			remainders[remainder]++
		} else {
			remainders[remainder] = 1
		}
	}

	return count
}

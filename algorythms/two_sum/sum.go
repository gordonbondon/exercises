package two_sum

func TwoSum(nums []int, target int) []int {
	result := make([]int, 2, 2)

	// map leftover to index
	leftovers := make(map[int]int, len(nums))

	for i, num := range nums {
		if v, ok := leftovers[num]; ok && v != i {
			result[0] = v
			result[1] = i

			return result
		}

		leftovers[target-num] = i
	}

	return result
}

func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)

	for i, n := range nums {
		needed := target - n
		if index, ok := indexes[needed]; ok {
			return []int{index, i}
		}
		indexes[n] = i
	}

	return []int{-1, -1}
}

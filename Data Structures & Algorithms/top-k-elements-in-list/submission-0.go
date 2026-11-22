func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n]++
	}

	for n, c := range count {
		freq[c] = append(freq[c], n)
	}

	var result []int

	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}

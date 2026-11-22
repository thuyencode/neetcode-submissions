func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, r := range str {
			count[r-'a']++
		}
		group[count] = append(group[count], str)
	}

	var result [][]string

	for _, list := range group {
		result = append(result, list)
	}

	return result
}

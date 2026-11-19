func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}

	countS:=make(map[rune]int)
	countT:=make(map[rune]int)

	for i, r:=range s{
		countS[r]++
		countT[rune(t[i])]++
	}

	for _,r:=range s{
		if countS[r]!=countT[r]{
			return false
		}
	}

	return true
}

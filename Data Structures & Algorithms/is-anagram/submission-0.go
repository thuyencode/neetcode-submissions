func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}

	seen:=make(map[rune]int)
	for _,r:=range s{
		seen[r]++
	}

	for r1,expected:=range seen{
		actual:=0

		for _,r2:=range t{
			if r2==r1{
				actual++
			}
		}

		if expected!=actual{
			return false
		}
	}

	return true
}

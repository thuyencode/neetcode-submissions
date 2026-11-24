type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := new(strings.Builder)
	for _,str:=range strs{
		fmt.Fprintf(encoded,"%d#%s",len(str),str)
	}
	return encoded.String()
}

func (s *Solution) Decode(encoded string) []string {
	var decoded []string

	i:=0
	for i < len(encoded){
		j := i
		for encoded[j] != '#'{
			j++
		}

		length,_ := strconv.Atoi(encoded[i:j])
		i = j + 1
		decoded = append(decoded,encoded[i:i+length])
		i += length
	}

	return decoded
}

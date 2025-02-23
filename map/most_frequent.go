package map_functions

func MostFrequent(s []string) string {
	frequency := make(map[string]int)
	maxFrequency := 0
	mostFrequent := ""

	for _, v := range s {
		frequency[v]++
		if frequency[v] > maxFrequency {
			maxFrequency = frequency[v]
			mostFrequent = v
		}
	}

	return mostFrequent
}

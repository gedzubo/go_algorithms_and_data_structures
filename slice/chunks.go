package slice_functions

// Chunk divides a slice into chunks of a given size.
func Chunk(s []string, chunkSize int) [][]string {
	chunks := make([][]string, 0)
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

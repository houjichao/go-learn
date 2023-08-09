package demo1

func Chunk[E any](s []E, chunkSize int) [][]E {
	var chunks [][]E
	if len(s) == 0 {
		return chunks
	}
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

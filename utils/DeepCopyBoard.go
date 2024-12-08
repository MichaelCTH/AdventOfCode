package utils

func DeepCopyBoard(source [][]string) [][]string {
	destination := make([][]string, len(source))
	for i, row := range source {
		if row == nil {
			destination[i] = nil
			continue
		}
		destination[i] = make([]string, len(row))
		copy(destination[i], row)
	}
	return destination
}

package utils

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func RangeOverlap(range_x []int, range_y []int) bool {
	// to be above parts indexes need to be within range_x - 1 -> range_x + 1
	// (StartA <= EndB) and (EndA >= StartB)
	return (range_x[0] <= range_y[1]) && (range_x[1] >= range_y[0])
}

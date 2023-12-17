package utils

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinArray(a []int) (i int) {
    i = a[0]
    for _, v := range a {
        if v < i {
            i = v
        }
    }
    return
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

func SetIntersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}

	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
        _, prs := hash[e]
		if prs {
			inter = append(inter, e)
		}
	}

	return
}

func Power(base, exp int) (value int) {
    if exp == 0 {
        value = 1
    } else {
        value = base * Power(base, exp - 1)
    }

    return
}

package utils

import (
	"math"
)

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
		value = base * Power(base, exp-1)
	}

	return
}

func QuadraticFormula(a, b, c int) (l, r float64) {
	l = (float64(-1*b) + math.Sqrt(float64(Power(b, 2))-float64(4*a*c))) / float64(2*a)
	r = (float64(-1*b) - math.Sqrt(float64(Power(b, 2))-float64(4*a*c))) / float64(2*a)
	return
}

func QuantifyString(s string) (m map[rune]int) {
	m = make(map[rune]int)
	for _, c := range s {
		_, prs := m[c]
		if prs {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	return
}

func GreatestCommonDenominator(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func LeastCommonMultiple(a, b int) int {
	gcd := GreatestCommonDenominator(a, b)
	prod := a * b
	return int(math.Floor(float64(prod) / float64(gcd)))
}

func LCMMultiple(nums []int) int {
    if len(nums) == 2 {
        return LeastCommonMultiple(nums[0], nums[1])
    } else {
        return LeastCommonMultiple(nums[0], LCMMultiple(nums[1:]))
    }
}

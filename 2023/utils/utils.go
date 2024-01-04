package utils

import (
	"fmt"
	"math"
)

// Math Helper Functions

func Abs(a int) int {
	if a < 0 {
		a = a * -1
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinArray(a []int) (min_value, index int) {
	min_value = a[0]
	index = 0
	for i, value := range a {
		if value < min_value {
			min_value = value
			index = i
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

// Array Helper functions
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

func Diff(nums []int) (diff []int) {
	for i := 0; i < len(nums)-1; i++ {
		diff = append(diff, nums[i+1]-nums[i])
	}
	return
}

func All(values []int) bool {
	for _, v := range values {
		if v == 0 {
			return false
		}
	}
	return true
}

func None(values []int) bool {
	for _, v := range values {
		if v != 0 {
			return false
		}
	}
	return true
}

func Sort(values []int) (sorted []int) {
	i := 0
	count := len(values)
	for i < count {
		min_value, index := MinArray(values)
		sorted = append(sorted, min_value)
		values = append(values[:index], values[index+1:]...)
		i++
	}
	return
}

func Sum(values []int) (total int) {
	for _, v := range values {
		total = total + v
	}
	return
}

func CompareArrays[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareArraysWithDifference[T comparable](a, b []T) int {
	difference := 0
	if len(a) != len(b) {
		difference = difference + Max(len(a), len(b)) - Min(len(a), len(b))
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			difference++
		}
	}
	return difference
}

func Reverse[T any](array []T) (new_array []T) {
	for i := len(array) - 1; i >= 0; i-- {
		new_array = append(new_array, array[i])
	}
	return
}

func FindIndex[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func Clone[T any](array []T) (clone []T) {
	for _, v := range array {
		clone = append(clone, v)
	}
	return
}

func Range(start, end, interval int) (r []int) {
	for i := start; i < end; i = i + interval {
        r = append(r, i)
	}
    return r
}

// Matrix Helper Functions

func PrintMatrix[T any](matrix [][]T) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func RotateMatrixClockwise[T any](matrix [][]T) (rotated [][]T) {
	for j := range matrix[0] {
		new_row := []T{}
		for i := len(matrix) - 1; i >= 0; i-- {
			new_row = append(new_row, matrix[i][j])
		}
		rotated = append(rotated, new_row)
	}
	return
}

func Transpose[T any](matrix [][]T) (transposed [][]T) {
	// Convert matrix so that columns are now rows
	for col_i := range matrix[0] {
		new_row := []T{}
		for row_i := range matrix {
			new_row = append(new_row, matrix[row_i][col_i])
		}
		transposed = append(transposed, new_row)
	}
	return
}

func CompareMatrix[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for row_i := 0; row_i < len(a); row_i++ {
		if !CompareArrays[T](a[row_i], b[row_i]) {
			return false
		}
	}
	return true
}

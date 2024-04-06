package dsa

import (
	"math"
)

// MergeSort O(n * log(n))
func MergeSort(v []int) []int {

	if len(v) == 1 {
		return v
	}

	q1 := 0
	q2 := int(math.Ceil(float64(len(v)) / 2))
	q3 := len(v)

	v1 := v[q1:q2]
	v2 := v[q2:q3]

	v1 = MergeSort(v1)
	v2 = MergeSort(v2)

	return merge(v1, v2)
}

// Merge the two halves of the mergeSort.
func merge(v1 []int, v2 []int) []int {
	var v3 []int

	v1pos, v2pos := 0, 0 // keeps track of the current index at which the next element should be picked

	// If both halves contain elements
	for len(v1) > v1pos && len(v2) > v2pos {
		if v1[v1pos] < v2[v2pos] {
			v3 = append(v3, v1[v1pos])
			v1pos++
		} else {
			v3 = append(v3, v2[v2pos])
			v2pos++
		}
	}

	// If right is empty
	for len(v1) > v1pos {
		v3 = append(v3, v1[v1pos])
		v1pos++
	}

	// If left is empty
	for len(v2) > v2pos {
		v3 = append(v3, v2[v2pos])
		v2pos++
	}

	return v3
}

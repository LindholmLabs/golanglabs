package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	v := []int{9, 8, 6, 3, 1, 0, 11, 4}
	printArr(v)
	printArr(mergeSort(v))
}

// MergeSort O(n * log(n))
func mergeSort(v []int) []int {

	if len(v) == 1 {
		return v
	}

	q1 := 0
	q2 := int(math.Ceil(float64(len(v)) / 2))
	q3 := len(v)

	v1 := v[q1:q2]
	v2 := v[q2:q3]

	v1 = mergeSort(v1)
	v2 = mergeSort(v2)

	return merge(v1, v2)
}

// Merge the two halves of the mergeSort.
func merge(v1 []int, v2 []int) []int {
	var v3 []int

	v1pos, v2pos := 0, 0 // keeps track of the current index at which the next element should be picked

	// If both halves contain elements
	for len(v1) > v1pos && len(v2) > v2pos {
		if v1[0] < v2[0] {
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

// print an []int array
func printArr(v []int) {

	var buffer bytes.Buffer
	buffer.WriteString("{")

	for i, n := range v {
		buffer.WriteString(fmt.Sprintf("%d", n))
		if i < len(v)-1 {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("}")

	fmt.Println(buffer.String())
}

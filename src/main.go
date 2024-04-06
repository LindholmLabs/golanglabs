package main

import (
	"bytes"
	"fmt"

	"github.com/LindholmLabs/golanglabs/src/dsa"
)

func main() {
	slice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	PrintArr(slice)
	slice = dsa.MergeSort(slice)
	PrintArr(slice)
	slice = dsa.FisherYates(slice)
	PrintArr(slice)
}

// print an []int array
func PrintArr(v []int) {

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

package main

import (
	"bytes"
	"fmt"

	"github.com/LindholmLabs/golanglabs/src/dsa"
)

func main() {
	slice := []int{1, 2, 3, 4}
	dsa.MergeSort(slice)
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

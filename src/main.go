package main

import (
	"bytes"
	"fmt"
	"github.com/LindholmLabs/golanglabs/src/dsa"
)

func main() {

	var tree dsa.Tree

	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	fmt.Println("constructed tree, view it using the debugger.")
}

// PrintArr print an []int array
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

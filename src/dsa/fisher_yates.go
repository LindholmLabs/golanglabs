package dsa

import (
	"math/rand/v2"
)

func FisherYates(s []int) []int {
	for i := int32(len(s) - 1); i > 0; i-- {
		r := rand.Int32N(i)
		s[i], s[r] = switchPos(s[i], s[r])
	}

	return s
}

func switchPos(x int, y int) (int, int) {
	z := x
	x = y
	y = z

	return x, y
}

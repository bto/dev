package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized")
}

func Area(len, wid float64) (area float64) {
	area = len * wid
	return
}

func Diagonal(len, wid float64) (diagonal float64) {
	diagonal = math.Sqrt((len * len) + (wid * wid))
	return
}

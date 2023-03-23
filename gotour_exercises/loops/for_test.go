package forloop

import (
	"log"
	"math"
	"testing"
)

func TestMain(m *testing.M) {
	z := 10.0
	x := sqrtOf(z)
	y := math.Sqrt(z)
	if x != y {
		log.Fatal(x, y)
	}
	z = 9.0
	x = sqrtOf(z)
	y = math.Sqrt(z)
	if x != y {
		log.Fatal(x, y)
	}

}

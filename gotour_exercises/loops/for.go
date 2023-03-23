package forloop

import "fmt"

var pl = fmt.Println

func sqrtOf(x float64) float64 {
	z := 1.0
	for (z*z > x+0.0000000001) || (z*z < x-0.00000000001) {
		z -= (z*z - x) / (2 * z)
		pl(z)
	}
	return z
}

func main() {
	pl(sqrtOf(2))
}

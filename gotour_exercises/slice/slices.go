package slice

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for x := range slice {
		slice[x] = make([]uint8, dx)
		for i := 0; i < dy; i++ {
			slice[x][i] = uint8(((x - i) ^ i) - (i*i*i)/(x+1))
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}

package maps

import "strings"

func WordCount(s string) map[string]int {

	arr := strings.Split(s, " ")
	words := make(map[string]int)

	for _, y := range arr {
		words[y]++
	}

	return words
}

func main() {
	wc.Test(WordCount)
}

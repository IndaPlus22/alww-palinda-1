package tps

import "testing"

func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

func Test1(t *testing.T) {

	arr := []int{1, -1, 2, -2, 3, -3}
	expected := 0

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}

}
func Test2(t *testing.T) {

	arr := []int{0}
	expected := 0

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}

}

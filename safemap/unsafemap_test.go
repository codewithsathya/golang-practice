package safemap

import "testing"

func TestUnsafeMap(t *testing.T) {
	data := make(map[int]int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			data[i] = i
		}(i)
	}
}
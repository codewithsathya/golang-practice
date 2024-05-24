package safemap

import "testing"

func TestSafeMap(t *testing.T) {
	m := NewSafeMap[int, int]()
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Insert(i, 2 * i)
			value, err := m.Get(i)
			if err != nil {
				t.Error(err)
			}
			if value != 2 * i {
				t.Errorf("Invalid value. Expected: %v, Got: %v", 2 * i, value)
			}
		}(i)
	}
}
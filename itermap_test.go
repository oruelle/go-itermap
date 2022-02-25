package itermap

import (
	"testing"
)

func TestIndexOf(t *testing.T) {

	for lenght := 1; lenght < 10; lenght++ {
		var v []interface{}

		// Test for empty slice
		i, ok := IndexOf(v, lenght-1)

		if i != 0 || ok != false {
			t.Errorf("Bad index for empty slice for len %d", len(v))
		}

		// Test for different lenght of slice
		for it := 0; it < lenght; it++ {
			v = append(v, it)
		}

		i, ok = IndexOf(v, lenght-1)

		// Check result
		if i != lenght-1 || ok == false {
			t.Errorf("Bad index of slice for len %d %v %d %d %v", len(v), v, lenght-1, i, ok)
		}
	}
}

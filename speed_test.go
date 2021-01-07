package converter

import (
	"testing"
)

func TestKnot(t *testing.T) {
	KnotKm := []struct {
		x float64
		y float64
	}{
		{1, 1.852},
		{-1, -1.852},
		{17, 31.484},
	}
	for _, table := range KnotKm {
		total := KnottoKilomerh(table.x)
		if total != table.y {
			t.Errorf("Unexpected %f was expected %f", total, table.y)
		}
	}

}

package ord

import (
	"fmt"
	"testing"
)

func TestVersionOrd(t *testing.T) {
	// Test cases for the VersionOrd
	testCases := []struct {
		x, y Version
		want int
	}{
		{Version{"1.0.0"}, Version{"1.0.0"}, 0},
		{Version{"1.0.0"}, Version{"2.0.0"}, -1},
		{Version{"2.0.0"}, Version{"1.0.0"}, 1},
	}

	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf("%s vs %s", tc.x.Num, tc.y.Num),
			func(t *testing.T) {
				got := VersionOrd.Compare(tc.x, tc.y)
				if got != tc.want {
					t.Errorf("VersionOrd.Compare(%v, %v) = %d; want %d", tc.x, tc.y, got, tc.want)
				}
			},
		)
	}
}

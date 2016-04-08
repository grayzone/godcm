package ofstd

import (
	"testing"
)

func TestCheck32BitAddOverflow(t *testing.T) {
	cases := []struct {
		in_1 uint32
		in_2 uint32
		want bool
	}{
		{uint32(0), uint32(0), false},
	}
	for _, c := range cases {
		got := Check32BitAddOverflow(c.in_1, c.in_2)
		if got != c.want {
			t.Errorf("Check32BitAddOverflow(), want %q got %q", c.want, got)
		}
	}
}

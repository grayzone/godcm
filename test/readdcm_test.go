package test

import "testing"

func TestDemo(t *testing.T) {
	cases := []struct {
		want int
	}{
		{1},
	}

	for _, c := range cases {

		got := Demo()
		if got != c.want {
			t.Errorf("Demo(), want '%v' got '%v'", c.want, got)
		}
	}

}

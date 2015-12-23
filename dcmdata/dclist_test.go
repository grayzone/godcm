package dcmdata

import "testing"

func TestEmpty(t *testing.T) {
	cases := []struct {
		want bool
	}{
		{true},
	}

	for _, c := range cases {
		var l DcmList
		t.Log(l)
		got := l.Empty()
		if got != c.want {
			t.Errorf("Empty(), want '%v' got '%v'", c.want, got)
		}
	}

}

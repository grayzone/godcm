package ofstd

import (
	"testing"
)

func TestofstdString(t *testing.T) {
	cases := []struct {
		in   OFStatus
		want string
	}{
		{OF_ok, "OF_ok"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("String() == want %v got %v", c.want, got)
		}
	}
}

func TestNewOFCondition(t *testing.T) {
	cases := []struct {
		in_m uint16
		in_c uint16
		in_s OFStatus
		in_t string
		want *OFCondition
	}{
		{0, 0, OF_ok, "test", &OFCondition{0, 0, OF_ok, "test"}},
	}
	for _, c := range cases {
		got := NewOFCondition(c.in_m, c.in_c, c.in_s, c.in_t)
		if *got != *c.want {
			t.Errorf("NewOFCondition() == want %v got %v", c.want, got)
		}
	}

}

func TestOFConditionModule(t *testing.T) {
	cases := []struct {
		in   OFCondition
		want uint16
	}{
		{OFCondition{}, 0},
	}
	for _, c := range cases {
		got := c.in.Module()
		if got != c.want {
			t.Errorf("Module() == want %v got %v", c.want, got)
		}
	}

}

func TestOFConditionCode(t *testing.T) {
	cases := []struct {
		in   OFCondition
		want uint16
	}{
		{OFCondition{}, 0},
	}
	for _, c := range cases {
		got := c.in.Code()
		if got != c.want {
			t.Errorf("Code() == want %v got %v", c.want, got)
		}
	}

}

func TestOFConditionStatus(t *testing.T) {
	cases := []struct {
		in   OFCondition
		want OFStatus
	}{
		{OFCondition{}, OF_ok},
	}
	for _, c := range cases {
		got := c.in.Status()
		if got != c.want {
			t.Errorf("Status() == want %v got %v", c.want, got)
		}
	}
}

func TestOFConditionText(t *testing.T) {
	cases := []struct {
		in   OFCondition
		want string
	}{
		{OFCondition{}, ""},
	}
	for _, c := range cases {
		got := c.in.Text()
		if got != c.want {
			t.Errorf("Text() == want %v got %v", c.want, got)
		}
	}
}

func TestOFConditionGood(t *testing.T) {
	cases := []struct {
		in   OFCondition
		want bool
	}{
		{OFCondition{}, true},
	}
	for _, c := range cases {
		got := c.in.Good()
		if got != c.want {
			t.Errorf("Good() == want %v got %v", c.want, got)
		}
	}
}

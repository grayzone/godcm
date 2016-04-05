package dcmdata

import (
	"strings"
	"testing"
)

func TestNewDcmTagKey(t *testing.T) {
	var v = NewDcmTagKey()
	if v.Group != 0xffff {
		t.Error("excepted 0xffff, got ", v.Group)
	}
	if v.Element != 0xffff {
		t.Error("excepted 0xffff, got ", v.Element)
	}
}

func BenchmarkHasValidGroup(b *testing.B) {
	for i := 0x0000; i <= 0xFFFF; i++ {
		var v DcmTagKey
		v.Group = uint16(i)
		switch v.Group {
		case 1, 3, 5, 7, 0xFFFF:
			if v.HasValidGroup() != false {
				b.Error(i, v, "excepted false, got ", v.HasValidGroup())
			}
		default:
			if v.HasValidGroup() != true {
				b.Error(i, v, "excepted true, got ", v.HasValidGroup())
			}
		}
	}
}

func BenchmarkIsGroupLength(b *testing.B) {
	for i := 0x0000; i <= 0xFFFF; i++ {
		for j := 0x0000; j <= 0xFFFF; j++ {
			var v DcmTagKey
			v.Group = uint16(i)
			v.Element = uint16(j)
			if j == 0x0000 {
				switch v.Group {
				case 1, 3, 5, 7, 0xFFFF:
					if v.IsGroupLength() != false {
						b.Error(i, v, "excepted false, got ", v.IsGroupLength())
					}
				default:
					if v.IsGroupLength() != true {
						b.Error(i, v, "excepted true, got ", v.IsGroupLength())
					}
				}
			} else {
				switch v.Group {
				case 1, 3, 5, 7, 0xFFFF:
					if v.IsGroupLength() != false {
						b.Error(i, v, "excepted false, got ", v.IsGroupLength())
					}
				default:
					if v.IsGroupLength() != false {
						b.Error(i, v, "excepted false, got ", v.IsGroupLength())
					}
				}

			}

		}
	}
}

func BenchmarkIsPrivate(b *testing.B) {
	for i := 0x0000; i <= 0xFFFF; i++ {
		var v DcmTagKey
		v.Group = uint16(i)
		switch v.Group {
		case 1, 3, 5, 7, 0xFFFF:
			if v.IsPrivate() != false {
				b.Error(i, v, "excepted false, got ", v.IsPrivate())
			}
		default:
			if (v.Group & 1) != 0 {
				if v.IsPrivate() != true {
					b.Error(i, v, "excepted true, got ", v.IsPrivate())
				}

			} else {
				if v.IsPrivate() != false {
					b.Error(i, v, "excepted false, got ", v.IsPrivate())
				}

			}
		}
	}
}

func BenchmarkIsPrivateReservation(b *testing.B) {
	for i := 0x0000; i <= 0xFFFF; i++ {
		for j := 0x0000; j <= 0xFFFF; j++ {
			v := DcmTagKey{Group: uint16(i), Element: uint16(j)}
			if v.Element >= 0x0010 && v.Element <= 0x00FF {
				switch v.Group {
				case 1, 3, 5, 7, 0xFFFF:
					if v.IsPrivateReservation() != false {
						b.Error(i, v, "excepted false, got ", v.IsPrivateReservation())
					}
				default:
					if (v.Group & 1) != 0 {
						if v.IsPrivateReservation() != true {
							b.Error(i, v, "excepted true, got ", v.IsPrivateReservation())
						}

					} else {
						if v.IsPrivateReservation() != false {
							b.Error(i, v, "excepted false, got ", v.IsPrivateReservation())
						}

					}
				}
			} else {
				switch v.Group {
				case 1, 3, 5, 7, 0xFFFF:
					if v.IsPrivateReservation() != false {
						b.Error(i, v, "excepted false, got ", v.IsPrivateReservation())
					}
				default:
					if (v.Group & 1) != 0 {
						if v.IsPrivateReservation() != false {
							b.Error(i, v, "excepted false, got ", v.IsPrivateReservation())
						}

					} else {
						if v.IsPrivateReservation() != false {
							b.Error(i, v, "excepted false, got ", v.IsPrivateReservation())
						}

					}
				}

			}

		}
	}

}

func TestHash(t *testing.T) {
	v := DcmTagKey{Group: 0x0002, Element: 0x0002}
	if v.Hash() != 131074 {
		t.Error("excepted 131074, got ", v.Hash())
	}
}

func TestToString(t *testing.T) {
	v := NewDcmTagKey()
	if v.ToString() != "(????,????)" {
		t.Error("excepted (????,????), got ", v.ToString())
	}
	v.Group = 0x001F
	v.Element = 0x002F
	if strings.ToUpper(v.ToString()) != "(001F,002F)" {
		t.Error("excepted (001F,002F), got ", v.ToString())
	}
}

func BenchmarkIsSignableTag(b *testing.B) {
	for i := 0x0000; i <= 0xFFFF; i++ {
		for j := 0x0000; j <= 0xFFFF; j++ {
			v := DcmTagKey{Group: uint16(i), Element: uint16(j)}

			if v.Element == 0x0000 {
				if v.IsSignableTag() != false {
					b.Error(v, "excepted false, got ", v.IsSignableTag())
				}
			} else if (v.Group == 0x0008) && (v.Element == 0x0001) {
				if v.IsSignableTag() != false {
					b.Error(v, "excepted false, got ", v.IsSignableTag())
				}
			} else if v.Group < 0x0008 {
				if v.IsSignableTag() != false {
					b.Error(v, "excepted false, got ", v.IsSignableTag())
				}
			} else if v.Group == 0xFFFA {
				if v.IsSignableTag() != false {
					b.Error(v, "excepted false, got ", v.IsSignableTag())
				}
			} else if (v.Group == 0x4FFE) && (v.Element == 0x0001) {
				if v.IsSignableTag() != false {
					b.Error(v, "excepted false, got ", v.IsSignableTag())
				}
			} else if (v.Group == 0xFFFC) && (v.Element == 0xFFFC) {
				if v.IsSignableTag() != false {
					b.Error(v, "excepted false, got ", v.IsSignableTag())
				}
			} else if (v.Group == 0xFFFE) && ((v.Element == 0xE00D) || (v.Element == 0xE0DD)) {
				if v.IsSignableTag() != false {
					b.Error(v, "excepted false, got ", v.IsSignableTag())
				}
			} else {
				if v.IsSignableTag() != true {
					b.Error(v, "excepted true, got ", v.IsSignableTag())
				}
			}
		}
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		base DcmTagKey
		in   DcmTagKey
		want bool
	}{
		{DcmTagKey{0x0010, 0x001F}, DcmTagKey{0x0010, 0x001F}, true},
		{DcmTagKey{0xFFFF, 0x001F}, DcmTagKey{0xFFFF, 0x001F}, true},
		{DcmTagKey{0xFFFF, 0x001F}, DcmTagKey{0xFFFF, 0x001F}, true},
		{DcmTagKey{0xFFFF, 0x001F}, DcmTagKey{0xFFFF, 0x001E}, false},
	}
	for _, c := range cases {
		got := c.base.Equal(c.in)
		if got != c.want {
			t.Errorf("%s Equal(%s)== %v, want %v ", c.base.ToString(), c.in.ToString(), got, c.want)
		}
	}
}

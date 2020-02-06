package parser

import (
	"testing"
)

func TestLengthFromData(t *testing.T) {
	cases := []struct{
		data     []byte
		index    int
		expected int
	}{
		{[]byte{0x00, 0x00}, 0, 0},
		{[]byte{0x00, 0xf0}, 0, 240},
		{[]byte{0x01, 0x01}, 0, 257},
		{[]byte{0x01, 0x00}, 0, 256},
		{[]byte{0x01, 0x00, 0x00}, 1, 0},
		{[]byte{0x01, 0x00, 0xff}, 0, 256},
	}

	for ix, data := range cases {
		seen := lengthFromData(data.data, data.index)
		if seen != data.expected {
			t.Errorf("Test case #%d, saw %d, expected %d", ix, seen, data.expected)
		}
	}
}

func TestGetSNIBlock(t *testing.T) {
	cases := []struct{
		data     []byte
		hostname string
		err      bool
	}{
		{
			[]byte{0x00, 0x06, 0x00, 0x00, 0x03, 0x66, 0x6f, 0x6f},
			"foo",
			false,
		},
		{
			[]byte{0x00, 0x06, 0x01, 0x00, 0x03, 0x66, 0x6f, 0x6f},
			"",
			true,
		},
		{
			[]byte{
				0x01, 0x06, 0x01, 0x00, 0x03, 0x66, 0x6f, 0x6f,
				0x00, 0x06, 0x00, 0x00, 0x03, 0x66, 0x6f, 0x6f,
			},
			"",
			true,
		},
		{
			[]byte{
				0x00, 0x06, 0x01, 0x00, 0x03, 0x66, 0x6f, 0x6f,
				0x00, 0x06, 0x00, 0x00, 0x03, 0x66, 0x6f, 0x6f,
			},
			"foo",
			false,
		},
	}

	for ix, data := range cases {
		sni, err := GetSNIBlock(data.data)

		sawErr := (err != nil)
		if sawErr != data.err {
			t.Errorf("Unexpected error status in test case #%d, error present %v, expected %v", ix, sawErr, data.err)
		} else {
			seen := string(sni)
			expected := data.hostname

			if seen != expected {
				t.Errorf("IN test case #%d, saw hostname %s, expected %s", ix, seen, expected)
			}
		}
	}
}

func cmpByteSlice(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for ix, val := range a {
		if val != b[ix] {
			return false
		}
	}

	return true
}

func TestGetSNBlock(t *testing.T) {
	cases := []struct{
		data     []byte
		err      bool
		expected []byte
	}{
		{
			[]byte{0x00, 0x07, 0x00, 0x00, 0x00, 0x03, 0x01, 0x02, 0x03},
			false,
			[]byte{0x01, 0x02, 0x03},
		},
		{
			[]byte{0x01, 0x07, 0x00, 0x00, 0x00, 0x03, 0x01, 0x02, 0x03},
			true,
			[]byte{},
		},
	}

	for ix, data := range cases {
		sn, err := GetSNBlock(data.data)

		sawErr := (err != nil)
		if sawErr != data.err {
			t.Errorf("Unexpected error status in test case #%d, error present %v, expected %v", ix, sawErr, data.err)
		} else {
			if !cmpByteSlice(sn, data.expected) {
				t.Errorf("IN test case #%d, saw SN %#v, expected %#v", ix, sn, data.expected)
			}
		}
	}
	
}

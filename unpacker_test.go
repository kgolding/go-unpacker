package unpacker

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Example(t *testing.T) {
	data := []byte{0xab, 0b11001011}
	definition := []string{
		"AAAABBBB", "CCCC10?ŋ",
	}
	m, err := Unpack(data, definition)
	if err != nil {
		panic(err)
	}
	for k, v := range m {
		fmt.Printf("'%s' = %X\n", string(k), v)
	}
	/* Outputs:
	'A' = A
	'B' = B
	'C' = C
	'ŋ' = 1
	*/
}

func Test_Unpacker(t *testing.T) {
	tests := []struct {
		Data []byte
		Def  []string
		Map  map[rune]uint
		Err  error
	}{
		{[]byte{0x00}, []string{"0000000"}, map[rune]uint{}, ErrBadDefinition},
		{[]byte{}, []string{"00000000"}, map[rune]uint{}, ErrBadData},
		{[]byte{0x00}, []string{"00000000", "00000000"}, map[rune]uint{}, ErrBadData},
		{[]byte{0x00}, []string{"00000000"}, map[rune]uint{}, nil},
		{[]byte{0xff}, []string{"11111111"}, map[rune]uint{}, nil},
		{[]byte{0x31}, []string{"hhhhllll"}, map[rune]uint{'h': 3, 'l': 1}, nil},
		{[]byte{0x12, 0x34}, []string{"aaaabbbb", "ccccdddd"}, map[rune]uint{'a': 1, 'b': 2, 'c': 3, 'd': 4}, nil},
		{[]byte{0b00110111}, []string{"aaabbbcc"}, map[rune]uint{'a': 1, 'b': 5, 'c': 3}, nil},
		{[]byte{0b00110111}, []string{"00110110"}, map[rune]uint{}, ErrFitBitMismatch},
		{[]byte{0x80, 0x80, 0x01, 0x84, 0x83, 0xD8,
			0xa5, 0x88, 0x20, 0x22,
			0xA2, 0x09, 0x53, 0x90,
			0xa6, 0x00, 0x10},
			[]string{"1ttt?Evv", "iiiiiiii", "iiiiiiii", "1HHHHHHH", "1mmmmmmm", "1sssssss",
				"1?WSLLLL", "LLLL1lll", "llllłłłł", "łłłłłłłł",
				"101GGGGG", "GGGG1ggg", "ggggŋŋŋŋ", "ŋŋŋŋŋŋŋŋ",
				"10??????", "KKKKKKKK", "KKKKkkkk"},
			map[rune]uint{'t': 0, 'E': 0, 'v': 0, 'H': 4, 'm': 3, 's': ToBCD(58),
				'W': 1, 'S': 0, 'L': ToBCD(58), 'l': ToBCD(2), 'ł': ToBCD(22),
				'G': ToBCD(20),
				'K': ToBCD(1), 'k': 0}, nil},
	}

	for i, test := range tests {
		// t.Logf("%d: %s", i, test.Def)
		m, err := Unpack(test.Data, test.Def)
		if test.Err == nil && err != nil {
			t.Errorf("%d: unexpected error: %v", i, err)
		} else if test.Err != nil && err == nil {
			t.Errorf("%d: expected error '%s' got nil", i, test.Err)
		} else if !errors.Is(err, test.Err) {
			t.Errorf("%d: expected error '%s' got '%v'", i, test.Err, err)
		}
		for k, v := range test.Map {
			x, ok := m[k]
			if !ok {
				t.Errorf("%d: missing '%s' from result map", i, string(k))
			} else {
				if x != v {
					t.Errorf("%d: for '%s' expected %d got %d", i, string(k), v, x)
				}
			}
		}
		// t.Logf("%d: %v", i, m)
	}
}

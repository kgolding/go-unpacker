package unpacker

import (
	"testing"
)

func Test_BCD(t *testing.T) {
	tests := []struct {
		Data uint
		Val  uint
	}{
		{uint(0x00), 0},
		{uint(0x01), 1},
		{uint(0x10), 10},
		{uint(0x99), 99},
		{uint(0x987654), 987654},
		{uint(0x1234567890), 1234567890},
		// {uint(0xaa), 0},
	}

	for i, test := range tests {
		v := BCD(test.Data)
		if v != test.Val {
			t.Errorf("BCD %d: expected %d got %d", i, test.Val, v)
		}
		d := ToBCD(test.Val)
		if d != test.Data {
			t.Errorf("ToBCD %d: expected %d got %d", i, test.Data, d)
		}
	}
}

package unpacker

import (
	"math/bits"
)

// BCD converts converts a BCD encoded uint to a normal uint
func BCD(v uint) uint {
	x := uint(0)
	col := uint(1)
	for i := bits.UintSize; i > 0; i -= 4 {
		x += col * uint(v&0xf)
		col *= 10
		v = v >> 4
	}
	return x
}

// ToBCD converts converts v to BCD encoded uint
func ToBCD(v uint) uint {
	// fmt.Println("ToBCD()", v)
	x := uint(0)
	for i := 0; i < bits.UintSize; i += 4 {
		t := v % 10 // Single digit
		t = t << i  // Shift into place
		x += t
		// fmt.Println("i:", i, "x:", x, "t:", t, "v:", v)
		v = v / 10
	}
	return x
}

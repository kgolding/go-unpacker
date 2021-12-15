package unpacker

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var ErrBadDefinition = errors.New("bad definition")
var ErrBadData = errors.New("incorrect length data")
var ErrFitBitMismatch = errors.New("fixed bit mismatch")

func Unpack(data []byte, definition []string) (map[rune]uint, error) {
	var def string
	// Validate definition
	for i, s := range definition {
		l := utf8.RuneCountInString(s)
		if l%8 != 0 {
			return nil, fmt.Errorf("length at index %d is not a multiple of 8 (%d): %w", i, l, ErrBadDefinition)
		}
		// def = append(def, []byte(s)...)
		def += s
	}

	// Validate data
	if utf8.RuneCountInString(def)/8 != len(data) {
		return nil, ErrBadData
	}

	// The result
	m := make(map[rune]uint)

	// Iterate over the full definition string of 'bits'
	idx := 0
	var b byte // current byte
	i := -1
	for _, c := range def {
		i++
		// fmt.Println(i, i%8, string(c))
		if i%8 == 0 {
			b = data[idx]
			idx++ // Move to next byte
		}
		// fmt.Printf("c:%s, b: %X, data: % X\n", string(c), b, data)
		// Get msb
		v := 0b10000000&b > 0
		// Shift all left
		b = b << 1

		switch c {
		case '1':
			if !v {
				return nil, fmt.Errorf("expected 1 in byte %d postion %d got 0: %w", idx-1, i%8, ErrFitBitMismatch)
			}

		case '0':
			if v {
				return nil, fmt.Errorf("expected 1 in byte %d postion %d got 0: %w", idx-1, i%8, ErrFitBitMismatch)
			}

		case '?':
			continue // ignore

		default:
			// If not exists, create map placeholder
			if _, ok := m[c]; !ok {
				m[c] = 0
			}
			// Shift left the current value
			m[c] = m[c] << 1
			// If the bit is set, add 1
			if v {
				m[c]++
			}

		}

	}

	return m, nil
}

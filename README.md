# go-unpacker for decoding bit values from byte array

The data is decoded according to a definition which consists of an array of strings,
one string per byte each with a rune per bit (8 runes). The rune is used to populate the results map.

Where each rune within the definition:

	* `1` Must match a bit 1
	* `0` Must match a bit 0
	* `?` Ignored bit
	* Otherwise is a returned variable being a single rune, with repeated runes having the bit's concatenated and converted to a `uint`

The returned map will have a value for each of the variables e.g. in the code below "AAAABBBB" will decode the top 4 bits of the first byte (0xab) into m['A'] and the bottom 4 bits into m['B'].

## Example usage
````golang
import "github.com/kgolding/go-unpacker"

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
````

## BCD

This package also includes `BCD()` and `ToBCD()` helpers to en/decode BCD encode `uint` to a normal `uint`.

## Notes

Runes's are used as the placeholder, allowing non-ascii chars such as æ, ł, ŋ etc.

Unpack() will return an error on:

* Length of the data does not match the length of the definition
* The definition does not have exactly 8 runes per string
* An `1` or `0` placeholder does not match the bit within the data

`uint`'s are used and both 32 & 64 bit platforms are supported.

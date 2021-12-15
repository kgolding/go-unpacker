# go-unpacker for decoding bit values from byte array

The data is decoded according to a definition which consists of an array of strings,
one string per byte each with 1 rune per bit (8 runes).

Where each rune within the definition:

	* `1` Must match a bit 1
	* `0` Must match a bit 0
	* `?` Ignored
	* Otherwise is a returned variable

The returned map will have a value for each of the variables.

````
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
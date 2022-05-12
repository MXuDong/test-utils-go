package patch

var bits = 8

// Assembles a jump to a function value
func jmpToFunctionValue(to uintptr) []byte {
	return []byte{
		0x48, 0xBA,
		byte(to >> bits * 0),
		byte(to >> bits * 1),
		byte(to >> bits * 2),
		byte(to >> bits * 3),
		byte(to >> bits * 4),
		byte(to >> bits * 5),
		byte(to >> bits * 6),
		byte(to >> bits * 7), // movabs rdx, 64 bits of value
		0xFF, 0x22,     // jmp QWORD PTR [rdx]
	}
}

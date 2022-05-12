package patch

var bits = 8

// Assembles a jump to a function value
func jmpToFunctionValue(to uintptr) []byte {
	return []byte{
		0xBA,
		byte(to >> (bits * 0)),
		byte(to >> (bits * 1)),
		byte(to >> (bits * 2)),
		byte(to >> (bits * 3)), // mov edx, to function with byte in 32 bit
		0xFF, 0x22,             // jmp DWORD PTR [edx]
	}
}

package patch

var bits = 8

// Assembles a jump to a function value
func jmpToFunctionValue(to uintptr) []byte {
	return []byte{
		0x48, 0xBA,
		byte(to >> (0 * bits)),
		byte(to >> (1 * bits)),
		byte(to >> (2 * bits)),
		byte(to >> (3 * bits)),
		byte(to >> (4 * bits)),
		byte(to >> (5 * bits)),
		byte(to >> (6 * bits)),
		byte(to >> (7 * bits)), // movabs rdx,to
		0xFF, 0x22,             // jmp QWORD PTR [rdx]
	}
}

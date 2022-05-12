// +build !windows

package patch

import (
	"syscall"
)

// pageStart return the ptr page index
func pageStart(ptr uintptr) uintptr {
	// x / n * n == x - x % n == x & (^(n)+1) == x & ^(n-1)
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}

func mprotectCrossPage(addr uintptr, length int, prot int) {
	pageSize := syscall.Getpagesize()
	// set prot to addr with length, if size > pageSize, set prot cover all page of data.
	for p := pageStart(addr); p < addr+uintptr(length); p += uintptr(pageSize) {
		page := rawMemoryAccess(p, pageSize)
		err := syscall.Mprotect(page, prot)
		if err != nil {
			panic(err)
		}
	}
}

// cpl will copy data to location
func cp(location uintptr, data []byte) {
	window := rawMemoryAccess(location, len(data))

	mprotectCrossPage(location, len(data), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(window, data[:])
}

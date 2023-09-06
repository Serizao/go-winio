//go:build windows

// Code generated by 'go generate' using "github.com/Serizao/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package process

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procK32EnumProcesses           = modkernel32.NewProc("K32EnumProcesses")
	procK32GetProcessMemoryInfo    = modkernel32.NewProc("K32GetProcessMemoryInfo")
	procQueryFullProcessImageNameW = modkernel32.NewProc("QueryFullProcessImageNameW")
)

func enumProcesses(pids *uint32, bufferSize uint32, retBufferSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procK32EnumProcesses.Addr(), 3, uintptr(unsafe.Pointer(pids)), uintptr(bufferSize), uintptr(unsafe.Pointer(retBufferSize)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getProcessMemoryInfo(process handle, memCounters *ProcessMemoryCountersEx, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procK32GetProcessMemoryInfo.Addr(), 3, uintptr(process), uintptr(unsafe.Pointer(memCounters)), uintptr(size))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func queryFullProcessImageName(process handle, flags uint32, buffer *uint16, bufferSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procQueryFullProcessImageNameW.Addr(), 4, uintptr(process), uintptr(flags), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(bufferSize)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

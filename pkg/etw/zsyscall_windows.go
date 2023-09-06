//go:build windows

// Code generated by 'go generate' using "github.com/Serizao/go-winio/tools/mkwinsyscall"; DO NOT EDIT.

package etw

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
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	procEventRegister       = modadvapi32.NewProc("EventRegister")
	procEventSetInformation = modadvapi32.NewProc("EventSetInformation")
	procEventUnregister     = modadvapi32.NewProc("EventUnregister")
	procEventWriteTransfer  = modadvapi32.NewProc("EventWriteTransfer")
)

func eventRegister(providerId *windows.GUID, callback uintptr, callbackContext uintptr, providerHandle *providerHandle) (win32err error) {
	r0, _, _ := syscall.Syscall6(procEventRegister.Addr(), 4, uintptr(unsafe.Pointer(providerId)), uintptr(callback), uintptr(callbackContext), uintptr(unsafe.Pointer(providerHandle)), 0, 0)
	if r0 != 0 {
		win32err = syscall.Errno(r0)
	}
	return
}

func eventSetInformation_64(providerHandle providerHandle, class eventInfoClass, information uintptr, length uint32) (win32err error) {
	r0, _, _ := syscall.Syscall6(procEventSetInformation.Addr(), 4, uintptr(providerHandle), uintptr(class), uintptr(information), uintptr(length), 0, 0)
	if r0 != 0 {
		win32err = syscall.Errno(r0)
	}
	return
}

func eventSetInformation_32(providerHandle_low uint32, providerHandle_high uint32, class eventInfoClass, information uintptr, length uint32) (win32err error) {
	r0, _, _ := syscall.Syscall6(procEventSetInformation.Addr(), 5, uintptr(providerHandle_low), uintptr(providerHandle_high), uintptr(class), uintptr(information), uintptr(length), 0)
	if r0 != 0 {
		win32err = syscall.Errno(r0)
	}
	return
}

func eventUnregister_64(providerHandle providerHandle) (win32err error) {
	r0, _, _ := syscall.Syscall(procEventUnregister.Addr(), 1, uintptr(providerHandle), 0, 0)
	if r0 != 0 {
		win32err = syscall.Errno(r0)
	}
	return
}

func eventUnregister_32(providerHandle_low uint32, providerHandle_high uint32) (win32err error) {
	r0, _, _ := syscall.Syscall(procEventUnregister.Addr(), 2, uintptr(providerHandle_low), uintptr(providerHandle_high), 0)
	if r0 != 0 {
		win32err = syscall.Errno(r0)
	}
	return
}

func eventWriteTransfer_64(providerHandle providerHandle, descriptor *eventDescriptor, activityID *windows.GUID, relatedActivityID *windows.GUID, dataDescriptorCount uint32, dataDescriptors *eventDataDescriptor) (win32err error) {
	r0, _, _ := syscall.Syscall6(procEventWriteTransfer.Addr(), 6, uintptr(providerHandle), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(activityID)), uintptr(unsafe.Pointer(relatedActivityID)), uintptr(dataDescriptorCount), uintptr(unsafe.Pointer(dataDescriptors)))
	if r0 != 0 {
		win32err = syscall.Errno(r0)
	}
	return
}

func eventWriteTransfer_32(providerHandle_low uint32, providerHandle_high uint32, descriptor *eventDescriptor, activityID *windows.GUID, relatedActivityID *windows.GUID, dataDescriptorCount uint32, dataDescriptors *eventDataDescriptor) (win32err error) {
	r0, _, _ := syscall.Syscall9(procEventWriteTransfer.Addr(), 7, uintptr(providerHandle_low), uintptr(providerHandle_high), uintptr(unsafe.Pointer(descriptor)), uintptr(unsafe.Pointer(activityID)), uintptr(unsafe.Pointer(relatedActivityID)), uintptr(dataDescriptorCount), uintptr(unsafe.Pointer(dataDescriptors)), 0, 0)
	if r0 != 0 {
		win32err = syscall.Errno(r0)
	}
	return
}

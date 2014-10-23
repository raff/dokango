package dokan

import (
	"syscall"
	"unsafe"
)

var (
	mod = syscall.NewLazyDLL("dokan.dll")

	fn_DokanMain             = mod.NewProc("DokanMain")
	fn_DokanUnmount          = mod.NewProc("DokanUnmount")
	fn_DokanRemoveMountPoint = mod.NewProc("DokanRemoveMountPoint")
	fn_DokanVersion          = mod.NewProc("DokanVersion")
	fn_DokanDriverVersion    = mod.NewProc("DokanDriverVersion")
	fn_DokanResetTimeout     = mod.NewProc("DokanResetTimeout")
)

func DokanMain() {
}

func DokanVersion() uint {
	ret, _, _ := fn_DokanVersion.Call()
	return uint(ret)
}

func DokanDriverVersion() uint {
	ret, _, _ := fn_DokanDriverVersion.Call()
	return uint(ret)
}

func DokanUnmount(driveLetter byte) bool {
	ret, _, _ := fn_DokanUnmount.Call(uintptr(uint(driveLetter)))
	return ret != 0
}

func DokanRemoveMountPoint(mountPoint string) bool {
	ret, _, _ := fn_DokanRemoveMountPoint.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(mountPoint))))
	return ret != 0
}

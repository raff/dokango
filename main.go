package main

import (
	"syscall"
	"unsafe"
	"fmt"
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

func DokanUnmount(driveLetter byte) int {
	ret, _, _ := fn_DokanUnmount.Call(uintptr(uint(driveLetter)))
	return int(ret)
}

func DokanRemoveMountPoint(mountPoint string) int {
	ret, _, _ := fn_DokanRemoveMountPoint.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(mountPoint))))
	return int(ret)
}

func main() {
	fmt.Println("version", DokanVersion())
	fmt.Println("drive version", DokanDriverVersion())
	fmt.Println("unmount", DokanUnmount('B'))
	fmt.Println("removeMountPoint", DokanRemoveMountPoint("/mnt/stuff"))
}

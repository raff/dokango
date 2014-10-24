package dokan

import "unsafe"
import "syscall"

var (
	mod = syscall.NewLazyDLL("dokan.dll")

	fn_DokanMain             = mod.NewProc("DokanMain")
	fn_DokanUnmount          = mod.NewProc("DokanUnmount")
	fn_DokanRemoveMountPoint = mod.NewProc("DokanRemoveMountPoint")
	fn_DokanVersion          = mod.NewProc("DokanVersion")
	fn_DokanDriverVersion    = mod.NewProc("DokanDriverVersion")
	fn_DokanResetTimeout     = mod.NewProc("DokanResetTimeout")
)

const (
	DOKAN_SUCCESS              = 0
	DOKAN_ERROR                = -1 // General Error
	DOKAN_DRIVE_LETTER_ERROR   = -2 // Bad Drive letter
	DOKAN_DRIVER_INSTALL_ERROR = -3 // Can't install driver
	DOKAN_START_ERROR          = -4 // Driver something wrong
	DOKAN_MOUNT_ERROR          = -5 // Can't assign a drive letter or mount point
	DOKAN_MOUNT_POINT_ERROR    = -6 // Mount point is invalid
)

type DokanOptions struct {
	Version      uint16
	ThreadCount  uint16
	DebugMode    bool
	UseStdErr    bool
	UseAltStream bool
	UseKeepAlive bool
	NetworkDrive bool
	VolumeLabel  string
	MountPoint   string
}

type DokanOperations struct {
}

func DokanMain(options DokanOptions, operations DokanOperations) int {
	if len(options.VolumeLabel) == 0 {
		options.VolumeLabel = "DOKAN"
	}

	//ret, _, _ := fn_DokanMain(dokanOptions, dokanOperations)
	ret := 0
	return int(ret)
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

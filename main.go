package main

import (
	"fmt"

	"github.com/raff/dokango/dokan"
)

func main() {
	fmt.Println("version", dokan.DokanVersion())
	fmt.Println("drive version", dokan.DokanDriverVersion())
	fmt.Println("unmount", dokan.DokanUnmount('B'))
	fmt.Println("removeMountPoint", dokan.DokanRemoveMountPoint("/mnt/stuff"))
}

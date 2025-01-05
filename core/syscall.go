package core;

import (
	"syscall"
	"unsafe"
)

// SysCall struct to receive info on the terminal screen
type TerminalSize struct {
    Row     uint16
    Col     uint16
    Xpixel  uint16 
    Ypixel  uint16
}

func GetTerminalSize() (*TerminalSize) {

    ts := new(TerminalSize)
        
    r1, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(0x5413),
        uintptr(unsafe.Pointer(ts)),
    )

    if int(r1) == -1 {
        // TODO: better handle
        return nil
    }
    

    return ts

}



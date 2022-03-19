package cage

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)


func PivotRoot(newroot string) error {
	// putold := filepath.Join(newroot, "/.pivot_root")
	
	// if err := syscall.Mount(newroot, newroot, "", syscall.MS_BIND | syscall.MS_REC, ""); err != nil {
	// 	return fmt.Errorf("failed to bind mount newroot: %w", err) 
	// }

	// if err := os.Mkdir(putold, 0700); err != nil {
	// 	return fmt.Errorf("failed to mkdir putold: %w", err) 
	// }

	// if err := syscall.PivotRoot(newroot, putold); err != nil {
	// 	return fmt.Errorf("failed to pivot: %w", err) 
	// }

	if err := syscall.Chroot(newroot); err != nil {
		return fmt.Errorf("failed to chroot: %w", err) 
	}

	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("failed to chdir: %w", err) 
	}

	// if err := syscall.Unmount("/.pivot_root", syscall.MNT_DETACH); err != nil {
	// 	return fmt.Errorf("failed to unmount: %w", err) 
	// }

	if err := os.RemoveAll("/.pivot_root"); err != nil {
		return fmt.Errorf("failed to rm putold: %w", err) 
	}

	return nil
}


func MountProc(newroot string) error {
    source := "proc"
    target := filepath.Join(newroot, "/proc")
    fstype := "proc"
    flags := 0
    data := ""

    // 创建Proc目录
    if err := os.MkdirAll(target, 0755); err != nil {
		return fmt.Errorf("failed to create proc: %w", err)
	}
    // 挂载proc
    // mount -t proc proc ./proc
    if err := syscall.Mount(source, target, fstype, uintptr(flags), data); err != nil {
		return fmt.Errorf("failed to mount proc: %w", err)
    }

    return nil
}

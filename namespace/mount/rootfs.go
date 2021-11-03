// Copyright 2021 liang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)
/* https://github.com/teddyking/ns-process/blob/master/rootfs.go */
func pivotRoot(newroot string) error {
	putold := filepath.Join(newroot, "/.pivot_root")

	// bind mount newroot to itself - this is a slight hack needed to satisfy the
	// pivot_root requirement that newroot and putold must not be on the same
	// filesystem as the current root
	if err := syscall.Mount(newroot, newroot, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return err
	}

	// create putold directory
	if err := os.MkdirAll(putold, 0700); err != nil {
		return err
	}

	// call pivot_root
	if err := syscall.PivotRoot(newroot, putold); err != nil {
		return err
	}

	// ensure current working directory is set to new root
	if err := os.Chdir("/"); err != nil {
		return err
	}

	// umount putold, which now lives at /.pivot_root
	putold = "/.pivot_root"
	if err := syscall.Unmount(putold, syscall.MNT_DETACH); err != nil {
		return err
	}

	// remove putold
	if err := os.RemoveAll(putold); err != nil {
		return err
	}

	return nil
}

func mountProc(newroot string) error {
	source := "proc"
	target := filepath.Join(newroot, "/proc")
	fstype := "proc"
	flags := 0
	data := ""

	os.MkdirAll(target, 0755)
	if err := syscall.Mount(source, target, fstype, uintptr(flags), data); err != nil {
		return err
	}

	return nil
}

func exitIfRootfsNotFound(rootfsPath string) {
	if _, err := os.Stat(rootfsPath); os.IsNotExist(err) {
		usefulErrorMsg := fmt.Sprintf(`
"%s" does not exist.
Please create this directory and unpack a suitable root filesystem inside it.
An example rootfs, BusyBox, can be downloaded and unpacked as follows:
wget "https://raw.githubusercontent.com/teddyking/ns-process/4.0/assets/busybox.tar"
mkdir -p %s
tar -C %s -xf busybox.tar
`, rootfsPath, rootfsPath, rootfsPath)

		fmt.Println(usefulErrorMsg)
		os.Exit(1)
	}
}

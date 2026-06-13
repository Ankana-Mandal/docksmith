package main

import (
    "os"
    "strconv"
)

func setupCgroup(pid int) {
    os.MkdirAll("/sys/fs/cgroup/docksmith", 0755)
    os.WriteFile("/sys/fs/cgroup/docksmith/memory.max", []byte("536870912"), 0644)
    os.WriteFile("/sys/fs/cgroup/docksmith/cgroup.procs", []byte(strconv.Itoa(pid)), 0644)
}

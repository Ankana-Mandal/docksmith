package main

import (
    "os"
    "strconv"
)

func setupCgroup(pid int,id string) {
    path := "/sys/fs/cgroup/docksmith-" + id
    os.MkdirAll(path, 0755)
    os.WriteFile(path+"/memory.max", []byte("536870912"), 0644)
    os.WriteFile(path+"/cgroup.procs", []byte(strconv.Itoa(pid)), 0644)
}
func cleanupCgroup(id string) {
    path := "/sys/fs/cgroup/docksmith-" + id
    os.Remove(path)
}

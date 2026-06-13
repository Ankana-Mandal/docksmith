package main
import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
    "time"
)

func main() {
    switch os.Args[1] {
    case "current":
        current()
    case "child":
        child()
    case "ps":
        ps()
    case "kill":
        killcontainer(os.Args[2])
    default:
        panic("bad command")
    }
}
func current() {
    fmt.Printf("current PID: %d\n", os.Getpid())
    base, _ := os.Getwd()
    id := fmt.Sprintf("c%d", time.Now().UnixNano())
    merged := base + "/containers/" + id + "/merged"
    upper := base + "/containers/" + id + "/upper"
    work := base + "/containers/" + id + "/work"
    rootfs := base + "/rootfs"

    cmd := exec.Command("/proc/self/exe", "child",merged,os.Args[2])
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
    }

   os.MkdirAll(merged, 0755)
   os.MkdirAll(upper, 0755)
   os.MkdirAll(work, 0755)

    err := syscall.Mount("overlay",merged,"overlay",0,"lowerdir=" + rootfs + ",upperdir=" + upper + ",workdir=" + work,)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }  
    cmd.Start()

    setupCgroup(cmd.Process.Pid,id)

    c := Container{
    ID: id,
    PID: cmd.Process.Pid,
    Status: "running",
    }
    updateJSON(c)
    cmd.Wait() 
    cleanupCgroup(id)

    c.Status = "exited"
    updateJSON(c)
    syscall.Unmount(merged, 0)
}


func child() {
   
    merged := os.Args[2]
    fmt.Printf("current PID: %d\n", os.Getpid())

    syscall.Sethostname([]byte("hi_child1"))
    syscall.Chroot(merged)
    syscall.Chdir("/")

    if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
        fmt.Printf("mount error: %v\n", err)
    }
    cmd := exec.Command(os.Args[3], os.Args[4:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Printf("run error: %v\n", err)
    }
    syscall.Unmount(
    "/proc",
    0)
}


package main

import (
    "encoding/json"
    "fmt"
    "os"
    "syscall"
)

type Container struct {
    ID     string
    PID    int
    Status string
}

func updateJSON(c Container) {
    os.MkdirAll("containers", 0755)
    data, err := json.Marshal(c)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    filename := "containers/" + c.ID + ".json"
    err = os.WriteFile(filename, data, 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func ps() {
    files, err := os.ReadDir("containers")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    for _, file := range files {
        data, err := os.ReadFile("containers/" + file.Name())
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            os.Exit(1)
        }
        var c Container
        json.Unmarshal(data, &c)
        fmt.Printf("ID: %s  PID: %d  Status: %s\n", c.ID, c.PID, c.Status)
    }
}

func killcontainer(id string) {
    data, err := os.ReadFile("containers/" + id + ".json")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    var c Container
    json.Unmarshal(data, &c)
    syscall.Kill(c.PID, syscall.SIGKILL)
    c.Status = "exited"
    updateJSON(c)
}


# Docksmith

Docksmith is a lightweight container runtime built from scratch in Go to understand how containers work internally. It uses Linux namespaces, OverlayFS, chroot, and cgroups to provide process and filesystem isolation.

## Architecture

current() creates a child process using Linux namespaces.

The child process enters a new PID and UTS namespace, mounts OverlayFS for an isolated filesystem

applies cgroup limits, performs `chroot()`

chroot is used for filesystem isolation. Unlike pivot_root, chroot is escapable by a root process — a known limitation and the next improvement planned for Docksmith.




## Features



* PID and UTS namespaces

* OverlayFS-based filesystem isolation

* `chroot()` container root filesystem

* Memory limits using cgroups v2

* Container metadata storage in JSON

* `ps` and `kill` commands

## Usage

Build:



```

go build -o docksmith .

```



Run a container:



```

sudo ./docksmith current /bin/bash

```



List containers:



```

./docksmith ps

```



Kill a container:



```

sudo ./docksmith kill <container-id>

```



## What I learned

Building Docksmith taught me how containers are simply Linux kernel features combined together. I gained hands-on experience with namespaces, OverlayFS, cgroups, and process isolation, and learned how container runtimes like Docker work under the hood.



## What's next



* Network namespaces

* Docksmithfile support

* Image management

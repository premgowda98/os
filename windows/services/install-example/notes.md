# Running a Go Program as a Windows Service: Two Approaches

This note explains how to install and manage a Go program as a Windows service using two methods:

---

## 1. Using Go's Native Service Manager (`golang.org/x/sys/windows/svc/mgr`)

### Overview
- Go provides packages (`windows/svc` and `windows/svc/mgr`) to create, install, and manage Windows services natively.
- You can add logic to your Go binary to handle `--install` and `--uninstall` flags, allowing the binary to self-register or remove itself as a service.

### How It Works
- The `mgr` package lets you connect to the Windows Service Control Manager (SCM) and create/delete services programmatically.
- The service implementation uses the `svc` package to handle service lifecycle events (start, stop, interrogate, etc).

### Example Usage
```sh
# Install the service
binary.exe --install

# Uninstall the service
binary.exe --uninstall
```

### Example Code Snippet
```go
import "golang.org/x/sys/windows/svc/mgr"

func installService(name, exepath string) error {
    m, err := mgr.Connect()
    if err != nil { return err }
    defer m.Disconnect()
    s, err := m.OpenService(name)
    if err == nil { s.Close(); return fmt.Errorf("service %s already exists", name) }
    s, err = m.CreateService(name, exepath, mgr.Config{DisplayName: name})
    if err != nil { return err }
    defer s.Close()
    return nil
}

func uninstallService(name string) error {
    m, err := mgr.Connect()
    if err != nil { return err }
    defer m.Disconnect()
    s, err := m.OpenService(name)
    if err != nil { return err }
    defer s.Close()
    return s.Delete()
}
```

### Pros
- No external dependencies required.
- Full control from within your Go code.
- Can be distributed as a single binary.

### Cons
- Slightly more complex code.
- Must run as Administrator to install/uninstall services.

---

## 2. Using NSSM (Non-Sucking Service Manager)

### Overview
- NSSM is a third-party tool that can wrap any executable (including Go binaries) and run it as a Windows service.
- You do not need to modify your Go code to use NSSM.

### How It Works
- Download and install NSSM from https://nssm.cc/
- Use NSSM commands to install or remove your Go binary as a service.

### Example Usage
```sh
# Install the service
nssm install MyGoService C:\path\to\binary.exe

# Remove the service
nssm remove MyGoService confirm
```

### Pros
- Very simple to use.
- No code changes required.
- Can wrap any executable, not just Go programs.

### Cons
- Requires distributing and installing NSSM.
- Service management is external to your binary.

---

## Summary Table
| Feature                | Go Service Manager         | NSSM                   |
|------------------------|---------------------------|------------------------|
| Self-contained binary  | Yes                       | No                     |
| Needs external tool    | No                        | Yes (nssm.exe)         |
| Code changes needed    | Yes                       | No                     |
| Fine-grained control   | Yes                       | No                     |
| Simplicity             | Moderate                  | Very easy              |

---

## When to Use Which?
- **Use Go's service manager** if you want a single binary, full control, and are comfortable with a little extra code.
- **Use NSSM** if you want the simplest possible setup and don't mind an extra dependency.

---

## References
- [Go Windows Service Example](https://pkg.go.dev/golang.org/x/sys/windows/svc)
- [NSSM Official Site](https://nssm.cc/)
- [Microsoft: Service Control Manager](https://docs.microsoft.com/en-us/windows/win32/services/service-control-manager)

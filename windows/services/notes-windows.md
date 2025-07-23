## Create a service in Windows

Windows manages services through the **Service Control Manager (SCM)**, while user-level startup tasks are typically handled by the **Registry** or the **Startup folder**.

---

## 📁 1. User-Level Startup Programs

This is the equivalent of user-level `LaunchAgents` on macOS or `systemd` user services on Linux. The program runs only when a specific user logs in.

### ✅ Purpose:
* Runs **only for the logged-in user**
* Starts **when that user logs in**
* Runs with **user permissions**
* Can interact with the user's desktop and session

### 📌 Methods:

1.  **Registry:** `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run`
2.  **Startup Folder:** `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup`

### 👤 Example use case:
* A background app like Dropbox, Slack, or a personal sync script.

---

## 📁 2. System-Level Services

These are true background services managed by the SCM. They are the equivalent of `LaunchDaemons` on macOS or system-level `systemd` services on Linux.

### ✅ Purpose:
* Runs for the **entire system**, independent of user logins
* Can start **at boot time**
* Typically runs with high privileges (e.g., `LocalSystem`, `NetworkService`)
* Cannot interact directly with a user's desktop by default

### 📌 Path:
* Managed by the **Service Control Manager (SCM)**, not a simple directory.

### 🧠 Example use case:
* System daemons, web servers, databases, hardware monitoring.

---

## 🔄 Summary Table

| Method                       | Runs As       | When It Runs           | Scope       | Interaction |
| ---------------------------- | ------------- | ---------------------- | ----------- | ----------- |
| **Registry (HKCU)**          | User          | On user login          | Per-user    | ✅ Yes      |
| **Startup Folder**           | User          | On user login          | Per-user    | ✅ Yes      |
| **Windows Service**          | `LocalSystem` | At boot or on-demand   | System-wide | ❌ No       |

---

## 🛠️ Setup Steps

### Step 1: Build and Install Binary

1.  **Build the Go binary for Windows:**
    ```bash
    # From your Linux/macOS environment
    GOOS=windows GOARCH=amd64 go build -o simplego.exe .
    ```

2.  **For User Startup:** Copy to a user-accessible location.
    *   Create a folder: `C:\Users\YourUser\AppData\Local\simplego`
    *   Copy `simplego.exe` into it.

3.  **For System Service:** Copy to a system-wide, secure location.
    *   Create a folder: `C:\Program Files\simplego`
    *   Copy `simplego.exe` into it.

---

## 📄 Create Startup Tasks & Services

### User-Level Startup (Registry Method)

This is the most common and reliable method for user-level tasks.

**Create the startup task (run in Command Prompt):**
```cmd
:: Add a new registry key to run simplego.exe on login
reg add "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v "SimpleGo" /t REG_SZ /d "C:\Users\YourUser\AppData\Local\simplego\simplego.exe" /f
```

**Logging:** The `simplego.exe` program will create `simplego.log` inside `C:\Users\YourUser\AppData\Local\simplego\`.

**To Remove:**
```cmd
reg delete "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v "SimpleGo" /f
```

---

### System-Level Service

This requires an **Administrator Command Prompt**.

**Create the service:**
```cmd
:: Create the service with sc.exe
sc.exe create SimpleGo binPath="C:\Program Files\simplego\simplego.exe"

:: Optional: Set a description
sc.exe description SimpleGo "A simple Go background service."

:: Optional: Configure the service to start automatically at boot
sc.exe config SimpleGo start=auto
```

**Logging:** The Go program will attempt to write `simplego.log` in `C:\Program Files\simplego\`. This will fail due to permissions. For a real service, you should modify the Go code to write logs to a location like `C:\ProgramData\simplego`.

---

## 🟢 Useful Windows Service Commands (`sc.exe`)

Run these in an **Administrator Command Prompt**.

| Command                          | Description                               |
| -------------------------------- | ----------------------------------------- |
| `sc.exe start SimpleGo`          | Start the service                         |
| `sc.exe stop SimpleGo`           | Stop the service                          |
| `sc.exe query SimpleGo`          | Show service status                       |
| `sc.exe qc SimpleGo`             | Show detailed service configuration       |
| `sc.exe config SimpleGo start=auto` | Set service to start automatically at boot |
| `sc.exe config SimpleGo start=demand`| Set service to start manually (on-demand) |
| `sc.exe delete SimpleGo`         | Delete the service                        |

---

## 🔍 Debugging and Monitoring

### View Logs
*   **File-based logs:** Check the path where your Go program writes its log file (e.g., `C:\Users\YourUser\AppData\Local\simplego\simplego.log`).
*   **Windows Event Viewer:** For system services, errors related to starting or stopping the service itself are logged here.
    1.  Press `Win + R` and type `eventvwr.msc`.
    2.  Navigate to `Windows Logs` -> `Application` or `System`.
    3.  Look for events with the source `Service Control Manager` or your service name.

---

## 🧹 Remove Service / Startup Task

### User Startup Task
```cmd
:: Remove the registry key
reg delete "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v "SimpleGo" /f

:: Remove the binary and logs (optional)
rmdir /s /q "C:\Users\YourUser\AppData\Local\simplego"
```

### System Service
```cmd
:: Stop the service if it's running
sc.exe stop SimpleGo

:: Delete the service
sc.exe delete SimpleGo

:: Remove the binary and logs (optional)
rmdir /s /q "C:\Program Files\simplego"
```

---

## 🔄 Comparison with macOS and Linux

| Windows                      | macOS (`launchd`)                 | Linux (`systemd`)               |
| ---------------------------- | --------------------------------- | ------------------------------- |
| **Registry/Startup Folder**  | `~/Library/LaunchAgents/`         | `~/.config/systemd/user/`       |
| **Windows Service (SCM)**    | `/Library/LaunchDaemons/`         | `/etc/systemd/system/`          |
| `sc.exe` / `powershell`      | `launchctl`                       | `systemctl`                     |
| **Event Viewer**             | `log show`                        | `journalctl`                    |
| **.exe binary**              | Mach-O binary                     | ELF binary                      |

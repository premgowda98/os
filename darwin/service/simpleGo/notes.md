## Create a service in macOS

macOS uses **launchd** to manage services. There are **two levels** of services: **user-level** and **system-level**.

---

## 📁 1. User-Level Services — `~/Library/LaunchAgents/`

### ✅ Purpose:
* Runs **only for the logged-in user**
* Starts **when that user logs in**
* Runs with **user permissions** (not `root`)
* Cannot access system-level services or hardware directly

### 📌 Path:
```
~/Library/LaunchAgents/
```

### 👤 Example use case:
* Personal background apps like Dropbox, Slack, user-specific monitoring tools

---

## 📁 2. System-Level Services — `/Library/LaunchDaemons/`

### ✅ Purpose:
* Runs for the **entire system**
* Starts **at boot time**
* Runs with **root privileges**
* Can access system resources and hardware

### 📌 Path:
```
/Library/LaunchDaemons/
```

### 🧠 Example use case:
* System daemons, web servers, databases, system monitoring

---

## 🔄 Summary Table

| Location                     | Runs As | When It Runs              | Scope                    |
| ---------------------------- | ------- | ------------------------- | ------------------------ |
| `~/Library/LaunchAgents/`    | User    | When that user logs in    | Per-user                 |
| `/Library/LaunchAgents/`     | User    | When **any** user logs in | All users (individually) |
| `/Library/LaunchDaemons/`    | `root`  | At boot                   | System-wide, background  |

[About plist files](plist.md) | [User vs System Level Details](user-system.md)

---

## 🛠️ Setup Steps

### Step 1: Build and Install Binary

1. **Build the Go binary:**
   ```bash
   go build -o simplego .
   ```

2. **For User Service:** Create user bin directory
   ```bash
   mkdir -p /usr/local/bin/simplego
   cp simplego /usr/local/bin/simplego/
   chmod +x /usr/local/bin/simplego/simplego
   chown -R $USER /usr/local/bin/simplego
   ```

3. **For System Service:** Copy to system location with root ownership
   ```bash
   sudo mkdir -p /usr/local/bin/simplego
   sudo cp simplego /usr/local/bin/simplego/
   sudo chmod +x /usr/local/bin/simplego/simplego
   sudo chown -R root:wheel /usr/local/bin/simplego
   ```

---

## 📄 Create Service Files (.plist)

### User-Level Service

**File:** `~/Library/LaunchAgents/com.user.simplego.plist`

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.user.simplego</string>

    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/simplego/simplego</string>
    </array>

    <key>RunAtLoad</key>
    <true/>

    <key>StandardOutPath</key>
    <string>/usr/local/bin/simplego/stdout.log</string>

    <key>StandardErrorPath</key>
    <string>/usr/local/bin/simplego/stderr.log</string>
</dict>
</plist>
```

**Create the service:**
```bash
# Create directories
mkdir -p ~/Library/LaunchAgents
mkdir -p /usr/local/bin/simplego

# Create plist file
nano ~/Library/LaunchAgents/com.user.simplego.plist

# Load service (registers with launchd and starts if RunAtLoad=true)
launchctl load ~/Library/LaunchAgents/com.user.simplego.plist

# Start service immediately (if not auto-started)
launchctl start com.user.simplego
```

---

### System-Level Service

**File:** `/Library/LaunchDaemons/com.user.simplego.daemon.plist`

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.user.simplego.daemon</string>

    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/simplego/simplego</string>
    </array>

    <key>RunAtLoad</key>
    <true/>

    <key>StandardOutPath</key>
    <string>/usr/local/bin/simplego/stdout.log</string>

    <key>StandardErrorPath</key>
    <string>/usr/local/bin/simplego/stderr.log</string>
</dict>
</plist>
```

**Create the service:**
```bash
# Create log directory (with proper permissions)
sudo mkdir -p /usr/local/bin/simplego

# Create plist file
sudo nano /Library/LaunchDaemons/com.user.simplego.daemon.plist

# Change ownership to root
sudo chown root:wheel /Library/LaunchDaemons/com.user.simplego.daemon.plist

# Load service (registers with launchd and starts if RunAtLoad=true)
sudo launchctl load /Library/LaunchDaemons/com.user.simplego.daemon.plist

# Start service immediately (if not auto-started)
sudo launchctl start com.user.simplego.daemon
```

---

## 🟢 Useful launchctl Commands

### User Services

| Command                                                    | Description                        |
| ---------------------------------------------------------- | ---------------------------------- |
| `launchctl load ~/Library/LaunchAgents/service.plist`     | Load and register service          |
| `launchctl start com.user.service`                        | Start service                      |
| `launchctl stop com.user.service`                         | Stop service                       |
| `launchctl unload ~/Library/LaunchAgents/service.plist`   | Unload and stop service            |
| `launchctl list`                                           | List all loaded services           |
| `launchctl list | grep service`                            | Find specific service              |

### System Services

| Command                                                      | Description                      |
| ------------------------------------------------------------ | -------------------------------- |
| `sudo launchctl load /Library/LaunchDaemons/service.plist`  | Load and register service        |
| `sudo launchctl start com.user.service.daemon`              | Start service                    |
| `sudo launchctl stop com.user.service.daemon`               | Stop service                     |
| `sudo launchctl unload /Library/LaunchDaemons/service.plist`| Unload and stop service          |
| `sudo launchctl list`                                        | List all loaded system services  |

---

## 🧰 .plist File Explained

### Key Plist Keys

| Key                    | Description                                  |
| ---------------------- | -------------------------------------------- |
| `Label`                | Unique name of the job (used in `launchctl`) |
| `ProgramArguments`     | Array of command and args to run             |
| `RunAtLoad`            | Run immediately on load/login                |
| `EnvironmentVariables` | Set env vars like `HOME`, `PATH`             |
| `StandardOutPath`      | File path for stdout                         |
| `StandardErrorPath`    | File path for stderr                         |
| `WorkingDirectory`     | Directory to run the program from            |
| `KeepAlive`            | Restart if program exits                     |

---

## 🔍 Debugging and Monitoring

### View Logs
```bash
# General launchd logs
log show --predicate 'process == "launchd"' --last 5m

# Service-specific logs (check your StandardOutPath/StandardErrorPath)
tail -f /usr/local/bin/simplego/stdout.log
tail -f /usr/local/bin/simplego/stderr.log

# System logs for launchd
log show --predicate 'subsystem == "com.apple.launchd"' --last 1h
```

### Check Service Status
```bash
# List all services and find yours
launchctl list | grep simplego

# Detailed service information
launchctl list com.user.simplego

# Check if service is loaded
launchctl print user/$(id -u)/com.user.simplego
```

---

## 🧹 Remove Service

### User Service
```bash
# Stop and unload
launchctl unload ~/Library/LaunchAgents/com.user.simplego.plist

# Remove plist file
rm ~/Library/LaunchAgents/com.user.simplego.plist

# Remove binary and logs (optional)
rm -rf /usr/local/bin/simplego
```

### System Service
```bash
# Stop and unload
sudo launchctl unload /Library/LaunchDaemons/com.user.simplego.daemon.plist

# Remove plist file
sudo rm /Library/LaunchDaemons/com.user.simplego.daemon.plist

# Remove binary and logs (optional)
sudo rm -rf /usr/local/bin/simplego
```

---

## 🚦 Quick Reference: User vs System

| Goal                                    | Use                            |
| --------------------------------------- | ------------------------------ |
| Personal background task for user      | `~/Library/LaunchAgents/`      |
| System-wide service at boot            | `/Library/LaunchDaemons/`      |
| Service for all users after login      | `/Library/LaunchAgents/`       |
| Service that needs root privileges     | `/Library/LaunchDaemons/`      |

---

## 📝 Notes

* **Agents** run in the user's login session and can interact with the GUI
* **Daemons** run in the background as `root` and **do not** have access to GUI components
* Always use **full paths** in plist files
* Use **reverse DNS notation** for Label (e.g., `com.company.service`)
* Test changes by unloading → editing → loading again
* `.plist` files must be in the correct folder and have `.plist` extension

---

## 🔄 Comparison with Linux

| macOS Command                                              | Linux Equivalent                               | Description                           |
| ---------------------------------------------------------- | ---------------------------------------------- | ------------------------------------- |
| `launchctl load ~/Library/LaunchAgents/service.plist`     | `systemctl --user enable service.service`     | Register service (load from file)    |
| `launchctl start com.user.service`                        | `systemctl --user start service.service`      | Start service immediately             |
| `launchctl stop com.user.service`                         | `systemctl --user stop service.service`       | Stop service                          |
| `launchctl unload ~/Library/LaunchAgents/service.plist`   | `systemctl --user disable service.service`    | Unregister service                    |
| `launchctl list`                                           | `systemctl --user list-units --type=service`  | List all services                     |
| `log show --predicate 'process == "service"'`             | `journalctl --user -u service.service`        | View service logs                     |

### Key Differences:

| Aspect           | macOS (`launchctl`)                     | Linux (`systemctl`)                    |
| ---------------- | --------------------------------------- | -------------------------------------- |
| **File Reference** | Uses **file path**: `~/Library/LaunchAgents/service.plist` | Uses **service name**: `service.service` |
| **Load Command**   | `launchctl load <filepath>`            | `systemctl enable <servicename>`       |
| **File Location**  | You specify the exact file path        | systemd looks in predefined directories |
| **Service Discovery** | Reads plist from specified path     | Scans standard directories automatically |
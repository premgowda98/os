## Create a service in Ubuntu/Linux

Ubuntu uses **systemd** to manage services, similar to how macOS uses `launchd`. There are **two levels** of services: **user-level** and **system-level**.

---

## 📁 1. User-Level Services — `~/.config/systemd/user/`

### ✅ Purpose:
* Runs **only for the logged-in user**
* Starts **when that user logs in** (or when user session is active)
* Runs with **user permissions** (not `root`)
* Cannot access system-level services or hardware directly

### 📌 Path:
```
~/.config/systemd/user/
```

### 👤 Example use case:
* Personal background services, user-specific monitoring tools

---

## 📁 2. System-Level Services — `/etc/systemd/system/`

### ✅ Purpose:
* Runs for the **entire system**
* Starts **at boot time**
* Runs with **root privileges** or specified user
* Can access system resources and hardware

### 📌 Path:
```
/etc/systemd/system/
```

### 🧠 Example use case:
* System daemons, web servers, databases, system monitoring

---

## 🔄 Summary Table

| Location                     | Runs As    | When It Runs           | Scope       |
| ---------------------------- | ---------- | ---------------------- | ----------- |
| `~/.config/systemd/user/`    | User       | On user login/session  | Per-user    |
| `/etc/systemd/system/`       | root/user  | At boot                | System-wide |
| `/lib/systemd/system/`       | root/user  | At boot (package mgmt) | System-wide |

---

## 🛠️ Setup Steps

### Step 1: Build and Install Binary

1. **Build the Go binary:**
   ```bash
   go build -o simplego .
   ```

2. **For User Service:** Copy to user bin directory
   ```bash
   mkdir -p ~/bin
   cp simplego ~/bin/
   chmod +x ~/bin/simplego
   ```

3. **For System Service:** Copy to system bin directory
   ```bash
   sudo cp simplego /usr/local/bin/simplego
   sudo chmod -R +x /usr/local/bin/simplego
   sudo chown -R root:root /usr/local/bin/simplego
   ```

---

## 📄 Create Service Files

### User-Level Service

**File:** `~/.config/systemd/user/simplego.service`

```ini
[Unit]
Description=Simple Go User Service
After=graphical-session.target

[Service]
Type=oneshot
ExecStart=%h/bin/simplego
StandardOutput=file:%h/.local/share/simplego/stdout.log
StandardError=file:%h/.local/share/simplego/stderr.log
RemainAfterExit=no

[Install]
WantedBy=default.target
```

**Create the service:**
```bash
# Create directories
mkdir -p ~/.config/systemd/user
mkdir -p ~/.local/share/simplego

# Create service file
nano ~/.config/systemd/user/simplego.service

# Reload systemd user daemon (tells systemd to scan for new/changed files)
systemctl --user daemon-reload

# Enable service (equivalent to launchctl load - registers and auto-starts on login)
systemctl --user enable simplego.service

# Start service immediately (equivalent to launchctl start)
systemctl --user start simplego.service
```

---

### System-Level Service

**File:** `/etc/systemd/system/simplego.service`

```ini
[Unit]
Description=Simple Go System Service
After=multi-user.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/simplego/simplego
User=root
Group=root
StandardOutput=file:/var/log/simplego/stdout.log
StandardError=file:/var/log/simplego/stderr.log
RemainAfterExit=no

[Install]
WantedBy=multi-user.target
```

**Create the service:**
```bash
# Create log directory
sudo mkdir -p /var/log/simplego

# Create service file
sudo nano /etc/systemd/system/simplego.service

# Reload systemd daemon (tells systemd to scan for new/changed files)
sudo systemctl daemon-reload

# Enable service (equivalent to launchctl load - registers and auto-starts on boot)
sudo systemctl enable simplego.service

# Start service immediately (equivalent to launchctl start)
sudo systemctl start simplego.service
```

---

## 🟢 Useful systemd Commands

### User Services

| Command                                      | Description                        |
| -------------------------------------------- | ---------------------------------- |
| `systemctl --user start simplego.service`   | Start service                      |
| `systemctl --user stop simplego.service`    | Stop service                       |
| `systemctl --user restart simplego.service` | Restart service                    |
| `systemctl --user enable simplego.service`  | Enable (auto-start on login)      |
| `systemctl --user disable simplego.service` | Disable auto-start                 |
| `systemctl --user status simplego.service`  | Show service status                |
| `systemctl --user list-units`               | List all user services             |
| `journalctl --user -u simplego.service`     | View service logs                  |

### System Services

| Command                                 | Description                      |
| --------------------------------------- | -------------------------------- |
| `sudo systemctl start simplego.service`   | Start service                    |
| `sudo systemctl stop simplego.service`    | Stop service                     |
| `sudo systemctl restart simplego.service` | Restart service                  |
| `sudo systemctl enable simplego.service`  | Enable (auto-start on boot)     |
| `sudo systemctl disable simplego.service` | Disable auto-start               |
| `sudo systemctl status simplego.service`  | Show service status              |
| `sudo systemctl list-units`               | List all system services         |
| `sudo journalctl -u simplego.service`     | View service logs                |

---

## 🧰 Service File Explained

### Key Systemd Directives

| Directive         | Description                                    |
| ----------------- | ---------------------------------------------- |
| `[Unit]`          | General service information                    |
| `Description`     | Human-readable description                     |
| `After`           | Start after these targets                      |
| `[Service]`       | Service-specific settings                      |
| `Type`            | Service type (oneshot, simple, forking, etc.) |
| `ExecStart`       | Command to run                                 |
| `User`            | User to run as (system services only)         |
| `StandardOutput`  | Where to send stdout                           |
| `StandardError`   | Where to send stderr                           |
| `[Install]`       | Installation settings                          |
| `WantedBy`        | When to start the service                      |

### Service Types

| Type       | Description                                         |
| ---------- | --------------------------------------------------- |
| `simple`   | Main process, doesn't fork                          |
| `oneshot`  | Process runs once and exits (like our Go program)  |
| `forking`  | Service forks and main process exits               |
| `notify`   | Service sends notification when ready              |

---

## 🔍 Debugging and Monitoring

### View Logs
```bash
# User service logs
journalctl --user -u simplego.service -f

# System service logs
sudo journalctl -u simplego.service -f

# Last 50 lines
journalctl --user -u simplego.service -n 50

# Since specific time
journalctl --user -u simplego.service --since "2025-01-15 10:00:00"
```

### Check Service Status
```bash
# Detailed status
systemctl --user status simplego.service

# Is service active?
systemctl --user is-active simplego.service

# Is service enabled?
systemctl --user is-enabled simplego.service
```

---

## 🧹 Remove Service

### User Service
```bash
# Stop and disable
systemctl --user stop simplego.service
systemctl --user disable simplego.service

# Remove service file
rm ~/.config/systemd/user/simplego.service

# Reload daemon
systemctl --user daemon-reload

# Remove binary and logs (optional)
rm ~/bin/simplego
rm -rf ~/.local/share/simplego
```

### System Service
```bash
# Stop and disable
sudo systemctl stop simplego.service
sudo systemctl disable simplego.service

# Remove service file
sudo rm /etc/systemd/system/simplego.service

# Reload daemon
sudo systemctl daemon-reload

# Remove binary and logs (optional)
sudo rm /usr/local/bin/simplego
sudo rm -rf /var/log/simplego
```

---

## 🚦 Quick Reference: User vs System

| Goal                                    | Use                            |
| --------------------------------------- | ------------------------------ |
| Personal background task for user      | `~/.config/systemd/user/`      |
| System-wide service at boot            | `/etc/systemd/system/`         |
| Service that needs root privileges     | `/etc/systemd/system/`         |
| Service that runs only when user login | `~/.config/systemd/user/`      |

---

## 📝 Notes

* **User services** run in the user's session and can access user resources
* **System services** run as system daemons and can access system resources
* Use `%h` in user service files to refer to user's home directory
* User services require `systemctl --user` commands
* System services require `sudo systemctl` commands
* Always run `daemon-reload` after modifying service files

---

## 🔄 Comparison with macOS

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

### Why This Difference?

- **macOS**: You tell `launchctl` exactly where to find the `.plist` file
- **Linux**: systemd automatically scans predefined directories in priority order

### 📁 Systemd Service Discovery Directories

#### User Services (searched in this order):
1. `~/.config/systemd/user/` - **User-specific services** (highest priority)
2. `/etc/systemd/user/` - **System administrator defined user services**
3. `/run/systemd/user/` - **Runtime user services** (temporary)
4. `/usr/local/lib/systemd/user/` - **Locally installed user services**
5. `/usr/lib/systemd/user/` - **Package-installed user services** (lowest priority)

#### System Services (searched in this order):
1. `/etc/systemd/system/` - **System administrator services** (highest priority)
2. `/run/systemd/system/` - **Runtime system services** (temporary)
3. `/usr/local/lib/systemd/system/` - **Locally installed system services**
4. `/lib/systemd/system/` - **Package-installed system services**
5. `/usr/lib/systemd/system/` - **Package-installed system services** (lowest priority)

### 🔍 How systemd Finds Your Service:

When you run `systemctl --user enable simplego.service`, systemd:
1. Searches directories in the order listed above
2. Stops at the **first match** it finds
3. Uses that service file

**Example:** If you have `simplego.service` in both `~/.config/systemd/user/` and `/usr/lib/systemd/user/`, systemd will use the one in `~/.config/systemd/user/` because it has higher priority.

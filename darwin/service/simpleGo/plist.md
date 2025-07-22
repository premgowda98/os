## 🧠 What Is a `.plist` in macOS Services?

A `.plist` (property list) is an **XML file** that defines the configuration for a background service (called a **launchd job**).

Each `.plist` tells macOS:

* What program to run
* When to run it (at login, periodically, on-demand)
* How to run it (env vars, working dir, arguments)
* Where to store logs

---

## 🧰 Components Involved

| Component       | Purpose                                                                            |
| --------------- | ---------------------------------------------------------------------------------- |
| `launchd`       | The macOS system daemon that manages all services (similar to `systemd` on Linux). |
| `.plist` file   | The config that describes your service.                                            |
| `launchctl`     | The command-line tool to load/unload/start/stop/manage `launchd` jobs.             |
| `LaunchAgents`  | Directory for **user-level services** that run at login.                           |
| `LaunchDaemons` | Directory for **system-wide services** that run at boot (require sudo).            |

---

## 📁 Common Service Locations

| Location                  | For What                             | Auto-start  |
| ------------------------- | ------------------------------------ | ----------- |
| `~/Library/LaunchAgents/` | Runs **per user**, on login          | ✅ Yes       |
| `/Library/LaunchAgents/`  | All users, on login                  | ✅ Yes       |
| `/Library/LaunchDaemons/` | System-wide, on boot                 | ✅ Yes       |
| `/System/Library/...`     | Apple system services (do not touch) | ⚠️ Reserved |

> Your Go program running as `~/Library/LaunchAgents/com.user.homecheck.plist` is a **user agent** that will start **each time you log in**.

---

## ⚙️ What Does `launchctl load` Do?

```bash
launchctl load ~/Library/LaunchAgents/com.example.myservice.plist
```

This:

* Registers the `.plist` with `launchd`
* Immediately starts the service (if `RunAtLoad` is `true`)
* Adds it to the **user session's persistent list** so it will restart at next login

So YES — if you run `launchctl load` on a valid plist in `~/Library/LaunchAgents/`, your program will **persist across restarts**, as long as:

* You log in as that user
* The plist stays in place

---

## 🟢 Other Useful Commands

| Command                    | Description                           |                         |
| -------------------------- | ------------------------------------- | ----------------------- |
| `launchctl start <label>`  | Start service manually (after loaded) |                         |
| `launchctl stop <label>`   | Stop service without unloading it     |                         |
| `launchctl unload <plist>` | Unload and stop the service           |                         |
| `launchctl list`           | Show all loaded services              |                         |
| \`launchctl list           | grep homecheck\`                      | Find a specific service |
| `launchctl bootout`        | (macOS 13+) alternative to `unload`   |                         |

---

## 🔁 Summary of `.plist` Keys

| Key                    | Description                                  |
| ---------------------- | -------------------------------------------- |
| `Label`                | Unique name of the job (used in `launchctl`) |
| `ProgramArguments`     | Array of command and args to run             |
| `RunAtLoad`            | Run immediately on load/login                |
| `EnvironmentVariables` | Set env vars like `HOME`, `PATH`             |
| `StandardOutPath`      | File path for stdout                         |
| `StandardErrorPath`    | File path for stderr                         |

Example minimal `.plist`:

```xml
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.user.homecheck</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/homecheck/homecheck</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/usr/local/bin/homecheck/stdout.log</string>
    <key>StandardErrorPath</key>
    <string>/usr/local/bin/homecheck/stderr.log</string>
</dict>
</plist>
```

---

## 🧹 To Remove a Service

```bash
launchctl unload ~/Library/LaunchAgents/com.user.homecheck.plist
rm ~/Library/LaunchAgents/com.user.homecheck.plist
```

(Also delete logs or binary if needed)

---

## 📝 Final Notes

* `.plist` must be **in the correct folder** and have **`.plist` extension**
* Don't forget: if you move your binary, update the plist
* To test changes, always `unload` → edit → `load` again
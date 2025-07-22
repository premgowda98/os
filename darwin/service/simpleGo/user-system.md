**How macOS handles services** using `launchd`.

macOS has **two levels** of `LaunchAgents`, each with a different purpose and scope:

---

## 📁 1. `~/Library/LaunchAgents` — **User-Level Agents**

### ✅ Purpose:

* Runs **only for the logged-in user**
* Starts **when that user logs in**
* Runs with **user permissions** (not `root`)
* Cannot access system-level services or hardware directly

### 📌 Path:

```
/Users/yourname/Library/LaunchAgents/
```

### 👤 Example use case:

* A background app like Dropbox, Slack, or a personal sync script

---

## 📁 2. `/Library/LaunchAgents` — **System-Available User Agents**

### ✅ Purpose:

* Runs for **all users**, but **only after login**
* macOS loads these for **each user that logs in**
* Still runs **as that user**, not `root`

### 📌 Path:

```
/Library/LaunchAgents/
```

### 🧠 Example use case:

* A corporate login item or helper tool that should run for **every user on the system**

---

## 🔄 Summary Table

| Location                 | Runs As | When It Runs              | Scope                    |
| ------------------------ | ------- | ------------------------- | ------------------------ |
| `~/Library/LaunchAgents` | User    | When that user logs in    | Per-user                 |
| `/Library/LaunchAgents`  | User    | When **any** user logs in | All users (individually) |
| `/Library/LaunchDaemons` | `root`  | At boot                   | System-wide, background  |

---

## ⚠️ Notes

* **Agents** run in the user's login session and can interact with the GUI.
* **Daemons** run in the background as `root` and **do not** have access to GUI components (no `$HOME`, no dock, no display).
* Agents **do not run at boot**, only on login.

---

### 🚦 Use the right location depending on your goal:

| Goal                             | Use                      |
| -------------------------------- | ------------------------ |
| GUI app or helper per user       | `~/Library/LaunchAgents` |
| Script for all users after login | `/Library/LaunchAgents`  |
| Background system service (root) | `/Library/LaunchDaemons` |

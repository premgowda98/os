## Windows Registry Comprehensive Notes

### What Is the Windows Registry?
The Windows Registry is a hierarchical database that stores configuration settings and options for the operating system, hardware, user profiles, and installed applications. It is essential for Windows operation and application behavior.

### Main Registry Hives
- **HKEY_LOCAL_MACHINE (HKLM):** Contains system-wide settings, applies to all users. Used for hardware, OS, and global application settings.
- **HKEY_CURRENT_USER (HKCU):** Contains settings for the currently logged-in user. User-specific preferences and configurations.
- **HKEY_USERS (HKU):** Stores settings for all user profiles on the system. Each user has a subkey (e.g., S-1-5-21-... for domain users, S-1-5-18/19/20 for system accounts).
- **HKEY_CLASSES_ROOT (HKCR):** Contains file association and COM object information. Merged from HKLM and HKCU.
- **HKEY_CURRENT_CONFIG (HKCC):** Contains information about the current hardware profile.

### Registry Keys and SIDs
- **SIDs (Security Identifiers):** Unique identifiers for user and system accounts. Examples:
  - `S-1-5-18`: LocalSystem
  - `S-1-5-19`: NT Authority LocalService
  - `S-1-5-20`: NT Authority NetworkService
  - `S-1-5-21-...`: Domain or local user accounts
- These SIDs appear under HKU to represent user profiles and system accounts.

### Registry vs. Windows Services
- **Registry:** Stores configuration, preferences, and system/application settings. Does not run code, but influences how code runs.
- **Services:** Executable processes that run in the background, often using registry settings for configuration.

### Application Behavior and Registry Hives
- **HKLM:** If an application stores settings here, those settings apply to all users. Requires admin rights to modify.
- **HKCU:** Settings here are user-specific. No admin rights needed; changes affect only the current user.
- **HKU:** Used for managing multiple user profiles. Applications may read/write here for specific users.
- **Services:** Services may read configuration from HKLM (global) or HKU/HKCU (user-specific if running as a user service).

#### Application Startup Scenarios
- **HKLM:** Application starts for all users, often as a system service or scheduled task.
- **HKCU:** Application starts only for the current user, e.g., via user login or user-specific startup.
- **Services:** Can start independently of user login, based on service configuration.

### Registry Editing and Management
- Use `regedit.exe` for GUI editing.
- Use `reg.exe` for CLI editing: `reg add`, `reg delete`, `reg query`.
- Use PowerShell: `Get-ItemProperty`, `Set-ItemProperty`, `New-Item`, `Remove-Item`.

### Additional Notes
- Registry changes can affect system stability; backup before editing.
- Permissions on registry keys control who can read/write settings.
- Malware often targets registry for persistence.

### References
- [Microsoft Docs: Registry](https://learn.microsoft.com/en-us/windows/win32/sysinfo/registry)
- [Security Identifiers (SIDs)](https://learn.microsoft.com/en-us/windows/security/identity-protection/access-control/security-identifiers)

---
This note covers the essentials of the Windows Registry, hives, SIDs, application behavior, and differences from Windows Services. Expand as needed for your use case.

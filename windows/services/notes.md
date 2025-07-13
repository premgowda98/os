## Windows Services Comprehensive Notes

### What Are Windows Services?
Windows Services are long-running executable applications that run in the background on Windows operating systems. They can start automatically when the computer boots, run without user intervention, and do not require a user to be logged in. Services are used for tasks like networking, security, backups, and more.

### Behavior and User Context
- **Startup:** Services can be set to start automatically, manually, or be disabled.
- **User Context:** Services typically run under system accounts (like LocalSystem, NetworkService, or LocalService) but can also run under specific user accounts. They are not tied to a logged-in user and can run even when no user is logged in.
- **Session:** Services run in their own session, separate from user sessions, and usually do not have a user interface.

### Types of Service Accounts
- **LocalSystem:** High privileges, access to most system resources.
- **NetworkService:** Limited privileges, can access network resources as the computer.
- **LocalService:** Limited privileges, minimal access to local and network resources.
- **Custom User:** Can run under a specific user account for more control.

### Creating Services with NSSM
NSSM (Non-Sucking Service Manager) is a tool for managing Windows services, especially for running non-service applications as services.

#### About NSSM
- Open-source service manager for Windows.
- Allows any executable (e.g., scripts, apps) to run as a service.
- Handles service restarts, logging, and more.

#### How to Create a Service with NSSM
1. Download and extract NSSM from [https://nssm.cc/](https://nssm.cc/).
2. Open Command Prompt as Administrator.
3. Run: `nssm install <ServiceName>`
4. In the GUI, set the path to your executable and configure options.
5. Click "Install service".

#### How to Delete an NSSM Service
1. Open Command Prompt as Administrator.
2. Run: `nssm remove <ServiceName> confirm`

#### Other Useful NSSM Commands
- `nssm edit <ServiceName>`: Edit service configuration.
- `nssm start <ServiceName>`: Start the service.
- `nssm stop <ServiceName>`: Stop the service.

### Service Types and Differences
- **Local Services:** Run on the local machine, usually with limited network access.
- **Network Services:** Can interact with network resources, may run as NetworkService account.
- **User Services:** Run under a specific user account, useful for user-specific tasks.

### Key Differences
- **Privileges:** LocalSystem has the most, LocalService the least.
- **Network Access:** NetworkService is designed for network resource access.
- **Security:** Running as LocalSystem is riskier; prefer least privilege.

### Managing Services
- Use `services.msc` for GUI management.
- Use `sc` command for CLI management: `sc create`, `sc delete`, `sc start`, `sc stop`.
- Use PowerShell: `New-Service`, `Remove-Service`, `Start-Service`, `Stop-Service`.

### Additional Notes
- Services can be set to restart on failure.
- Dependencies can be configured so services start in order.
- Logging and recovery options are available in service properties.

### References
- [Microsoft Docs: Services](https://learn.microsoft.com/en-us/windows/win32/services/services)
- [NSSM Documentation](https://nssm.cc/usage)

---
This note covers the essentials of Windows Services, NSSM usage, service types, and management. Add more details as needed for your specific use case.

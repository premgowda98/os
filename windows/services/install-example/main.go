package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type myService struct{}

func (m *myService) Execute(args []string, r <-chan svc.ChangeRequest, s chan<- svc.Status) (bool, uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	s <- svc.Status{State: svc.StartPending}
	go runMainLogic()
	s <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	for c := range r {
		switch c.Cmd {
		case svc.Interrogate:
			s <- c.CurrentStatus
		case svc.Stop, svc.Shutdown:
			s <- svc.Status{State: svc.StopPending}
			return false, 0
		default:
		}
	}
	return false, 0
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--install":
			exePath, err := os.Executable()
			if err != nil {
				fmt.Println("Failed to get executable path:", err)
				os.Exit(1)
			}
			err = installService("MyGoService", exePath)
			if err != nil {
				fmt.Println("Failed to install service:", err)
				os.Exit(1)
			}
			fmt.Println("Service installed successfully.")
			return
		case "--uninstall":
			err := uninstallService("MyGoService")
			if err != nil {
				fmt.Println("Failed to uninstall service:", err)
				os.Exit(1)
			}
			fmt.Println("Service uninstalled successfully.")
			return
		}
	}

	isWinService, err := svc.IsWindowsService()
	if err != nil {
		isWinService = false
	}
	if isWinService {
		// Running as a Windows service
		svc.Run("MyGoService", &myService{})
		return
	}
	// Running as console app
	runMainLogic()
}

// Install the service using the service manager
func installService(name, exepath string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err == nil {
		s.Close()
		return fmt.Errorf("service %s already exists", name)
	}
	s, err = m.CreateService(name, exepath, mgr.Config{DisplayName: name})
	if err != nil {
		return err
	}
	defer s.Close()
	return nil
}

// Uninstall the service using the service manager
func uninstallService(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return err
	}
	defer s.Close()
	return s.Delete()
}

func runMainLogic() {
	exePath, err := os.Executable()
	logFilePath := "service_output.txt"
	if err == nil {
		exeDir := exePath
		if idx := strings.LastIndex(exePath, string(os.PathSeparator)); idx != -1 {
			exeDir = exePath[:idx]
		}
		logFilePath = exeDir + string(os.PathSeparator) + "service_output.txt"
	}
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
	} else {
		log.SetOutput(f)
		defer f.Close()
	}

	userName, _ := user.Current()
	log.Println("Current username:", userName.Username, "name:", userName.Name, "uid:", userName.Uid, "gid:", userName.Gid)
}

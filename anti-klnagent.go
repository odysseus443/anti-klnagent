package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"os"
)

func main() {
	killService("klnagent")
	err := os.RemoveAll("C:\\Program Files (x86)\\Kaspersky Lab\\NetworkAgent")
	if err != nil {
		return
	}
}

func killService(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return fmt.Errorf("Cannot connect to manager %v", err)
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("service %s does not exist", name)
	}
	defer s.Close()
	status, err := s.Query()
	if status.State != svc.Stopped {
		s.Control(svc.Stop)
	}
	if err != nil {
		return fmt.Errorf("could not start the service: %v", err)
	}
	return nil
}

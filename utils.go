package main

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunCommand(command string) {
	parts := strings.Fields(command)

	cmd := exec.Command(parts[0], parts[1:]...)

	cmd.Run()
}

func RunCommandErr(command string) error {
	parts := strings.Fields(command)

	cmd := exec.Command(parts[0], parts[1:]...)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func CheckIsRoot() bool {
	return os.Geteuid() == 0
}

func WithDots(message string, fn func()) {
	done := make(chan bool)

	go func() {
		dots := "."
		for {
			select {
			case <-done:
				logMessage := "\n" + message + "✔️"
				LogMessage(logMessage, Bold, Green)
				return
			default:
				dots += "."
				if len(dots) > 3 {
					dots = "."
				}
				logMessage := "\r" + message + dots
				LogMessagef(logMessage, Normal, Yellow)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	fn()
	done <- true
	time.Sleep(100 * time.Millisecond)
}

func GetUserHome() string {
	user := os.Getenv("SUDO_USER")
	home := "/home/" + user
	return home
}

func CheckDir(dir string) bool {
	if _, err := os.Stat(dir); err == nil {
		return true
	}

	return false
}

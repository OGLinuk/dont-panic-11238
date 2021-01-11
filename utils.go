package main

import (
	"os"
	"os/exec"
)

// checkExists always returns false, unless path exists
func checkExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// runCmd wraps exec.Command
func runCmd(cmd string, args ...string) error {
	execCmd := exec.Command(cmd, args...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	if err = execCmd.Run(); err != nil {
		return err
	}
	return nil
}

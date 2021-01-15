package main

import (
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
	"os/exec"
)

// checkExists always returns false, unless path exists
func checkExists(path string) bool {
	if _, err = os.Stat(path); err == nil {
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

// crc32Checksum of path
func crc32Checksum(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	poly := crc32.MakeTable(0xEDB88320)
	hash := crc32.New(poly)

	if _, err = io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

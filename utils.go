package dontpanic

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"os"
	"os/exec"
)

var (
	// ROOTDIR is a directory containing the manifests and services directories used by dont-panic-11238
	ROOTDIR = "DONTPANIC"
	// MANIFESTDIR is a directory containing the manifest files
	MANIFESTSDIR = fmt.Sprintf("%s/%s", ROOTDIR, "manifests")
	// SERVICEDIR is a directory containing the git repositories
	SERVICESDIR = fmt.Sprintf("%s/%s", ROOTDIR, "services")

	err error
)

// CheckExists always returns false, unless path exists
func CheckExists(path string) bool {
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

// Checksum of path using a crc32 hash
func Checksum(path string) (string, error) {
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

// GenerateFile path if it does not exist with content
func GenerateFile(path string, content []byte) error {
	if !CheckExists(path) {
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("gomain.go::os.Create(%s)::ERROR: %s", path, err.Error())
		}
		defer f.Close()

		f.Write(content)
	}
	return nil
}

// GetLocalIP of hostmachine
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			return ipnet.IP.String()
		}
	}
	return ""
}

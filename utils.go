package dontpanic

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

var (
	// ROOTDIR is a directory containing the manifests and services directories used by dont-panic-11238
	ROOTDIR = "DONTPANIC"
	// MANIFESTDIR is a directory containing the manifest files
	MANIFESTSDIR = fmt.Sprintf("%s/%s", ROOTDIR, "manifests")
	// SERVICEDIR is a directory containing the git repositories
	SERVICESDIR = fmt.Sprintf("%s/%s", ROOTDIR, "services")

	err         error
	activePorts []string
	wg          = &sync.WaitGroup{}
	timeTaken   time.Duration
	timeSince   time.Time
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

package main

import "os"

// checkExists always returns false, unless path exists
func checkExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

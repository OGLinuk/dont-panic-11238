package main

import "testing"

func TestScanLocalhost(t *testing.T) {
	actual := ScanLocalhost()
	if len(actual) < 1 {
		t.Fatalf("Expected len() <= 1; Got len() < 1 ...\n")
	}
}

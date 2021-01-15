package main

import (
	"os"
	"testing"
)

var (
	fname = "Dockerfile"
)

func TestGenerateDockerfile(t *testing.T) {
	// crc32 checksum of rwx.gg-50501
	expected := "33d74d0d"

	GenerateDockerfile("rwx.gg", "50501", ".")

	if _, err = os.Stat(fname); err != nil {
		t.Fatalf("Expected nil; Got: %s", err.Error())
	} else {
		actual, err := crc32Checksum(fname)
		os.RemoveAll(fname)
		if err != nil {
			t.Fatalf("crc32Checksum::ERROR: %s", err.Error())
		}

		if actual != expected {
			t.Fatalf("\nExpected: %s\nGot: %s\n", expected, actual)
		}
	}
}

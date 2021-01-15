package main

import (
	"os"
	"testing"
)

var (
	dockercomposefn = "docker-compose.yml"
)

func TestGenerateDockerCompose(t *testing.T) {
	manifestsPath := "DONTPANIC/manifests"

	// crc32 checksum of docker-compose.yml
	expected := "5b712562"

	if _, err = os.Stat(manifestsPath); err != nil {
		GenerateManifests("default")
	}

	GenerateDockerCompose()

	if _, err = os.Stat(dockercomposefn); err != nil {
		t.Fatalf("Expected nil; Got: %s", err.Error())
	} else {
		actual, err := crc32Checksum(dockercomposefn)
		//os.RemoveAll(dockercomposefn)
		if err != nil {
			t.Fatalf("crc32Checksum::ERROR: %s", err.Error())
		}

		if actual != expected {
			t.Fatalf("\nExpected: %s\nGot: %s\n", expected, actual)
		}
	}
}

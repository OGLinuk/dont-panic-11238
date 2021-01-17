package dontpanic

import (
	"fmt"
	"log"
	"os"
)

// CreateManifestFile called name which contains defaults given
func CreateManifestFile(name string, defaults []string) {
	manif := fmt.Sprintf("%s/%s", MANIFESTSDIR, name)
	f, err := os.Create(manif)
	if err != nil {
		// TODO: better ...
		log.Printf("manifests.go::os.Create(%s)::ERROR: %s", manif, err.Error())
	}
	defer f.Close()

	for _, d := range defaults {
		f.WriteString(fmt.Sprintf("%s\n", d))
	}
}

// GenerateManifests directory && create default manifest files defined in e
func GenerateManifests(e string) error {
	if !CheckExists(MANIFESTSDIR) {
		if err = os.MkdirAll(MANIFESTSDIR, 0744); err != nil {
			return fmt.Errorf("manifests.go::os.MkdirAll::ERROR: %s", err.Error())
		}

		for manifestname, defaults := range DefaultEnvs[e] {
			CreateManifestFile(manifestname, defaults)
		}
	}
	return nil
}

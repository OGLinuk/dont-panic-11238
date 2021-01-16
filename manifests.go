package dontpanic

import (
	"fmt"
	"log"
	"os"
)

// CreateManifestFile called name which contains defaults given
func CreateManifestFile(name string, defaults []string) {
	defer wg.Done()

	manif := fmt.Sprintf("%s/%s", MANIFESTSDIR, name)
	f, err := os.Create(manif)
	if err != nil {
		log.Printf("manifests.go::os.Create(%s)::ERROR: %s", manif, err.Error())
	}
	defer f.Close()

	for _, d := range defaults {
		f.WriteString(fmt.Sprintf("%s\n", d))
	}
}

// GenerateManifests directory && create default manifest files defined in e
func GenerateManifests(e string) {
	if _, err = os.Stat(MANIFESTSDIR); err != nil {
		log.Println("`manifests` dir not found, generating manifests ...")
		if _, err = os.Stat(MANIFESTSDIR); err != nil {
			if err = os.MkdirAll(MANIFESTSDIR, 0744); err != nil {
				log.Fatalf("manifests.go::os.MkdirAll::ERROR: %s", err.Error())
			}

			for manifestname, defaults := range DefaultEnvs[e] {
				wg.Add(1)
				go CreateManifestFile(manifestname, defaults)
			}
			wg.Wait()
		}
	}
}

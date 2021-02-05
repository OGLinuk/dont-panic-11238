package dontpanic

import (
	"fmt"
	"os"
)

// CreateManifestFile called name which contains defaults given
func CreateManifestFile(name string, defaults []string) error {
	manif := fmt.Sprintf("%s/%s", MANIFESTSDIR, name)
	f, err := os.Create(manif)
	if err != nil {
		return fmt.Errorf("CreateManifestFile::os.Create(%s): %s", manif, err.Error())
	}
	defer f.Close()

	for _, d := range defaults {
		f.WriteString(fmt.Sprintf("%s\n", d))
	}
	return nil
}

// GenerateManifests directory && create default manifest files defined in env
func GenerateManifests(env string) error {
	if !CheckExists(MANIFESTSDIR) {
		if err = os.MkdirAll(MANIFESTSDIR, 0744); err != nil {
			return fmt.Errorf("GenerateManifests::os.MkdirAll(%s): %s", MANIFESTSDIR, err.Error())
		}

		for manifestname, defaults := range DefaultEnvs[env] {
			if err = CreateManifestFile(manifestname, defaults); err != nil {
				return err
			}
		}
	}
	return nil
}

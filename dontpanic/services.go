package dontpanic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// GenerateService reads the manifest file (manifestPath) contents, then for
// every service check if it exists. If it doesnt; clone it and call the passed
// fn using the values: name, port, and git repo path. The fn func passed
// should contain the logic to generate everything needed for the type of
// service (ex: passing GenerateFileServer for a service of type docs).
func GenerateService(manifestName, manifestPath string, fn func(name, port, link string) error) error {
	f, err := os.Open(manifestPath)
	if err != nil {
		return fmt.Errorf("services.go::os.Open(%s)::ERROR: %s", manifestPath, err.Error())
	}
	defer f.Close()

	bs := bufio.NewScanner(f)
	for bs.Scan() {
		service := bs.Text()
		sc := strings.Split(service, " ")

		if len(sc) != 3 {
			return fmt.Errorf("service [%s] does not have all 3 required components <name> <port> <link>", service)
		}

		// Ex: services/functions/sbh-9001
		toCheck := fmt.Sprintf("%s/%s/%s-%s", SERVICESDIR, manifestName, sc[0], sc[1])
		go func() error {
			// TODO: Change below to use go git
			if CheckExists(toCheck) == false {
				log.Printf("services.go: %s does not exists, [CLONING] now ...", toCheck)
				if err = runCmd("git", "clone", sc[2], toCheck); err != nil {
					// TODO: Retry X times before resulting to error
					return fmt.Errorf("GenerateService::runCmd::gitCLONE: %s", err.Error())
				}

				if fn != nil {
					fn(sc[0], sc[1], toCheck)
				}
			} else {
				log.Printf("services.go: %s exists, [UPDATING] now ...", toCheck)
				if err = runCmd("git", "-C", toCheck, "pull"); err != nil {
					return fmt.Errorf("GenerateService::runCmd::gitPULL: %s", err.Error())
				}
			}
			return nil
		}()
	}
	return nil
}

// GenerateServices defined in all manifest files located in MANIFESTSDIR
func GenerateServices() error {
	// Ensure `services` directory exists
	if !CheckExists(SERVICESDIR) {
		if err = os.MkdirAll(SERVICESDIR, 0744); err != nil {
			return fmt.Errorf("GenerateServices::os.MkdirAll(%s): %s", SERVICESDIR, err.Error())
		}
	}

	// TODO: Find a way to compress opening and iterating over the content of manifest files
	// Get all manifest files
	manifests, err := ioutil.ReadDir(MANIFESTSDIR)
	if err != nil {
		// TODO: Generate defaults X times before erroring
		return fmt.Errorf("GenerateServices::ioutil.ReadDir(%s): %s", MANIFESTSDIR, err.Error())
	}

	// For every manifest file, read its contents and download/update entries
	for _, manifest := range manifests {
		manifestName := manifest.Name()

		servicesPath := fmt.Sprintf("%s/%s", SERVICESDIR, manifestName)
		if _, err = os.Stat(servicesPath); err != nil {
			// TODO: When `readme.go` is implemented, create a default README.md file in servicesPath dir
			// if one does not exists, then render with `doc.go`
			if err = os.MkdirAll(servicesPath, 0744); err != nil {
				return fmt.Errorf("GenerateServices::os.MkdirAll(%s):%s", servicesPath, err.Error())
			}
		}

		manifestPath := fmt.Sprintf("%s/%s", MANIFESTSDIR, manifestName)
		go func(mp, manifestName string) error {
			var maniFunc func(name, port, link string) error
			if manifestName == "docs" {
				maniFunc = GenerateFileServer
			} else if manifestName == "individuals-blog" {
				maniFunc = GenerateFileServer
			} else if manifestName == "projects-blog" {
				maniFunc = GenerateFileServer
			} else if manifestName == "fileserver" {
				maniFunc = GenerateFileServer
			} else {
				maniFunc = nil
			}
			if err = GenerateService(manifestName, mp, maniFunc); err != nil {
				return err
			}
			return nil
		}(manifestPath, manifestName)
	}
	return nil
}

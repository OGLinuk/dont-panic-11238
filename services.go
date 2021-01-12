package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// genService reads the manifest file located at sPath, git clones, then
// invokes the functions corresponding to the sType
func genService(sType, sPath string) {
	f, err := os.Open(sPath)
	if err != nil {
		// TODO: Improve ...
		log.Printf("genService::os.Open(%s)::ERROR: %s", sPath, err.Error())
	}
	defer f.Close()

	bs := bufio.NewScanner(f)
	for bs.Scan() {
		service := bs.Text()
		sc := strings.Split(service, " ")

		if len(sc) != 3 {
			log.Fatal(fmt.Errorf(`service [%s] does not have all 3 required components <name> <port> <link>`, service))
		}

		// Ex: services/sbh-9001
		toCheck := fmt.Sprintf("%s/%s-%s", SERVICESDIR, sc[0], sc[1])
		result := checkExists(toCheck)

		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: Change below to use go git
			if result == false {
				log.Printf("services.go: %s does not exists, [CLONING] now ...", toCheck)
				if err = runCmd("git", "clone", sc[2], toCheck); err != nil {
					// TODO: Retry X times before resulting to error
					log.Printf("services::runCmd::gitCLONE::ERROR: %s", err.Error())
				}

				if sType == "fileserver" {
					NewFileServer(sc[0], sc[1])
				} else {
					// TODO: Determine which service types require what
				}
			} else {
				log.Printf("services.go: %s exists, [UPDATING] now ...", toCheck)
				if err = runCmd("git", "-C", toCheck, "pull", "origin", "master"); err != nil {
					log.Printf("services.go::runCmd::ERROR:%s", err.Error())
				}
			}
		}()
	}
}

func init() {
	flag.Parse()

	// Ensure `services` directory exists
	if _, err = os.Stat(SERVICESDIR); err != nil {
		if err = os.MkdirAll(SERVICESDIR, 0744); err != nil {
			log.Fatalf("services.go::checkExists(%s)::ERROR: %s", SERVICESDIR, err.Error())
		}
	}

	f, err := os.OpenFile("gen.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Could not open log file ...")
	}

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

// GenerateServices defined in all manifest files located in MANIFESTSDIR
func GenerateServices() {
	// Get all manifest files
	manifests, err := ioutil.ReadDir(MANIFESTSDIR)
	if err != nil {
		// TODO: Generate defaults X times before erroring
		log.Printf("services.go::ioutil.ReadDir(%s)::ERROR: %s", MANIFESTSDIR, err.Error())
	}

	// For every manifest file, read its contents and download/update entries
	for _, manifest := range manifests {
		manifestName := manifest.Name()
		manifestPath := fmt.Sprintf("%s/%s", MANIFESTSDIR, manifestName)

		wg.Add(1)
		go func(mp, manifestName string) {
			defer wg.Done()
			// TODO: Do better ... switch statement?
			if manifestName == "blogs" || manifestName == "docs" || manifestName == "fileserver" {
				genService("fileserver", mp)
			} else {
				genService(manifestName, mp)
			}
		}(manifestPath, manifestName)
	}
}

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

// GenerateService reads the manifest file (manifestPath) contents, then for
// every service check if it exists. If it doesnt; clone it and call the passed
// fn using the values: name, port, and git repo path. The fn func passed
// should contain the logic to generate everything needed for the type of
// service (ex: passing GenerateFileServer for a service of type docs).
func GenerateService(manifestName, manifestPath string, fn func(name, port, link string)) {
	f, err := os.Open(manifestPath)
	if err != nil {
		// TODO: Improve ...
		log.Printf("services.go::os.Open(%s)::ERROR: %s", manifestPath, err.Error())
	}
	defer f.Close()

	bs := bufio.NewScanner(f)
	for bs.Scan() {
		service := bs.Text()
		sc := strings.Split(service, " ")

		if len(sc) != 3 {
			log.Fatal(fmt.Errorf(`service [%s] does not have all 3 required components <name> <port> <link>`, service))
		}

		// Ex: services/functions/sbh-9001
		toCheck := fmt.Sprintf("%s/%s/%s-%s", SERVICESDIR, manifestName, sc[0], sc[1])
		result := checkExists(toCheck)

		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: Change below to use go git
			if result == false {
				log.Printf("services.go: %s does not exists, [CLONING] now ...", toCheck)
				if err = runCmd("git", "clone", sc[2], toCheck); err != nil {
					// TODO: Retry X times before resulting to error
					log.Printf("services.go::runCmd::gitCLONE::ERROR: %s", err.Error())
				}

				if fn != nil {
					fn(sc[0], sc[1], toCheck)
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
	// TODO: Find a way to compress opening and iterating over the content of manifest files
	// Get all manifest files
	manifests, err := ioutil.ReadDir(MANIFESTSDIR)
	if err != nil {
		// TODO: Generate defaults X times before erroring
		log.Printf("services.go::ioutil.ReadDir(%s)::ERROR: %s", MANIFESTSDIR, err.Error())
	}

	// For every manifest file, read its contents and download/update entries
	for _, manifest := range manifests {
		manifestName := manifest.Name()

		servicesPath := fmt.Sprintf("%s/%s", SERVICESDIR, manifestName)
		if _, err = os.Stat(servicesPath); err != nil {
			// TODO: When `readme.go` is implemented, create a default README.md file in servicesPath dir
			if err = os.MkdirAll(servicesPath, 0744); err != nil {
				log.Printf("services.go::os.MkdirAll(%s)::ERROR:%s", servicesPath, err.Error())
			}
		}

		manifestPath := fmt.Sprintf("%s/%s", MANIFESTSDIR, manifestName)
		wg.Add(1)
		go func(mp, manifestName string) {
			var maniFunc func(name, port, link string)

			defer wg.Done()
			// TODO: Figure out way to dynamically defined
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

			GenerateService(manifestName, mp, maniFunc)
		}(manifestPath, manifestName)
	}
	wg.Wait()
}

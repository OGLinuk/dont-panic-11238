package dontpanic

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type dockerservice struct {
	Hostname      string
	ContainerName string `yaml:"container_name"`
	Build         string
	Ports         []string
	Networks      []string
}

type dockercompose struct {
	Version  string
	Services map[string]dockerservice
	Networks map[string]interface{}
}

var (
	// DOCKERCOMPOSEVERSION is the version used in the generated docker-compose.yml
	DOCKERCOMPOSEVERSION = "3.2"
	dcFilename           = "docker-compose.yml"
)

func appendService(dc *dockercompose, ds dockerservice) {
	dc.Services[ds.Hostname] = ds
}

// GenerateDockerCompose file with all entries of all manifest files in DONTPANIC/manifests
func GenerateDockerCompose() {
	dc := &dockercompose{
		Version:  DOCKERCOMPOSEVERSION,
		Services: make(map[string]dockerservice),
		Networks: map[string]interface{}{
			"tiered": nil,
		},
	}

	dcf, err := os.Create(dcFilename)
	if err != nil {
		log.Printf("dockercompose.go::os.Create(%s)::ERROR: %s", dcFilename, err.Error())
	}
	defer dcf.Close()

	// TODO: Find a way to compress opening and iterating over the content of manifest files
	// Get all manifest files
	manifests, err := ioutil.ReadDir(MANIFESTSDIR)
	if err != nil {
		// TODO: Generate defaults X times before erroring
		log.Printf("dockercompose.go::ioutil.ReadDir(%s)::ERROR: %s", MANIFESTSDIR, err.Error())
	}

	// For every manifest file, read its contents and download/update entries
	for _, manifest := range manifests {
		manifestName := manifest.Name()
		// Skip manifest files like games that dont use Docker // TODO: Better ...
		if manifestName != "games" {
			manifestPath := fmt.Sprintf("%s/%s", MANIFESTSDIR, manifestName)

			manif, err := os.Open(manifestPath)
			if err != nil {
				// TODO: Improve ...
				log.Printf("dockercompose.go::os.Open(%s)::ERROR: %s", manifestPath, err.Error())
			}
			defer manif.Close()

			bs := bufio.NewScanner(manif)
			for bs.Scan() {
				service := bs.Text()
				sc := strings.Split(service, " ")

				if len(sc) != 3 {
					log.Fatal(fmt.Errorf(`service [%s] does not have all 3 required components <name> <port> <link>`, service))
				}

				// TODO: Scan the repo for Dockerfile(s) if not in root
				// URGENT as will break if any other services used are like sbh
				var bp string
				if sc[0] == "sbh" {
					bp = fmt.Sprintf("./DONTPANIC/services/%s/%s-%s/examples/web", manifestName, sc[0], sc[1])
				} else {
					bp = fmt.Sprintf("./DONTPANIC/services/%s/%s-%s", manifestName, sc[0], sc[1])
				}
				appendService(dc, dockerservice{
					Hostname:      sc[0],
					ContainerName: fmt.Sprintf("%s-container", sc[0]),
					Build:         bp,
					Ports:         []string{fmt.Sprintf("%s:%s", sc[1], sc[1])},
					Networks:      []string{"tiered"},
				})
			}
		}
	}

	data, err := yaml.Marshal(&dc)
	if err != nil {
		log.Printf("dockercompose.go::yaml.Marshal::ERROR: %s", err.Error())
	}

	dcf.Write(data)
}

package main

import (
	"fmt"
	"log"
	"os"
)

var (
	// BASEIMAGE used
	BASEIMAGE = "golang:1.15"
)

// GenerateDockerfile using the given name and port values
func GenerateDockerfile(name, port, path string) {
	dockerfile, err := os.Create(fmt.Sprintf("%s/Dockerfile", path))
	if err != nil {
		log.Fatalf("dockerfile.go::os.Create(Dockerfile)::ERROR: %s", err.Error())
	}
	defer dockerfile.Close()

	log.Printf("[docker.go] Generated Dockerfile at %s ...", path)
	fullServiceName := fmt.Sprintf("%s-%s", name, port)

	// TODO: better ...
	dockerfile.WriteString(fmt.Sprintf("FROM %s\n", BASEIMAGE))
	dockerfile.WriteString(fmt.Sprintf("ADD . /go/src/%s\n", fullServiceName))
	dockerfile.WriteString(fmt.Sprintf("WORKDIR /go/src/%s\n", fullServiceName))
	dockerfile.WriteString(fmt.Sprintf("RUN go build -o %s\n", fullServiceName))
	dockerfile.WriteString(fmt.Sprintf("EXPOSE %s\n", fmt.Sprintf("%s:%s", port, port)))
	dockerfile.WriteString(fmt.Sprintf("CMD [\"./%s\"]", fullServiceName))
}

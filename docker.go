package main

import (
	"fmt"
	"log"
	"os"
)

// GenerateDockerfile using the given name and port values
func GenerateDockerfile(name, port, path string) {
	dockerfile, err := os.Create(fmt.Sprintf("%s/Dockerfile", path))
	if err != nil {
		log.Fatalf("dockerfile.go::os.Create(Dockerfile)::ERROR: %s", err.Error())
	}
	defer dockerfile.Close()

	log.Printf("Generated Dockerfile at %s ...", path)
	fullServiceName := fmt.Sprintf("%s-%s", name, port)

	// TODO: Abstract behind a way to specify fill-in areas (ie FROM <base>)
	dockerfile.WriteString(fmt.Sprintf(`FROM golang:1.15
ADD . /go/src/%s
WORKDIR /go/src/%s
RUN go build -o %s-container
EXPOSE %s
CMD ["./%s-container"]`,
		fullServiceName, fullServiceName, fullServiceName,
		fmt.Sprintf("%s:%s", port, port), fullServiceName))
}

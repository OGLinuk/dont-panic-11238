package dontpanic

import (
	"bytes"
	"fmt"
)

var (
	// BASEIMAGE used
	BASEIMAGE = "golang:1.15"
)

// GenerateDockerfile using the given name and port values
func GenerateDockerfile(name, port, path string) error {
	var buffer bytes.Buffer
	fullServiceName := fmt.Sprintf("%s-%s", name, port)

	// TODO: better ...
	buffer.WriteString(fmt.Sprintf("FROM %s\n", BASEIMAGE))
	buffer.WriteString(fmt.Sprintf("ADD . /go/src/%s\n", fullServiceName))
	buffer.WriteString(fmt.Sprintf("WORKDIR /go/src/%s\n", fullServiceName))
	buffer.WriteString(fmt.Sprintf("RUN go build -o %s\n", fullServiceName))
	buffer.WriteString(fmt.Sprintf("EXPOSE %s\n", fmt.Sprintf("%s:%s", port, port)))
	buffer.WriteString(fmt.Sprintf("CMD [\"./%s\"]", fullServiceName))

	GenerateFile(fmt.Sprintf("%s/Dockerfile", path), buffer.Bytes())
	return nil
}

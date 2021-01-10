package main

import (
	"fmt"
	"log"
	"os"
)

// TODO: Abstract the entire Dockerfile to allow for more granular config
type dockerfile struct {
	Name string
	Port string
}

func NewDockerfile(name, port string) *dockerfile {
	return &dockerfile{
		Name: name,
		Port: port,
	}
}

// TODO: Allow for more configurability
func (df *dockerfile) GenerateDockerfile() {
	serviceSourcePath := fmt.Sprintf("%s/%s-%s", SERVICESDIR, df.Name, df.Port)

	// TODO: Abstract main.go generation to its own goscript
	log.Println(serviceSourcePath)
	dockerfile, err := os.Create(fmt.Sprintf("%s/Dockerfile", serviceSourcePath))
	if err != nil {
		log.Fatalf("os.Create(Dockerfile)::ERROR: %s", err.Error())
	}
	defer dockerfile.Close()

	fullServiceName := fmt.Sprintf("%s-%s", df.Name, df.Port)

	// TODO: Abstract behind a way to specify fill-in areas (ie FROM <base>)
	dockerfile.WriteString(fmt.Sprintf(`FROM golang:1.15
ADD . /go/src/%s
WORKDIR /go/src/%s
RUN go build -o %s-container
EXPOSE %s
CMD ["./%s-container"]`,
		fullServiceName, fullServiceName, fullServiceName,
		fmt.Sprintf("%s:%s", df.Port, df.Port), fullServiceName))
}

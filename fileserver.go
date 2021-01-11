package main

import (
	"fmt"
	"log"
	"os"
)

// fileserver is anything that requires http.FileServer
type fileserver struct {
	Name string
	Port string
}

// NewFileServer constructor
func NewFileServer(name, port string) {
	fs := &fileserver{
		Name: name,
		Port: port,
	}
	fs.generate()
}

// generate a main.go file at serviceSourcePath
func (fs *fileserver) generate() {
	serviceSourcePath := fmt.Sprintf("%s/%s-%s", SERVICESDIR, fs.Name, fs.Port)

	NewDockerfile(fs.Name, fs.Port)

	f, err := os.Create(fmt.Sprintf("%s/main.go", serviceSourcePath))
	if err != nil {
		log.Fatalf("fileserver.go::os.Create(%s/main.go)::ERROR: %s", serviceSourcePath, err.Error())
	}
	defer f.Close()

	// TODO: Refactor ...
	f.WriteString(fmt.Sprintf(`package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	NAME = "%s"
	PORT = %s
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Printf("Serving %s on %s ...", NAME, PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", PORT), nil))
}`, fs.Name, fs.Port, "%s", "%d", "%d"))
}

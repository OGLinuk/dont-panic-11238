package main

import (
	"fmt"
	"log"
	"os"
)

// GenerateFileServer creates Dockerfile and main.go files using the given name and port at path
func GenerateFileServer(name, port, path string) {
	GenerateDockerfile(name, port, path)

	f, err := os.Create(fmt.Sprintf("%s/main.go", path))
	if err != nil {
		log.Fatalf("fileserver.go::os.Create(%s/main.go)::ERROR: %s", path, err.Error())
	}
	defer f.Close()

	// TODO: Same as dockerfile.go ...
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
}`, name, port, "%s", "%d", "%d"))

	// TODO: Do check to ensure the Dockerfile and main.go file were created
}

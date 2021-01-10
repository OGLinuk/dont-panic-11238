package main

import (
	"fmt"
	"log"
	"os"
)

// TODO: Abstract the entire Dockerfile to allow for more granular config
type fileserver struct {
	Name string
	Port string
}

func NewFileserver(name, port string) *fileserver {
	return &fileserver{
		Name: name,
		Port: port,
	}
}

func (fs *fileserver) ScaffoldFileserver() {
	serviceSourcePath := fmt.Sprintf("%s/%s-%s", SERVICESDIR, fs.Name, fs.Port)

	df := NewDockerfile(fs.Name, fs.Port)
	df.GenerateDockerfile()

	maindotgo, err := os.Create(fmt.Sprintf("%s/main.go", serviceSourcePath))
	if err != nil {
		log.Fatalf("os.Create(%s/main.go)::ERROR: %s", serviceSourcePath, err.Error())
	}
	defer maindotgo.Close()

	maindotgo.WriteString(fmt.Sprintf(`package main

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

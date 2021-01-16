package dontpanic

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

var (
	defaultMain = map[string][]string{
		"import": []string{
			"\"fmt\"",
			"\"log\"",
			"\"net/http\"",
		},
		"main": []string{
			"http.Handle(\"/\", http.FileServer(http.Dir(\".\")))",
			"log.Printf(\"Serving %s on %d ...\", NAME, PORT)",
			"log.Fatal(http.ListenAndServe(fmt.Sprintf(\"0.0.0.0:%d\", PORT), nil))",
		},
	}
)

// GenerateMain .go file if one does not exist with content
func GenerateMain(path string, content []byte) {
	mainpath := fmt.Sprintf("%s/main.go", path)
	if checkExists(mainpath) == false {
		maindotgo, err := os.Create(mainpath)
		if err != nil {
			log.Printf("gomain.go::os.Create(%s)::ERROR: %s", mainpath, err.Error())
		}
		defer maindotgo.Close()

		maindotgo.Write(content)
	}
}

// GenerateFileServer creates a Dockerfile and a main.go file using the
// given name, port, and path
func GenerateFileServer(name, port, path string) {
	GenerateDockerfile(name, port, path)

	var buffer bytes.Buffer

	// TODO: Better ... There must be a way to store a template that name and
	// port can be passed to, then written to maindotgo
	buffer.WriteString("package main\n")

	buffer.WriteString("\nimport (\n")
	for _, imp := range defaultMain["import"] {
		buffer.WriteString(fmt.Sprintf("\t%s\n", imp))
	}
	buffer.WriteString("\n)")

	buffer.WriteString("\nvar (\n")
	buffer.WriteString(fmt.Sprintf("\tNAME = \"%s\"\n", name))
	buffer.WriteString(fmt.Sprintf("\tPORT = %s", port))
	buffer.WriteString("\n)\n")

	buffer.WriteString("\nfunc main() {\n")
	for _, imp := range defaultMain["main"] {
		buffer.WriteString(fmt.Sprintf("\t%s\n", imp))
	}
	buffer.WriteString("}\n")

	GenerateMain(path, buffer.Bytes())

	// TODO: Do check to ensure the Dockerfile and main.go file were created;
	// if not, try X times before resorting to error
}

package dontpanic

import (
	"bytes"
	"fmt"
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

	GenerateFile(fmt.Sprintf("%s/main.go", path), buffer.Bytes())

	// TODO: Do check to ensure the Dockerfile and main.go file were created;
	// if not, try X times before resorting to error
}

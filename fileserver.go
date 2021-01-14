package main

import (
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

// GenerateFileServer creates a Dockerfile and a main.go file using the
// given name, port, and path
func GenerateFileServer(name, port, path string) {
	GenerateDockerfile(name, port, path)

	maindotgo, err := os.Create(fmt.Sprintf("%s/main.go", path))
	if err != nil {
		log.Fatalf("fileserver.go::os.Create(%s/main.go)::ERROR: %s", path, err.Error())
	}
	defer maindotgo.Close()

	// TODO: Better ... There must be a way to store a template that name and
	// port can be passed to, then written to maindotgo
	maindotgo.WriteString("package main\n")

	maindotgo.WriteString("\nimport (\n")
	for _, imp := range defaultMain["import"] {
		maindotgo.WriteString(fmt.Sprintf("\t%s\n", imp))
	}
	maindotgo.WriteString("\n)")

	maindotgo.WriteString("\nvar (\n")
	maindotgo.WriteString(fmt.Sprintf("\tNAME = \"%s\"\n", name))
	maindotgo.WriteString(fmt.Sprintf("\tPORT = %s", port))
	maindotgo.WriteString("\n)\n")

	maindotgo.WriteString("\nfunc main() {\n")
	for _, imp := range defaultMain["main"] {
		maindotgo.WriteString(fmt.Sprintf("\t%s\n", imp))
	}
	maindotgo.WriteString("}\n")

	// TODO: Do check to ensure the Dockerfile and main.go file were created;
	// if not, try X times before resorting to error
}

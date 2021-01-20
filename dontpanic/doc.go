package dontpanic

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GenerateDoc for every markdown file in root
func GenerateDoc(root string) error {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			log.Printf("PATH: %s", path)
			if strings.Contains(path, "README.md") {
				if err = runCmd("pandoc", path, "-o", fmt.Sprintf("%s/index.html", filepath.Dir(path))); err != nil {
					return err
				}
			} else if strings.Contains(path, ".md") {
				splitpath := strings.Split(path, ".")[0]
				if err = runCmd("pandoc", splitpath, "-o", fmt.Sprintf("%s/%s.html", filepath.Dir(path), splitpath)); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

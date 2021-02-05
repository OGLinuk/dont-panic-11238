package dontpanic

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Render src as target with pandoc
func Render(src, target string) error {
	if err = runCmd("pandoc", src, "-o", target); err != nil {
		return err
	}
	return nil
}

// GenerateDoc for every file in root
func GenerateDoc(root string) error {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if strings.Contains(path, "README.md") {
				target := fmt.Sprintf("%s/index.html", filepath.Dir(path))
				if err = Render(path, target); err != nil {
					return err
				}
			} else if strings.Contains(path, ".md") {
				source := strings.Split(path, ".")[0]
				target := fmt.Sprintf("%s/%s.html", filepath.Dir(path), source)
				if err = Render(source, target); err != nil {
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

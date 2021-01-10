package main

import (
	"fmt"
	"log"
	"os"
)

//					 (manifestType) (defaults)
// TODO: Shift to using map[string][]string?
var (
	// The sub-types of a manifest(service) file
	manifestTypes = []string{
		"functions",
		"applications",
		"blogs",
		"docs",
		"fileserver",
		"games",
		"personal",
	}

	// A function is a stateless piece of software, that usually only does one
	// or two things
	functionsDefault = []string{
		"sbh 9001 https://gitlab.com/oglinuk/sbh.git",
		"quotitioner 9429 https://gitlab.com/oglinuk/quotitioner.git",
	}

	// An application is a piece of software that requires state
	applicationsDefault = []string{
		"gitea 3000 https://github.com/go-gitea/gitea.git", // TODO: Fix (to 9090) when gendockercompose is implemented
		"www-archiver 11111 https://gitlab.com/oglinuk/www-archiver.git",
		"mediafs 12121 https://gitlab.com/oglinuk/mediafs.git",
	}

	// A blog is a git hosted collection of code/.md/.html/site/... files around
	// discussion of a topic
	// TODO: Split blogs into individualsBlog/projectsBlog?
	blogsDefault = []string{
		"fragglet-blog 50006 https://github.com/fragglet/soulsphere.org.git",
		"rwxrob-blog 50007 https://github.com/rwxrob/rwxrob.git",
		"jessfraz-blog 50008 https://github.com/jessfraz/blog.git",
		"gitea-blog 50300 https://gitea.com/gitea/blog.git",
	}

	// A doc source is a git hosted collection of code/.md/.html/.txt/... files
	// documenting or explaining something
	docsDefault = []string{
		"rwx.gg 50500 https://gitlab.com/rwx.gg/README.git",
		"freedoom-docs 50502 https://github.com/freedoom/freedoom.github.io.git",
		"11238-docs 65535 https://gitlab.com/oglinuk/11238.git",
	}

	// A fileserver source is a git hosted repo that does not necessarily fall
	// into the docs or blogs categories, but still requires the scaffolding of
	// a fileserver
	fileserverDefault = []string{
		"directories 50000 https://gitlab.com/oglinuk/directories.git",
		"library 50001 https://gitlab.com/oglinuk/library.git",
		"ptp 50002 https://github.com/oglinuk/ptp.git",
		"awesome-gitea 50700 https://gitea.com/gitea/awesome-gitea.git",
	}

	// A game source is a git hosted game which is capable of being run locally
	// TODO: Dockerize?
	gamesDefault = []string{
		"chocolate-doom 30303 https://github.com/chocolate-doom/chocolate-doom.git",
		"quakespasm 30304 https://github.com/ericwa/quakespasm.git",
	}

	// A personal source is ones creation(s) that falls under any type of service
	personalDefault = []string{
		//"fourohfournotfound 40404 https://gitlab.com/oglinuk/fourohfournotfound.git"
	}
)

// createManifestFile called name with defaults for the content
func createManifestFile(name string, defaults []string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", MANIFESTSDIR, name))
	if err != nil {
		return err
	}
	defer f.Close()

	for _, d := range defaults {
		f.WriteString(fmt.Sprintf("%s\n", d))
	}
	return nil
}

// generateManifest with defaults of manifestType
func generateManifest(manifestType string) error {
	switch manifestType {
	case "functions":
		return createManifestFile(manifestType, functionsDefault)
	case "applications":
		return createManifestFile(manifestType, applicationsDefault)
	case "blogs":
		return createManifestFile(manifestType, blogsDefault)
	case "docs":
		return createManifestFile(manifestType, docsDefault)
	case "fileserver":
		return createManifestFile(manifestType, fileserverDefault)
	case "games":
		return createManifestFile(manifestType, gamesDefault)
	case "personal":
		return createManifestFile(manifestType, personalDefault)
	default:
		return fmt.Errorf("generateManifest::ERROR: Unknown type (%s)", manifestType)
	}
}

// createDefaultManifests for each value in manifestTypes
func createDefaultManifests() error {
	for _, t := range manifestTypes {
		if err = generateManifest(t); err != nil {
			log.Printf("generateManifest(%s)::ERROR: %s", t, err.Error())
			return err
		}
	}

	return nil
}

// createManifestsDir if not exists; then createDefaultManifests
func createManifestsDir() error {
	if _, err = os.Stat(MANIFESTSDIR); err != nil {
		log.Println("Creating manifests dir and creating default manifests ...")
		if err = os.MkdirAll(MANIFESTSDIR, 0744); err != nil {
			return err
		}

		if err = createDefaultManifests(); err != nil {
			return err
		}
	}

	return nil
}

// GenerateManifests directory && create default manifest files
func GenerateManifests() {
	if _, err = os.Stat(MANIFESTSDIR); err != nil {
		log.Println("manifests dir not found, generating manifests ...")

		if err = createManifestsDir(); err != nil {
			log.Fatalf("createManifestDir::ERROR: %s", err.Error())
		}
	}
}

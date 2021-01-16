package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jasonlvhit/gocron"
)

var (
	dontpanicInterval = flag.Uint64("dpi", 15, "Interval in minutes to execute DONTPANIC (Default is 15)")
	scanInterval      = flag.Uint64("si", 1, "Interval in minutes to execute ScanLocalhost (Default is 1)")
	env               = flag.String("e", "default", "Environment to generate")

	// ROOTDIR is a directory containing the manifests and services directories used by dont-panic-11238
	ROOTDIR = "DONTPANIC"
	// MANIFESTDIR is a directory containing the manifest files
	MANIFESTSDIR = fmt.Sprintf("%s/%s", ROOTDIR, "manifests")
	// SERVICEDIR is a directory containing the git repositories
	SERVICESDIR          = fmt.Sprintf("%s/%s", ROOTDIR, "services")
	err                  error
	wg                   = &sync.WaitGroup{}
	timeTaken            time.Duration
	timeSince            time.Time
	activeLocalhostPorts map[int]struct{}
)

// DONTPANIC handles the manifests and services
func DONTPANIC() {
	log.Println("=== DONTPANIC ===")
	if _, err = os.Stat(ROOTDIR); err != nil {
		if err = os.MkdirAll(ROOTDIR, 0744); err != nil {
			log.Fatalf("os.MkdirAll(%s)::ERROR: %s", ROOTDIR, err.Error())
		}
	}

	GenerateManifests(*env)
	GenerateServices()
	GenerateDockerCompose()
}

func init() {
	// TODO: improve testing and figure out way to not need to do this
	// "Fix" for flag provided but not defined: -test.timeout
	// https://github.com/golang/go/issues/31859
	testing.Init()
}

func main() {
	flag.Parse()

	activeLocalhostPorts = make(map[int]struct{})

	f, err := os.OpenFile("main.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Could not open log file ...")
	}

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	go DONTPANIC()
	go func() {
		gocron.Every(*dontpanicInterval).Minutes().Do(DONTPANIC)
		<-gocron.Start()
	}()

	go ScanLocalhost()
	go func() {
		gocron.Every(*scanInterval).Minutes().Do(ScanLocalhost)
	}()

	PORT := 11238
	HOST := "0.0.0.0"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Refactor ...
		var portLinks []string
		for port := range activeLocalhostPorts {
			portLinks = append(portLinks,
				fmt.Sprintf("<a href=\"http://localhost:%d\">%d</a>", port, port))
		}
		sort.Strings(portLinks)

		// TODO: Refactor ...
		var serviceLinks []string
		for _, v := range DefaultEnvs[*env] {
			for _, entry := range v {
				parts := strings.Split(entry, " ")
				serviceName := fmt.Sprintf("%s-%s", parts[0], parts[1])
				serviceLinks = append(serviceLinks,
					fmt.Sprintf("<br><a href=\"http://localhost:%d\">%s</a>", parts[1], serviceName))
			}
		}
		sort.Strings(serviceLinks)

		// TODO: Refactor using a template
		fmt.Fprintf(w, `
			<h1>Hello %d!</h1><br><br>

			<p>
			This is your <b><i>localhost: DONT PANIC!</i></b> A bazaar of
			services all running on various ports locally! Below is a list
			of all active ports.
			</p><br><br>

			<p>
			This page took %s to scan.
			<br>
			Scan interval: %d (minutes) - use '-si <1-60>' to change the interval
			<br>
			Last scan was %s.
			</p><br><br>

			<h3>Current active ports:</h3>
			<b>Total Ports Active: %d</b><br><br>

			%v

			<br>
			<br>

			%v
		`, PORT, timeTaken, 1, time.Since(timeSince), len(activeLocalhostPorts), portLinks, serviceLinks)
	})

	log.Printf("Serving at localhost:%d ...", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), nil))
}

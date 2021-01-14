package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jasonlvhit/gocron"
)

var (
	interval = flag.Uint64("i", 15, "Interval to scan and check manifests/services (Default is 15)")
	env      = flag.String("e", "default", "Environment to generate")

	// ROOTDIR is a directory containing the manifests and services directories used by dont-panic-11238
	ROOTDIR = "DONTPANIC"
	// MANIFESTDIR is a directory containing the manifest files
	MANIFESTSDIR = fmt.Sprintf("%s/%s", ROOTDIR, "manifests")
	// SERVICEDIR is a directory containing the git repositories
	SERVICESDIR          = fmt.Sprintf("%s/%s", ROOTDIR, "services")
	err                  error
	wg                   = &sync.WaitGroup{}
	activeLocalhostPorts []string
	timeTaken            time.Duration
	timeSince            time.Time
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

	// TODO: add manifest of standard/reserved service ports to check initially
	// as a heartbeat analytics measure
	sTime := time.Now()

	// TODO: Improve ...
	activeLocalhostPorts = ScanLocalhost()

	timeTaken = time.Since(sTime)
	timeSince = time.Now()
	log.Printf("ScanLocalhost took %s ...", timeTaken)
}

func init() {
	flag.Parse()

	f, err := os.OpenFile("main.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Could not open log file ...")
	}

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	go DONTPANIC()
	go func() {
		gocron.Every(*interval).Minutes().Do(DONTPANIC)
		<-gocron.Start()
	}()
}

func main() {
	PORT := 11238
	HOST := "0.0.0.0"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Refactor ...
		var portLinks []string
		for _, port := range activeLocalhostPorts {
			portLinks = append(portLinks,
				fmt.Sprintf("<a href=\"http://%s\">%s</a>", port, strings.Split(port, ":")[1]))
		}

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
			Scan interval: %d (minutes) - use '-i <1-60>' to change the interval
			<br>
			Last scan was %s.
			</p><br><br>

			<h3>Current active ports:</h3>
			<b>Total Ports Active: %d</b><br><br>

			%v
		`, PORT, timeTaken, *interval, time.Since(timeSince), len(activeLocalhostPorts), portLinks)
	})

	log.Printf("Serving at localhost:%d ...", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), nil))
}

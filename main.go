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
	"time"

	"github.com/jasonlvhit/gocron"
	"gitlab.com/OGLinuk/dont-panic-11238/dontpanic"
)

var (
	interval    = flag.Uint64("i", 15, "Interval in minutes to execute (Default is 15)")
	env         = flag.String("e", "default", "Environment to generate")
	err         error
	timeTaken   time.Duration
	timeSince   time.Time
	activePorts []string
)

// DONTPANIC generates manifests (of *env), services, and docker-compose.yml
func DONTPANIC() {
	log.Println("=== DONTPANIC ===")
	dontpanic.GenerateManifests(*env)
	dontpanic.GenerateServices()
	dontpanic.GenerateDockerCompose()
}

func main() {
	flag.Parse()

	f, err := os.OpenFile("main.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Could not open log file ...")
	}

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	go func() {
		DONTPANIC()
		gocron.Every(*interval).Minutes().Do(DONTPANIC)
	}()

	go func() {
		activePorts = dontpanic.ScanLocalhost()
		gocron.Every(1).Minutes().Do(func() {
			sTime := time.Now()
			activePorts = dontpanic.ScanLocalhost()
			timeSince = time.Now()
			timeTaken = time.Since(sTime)
		})
		<-gocron.Start()
	}()

	PORT := 11238
	HOST := "0.0.0.0"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Refactor ...
		var portLinks []string
		for _, addr := range activePorts {
			portLinks = append(portLinks,
				fmt.Sprintf("<a href=\"http://%s\">%s</a>", addr, strings.Split(addr, ":")[1]))
		}
		sort.Strings(portLinks)

		// TODO: Refactor ...
		var serviceLinks []string
		for k, v := range dontpanic.DefaultEnvs[*env] {
			// TODO: Filter better; no reason to show services like games if theyre not web based
			if k != "games" {
				for _, entry := range v {
					parts := strings.Split(entry, " ")
					serviceName := fmt.Sprintf("%s-%s", parts[0], parts[1])
					serviceLinks = append(serviceLinks,
						fmt.Sprintf("<br><a href=\"http://localhost:%s\">%s</a>", parts[1], serviceName))
				}
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
			Scan interval: %d (minutes) - use '-i <1-60>' to change the interval
			<br>
			Last scan was %s.
			</p><br><br>

			<h3>Current active ports:</h3>
			<b>Total Ports Active: %d</b><br><br>

			%v

			<br>
			<br>

			%v
		`, PORT, timeTaken, 1, time.Since(timeSince), len(activePorts), portLinks, serviceLinks)
	})

	log.Printf("Serving at localhost:%d ...", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", HOST, PORT), nil))
}

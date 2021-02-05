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
	"text/template"
	"time"

	"github.com/jasonlvhit/gocron"
	dontpanic "gitlab.com/OGLinuk/dont-panic-11238"
)

var (
	interval = flag.Uint64("i", 15, "Interval in minutes to execute (Default is 15)")
	env      = flag.String("e", "default", "Environment to generate")

	host = "0.0.0.0"
	port = 11238

	err         error
	timeTaken   time.Duration
	timeSince   time.Time
	activePorts map[string]struct{}

	tpl *template.Template
)

// dontpanicResponse is the data used to pass to the index.html template
type dontpanicResponse struct {
	Title             string
	Description       string
	TimeComplexity    time.Duration // timeTaken
	LastExecution     time.Duration // time.Since(timeSince)
	NumActiveServices int           // len(activePorts)
	ServiceLinks      []string
}

// DONTPANIC generates manifests (of *env), services, and docker-compose.yml
func DONTPANIC() {
	log.Println("=== DONTPANIC ===")
	dontpanic.GenerateManifests(*env)
	dontpanic.GenerateServices()
	dontpanic.GenerateDockerCompose()
}

func main() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	flag.Parse()

	/* ----- Logging ----- */
	f, err := os.OpenFile("main.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Could not open log file ...")
	}
	defer f.Close()

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	/* ----- Start DONTPANIC && gocron ----- */
	go func() {
		DONTPANIC()
		gocron.Every(*interval).Minutes().Do(DONTPANIC)
	}()

	go func() {
		sTime := time.Now()
		activePorts = dontpanic.ScanLocalhost()
		timeSince = time.Now()
		timeTaken = time.Since(sTime)
		gocron.Every(1).Minutes().Do(func() {
			sTime = time.Now()
			activePorts = dontpanic.ScanLocalhost()
			timeSince = time.Now()
			timeTaken = time.Since(sTime)
		})
		log.Printf("Serving at localhost:%d ...", port)
		<-gocron.Start()
	}()

	/* ----- Index HandlerFunc ----- */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var serviceLinks []string
		for k := range activePorts {
			servicePort := strings.Split(k, ":")[1]

			// TODO: Refactor and also group the links by their defined ek
			for ek, ev := range dontpanic.DefaultEnvs[*env] {
				if ek != "games" {
					for _, sv := range ev {
						parts := strings.Split(sv, " ")
						name := parts[0]
						port := parts[1]
						if servicePort == port {
							serviceLinks = append(serviceLinks,
								fmt.Sprintf("<a href=\"http://%s\">%s</a>", k, name))
						}
					}
				}
			}
		}
		sort.Strings(serviceLinks)

		if err = tpl.ExecuteTemplate(w, "index.html", &dontpanicResponse{
			Title: "11238",
			Description: `
			This is your <b><i>localhost:11238 - DONT PANIC!</i></b> A bazaar of
			services all running on various ports <b><i>locally</i></b> and
			<b><i>offline</i></b>! Below is a list of all active ports.`,
			TimeComplexity:    timeTaken,
			LastExecution:     time.Since(timeSince),
			NumActiveServices: len(serviceLinks),
			ServiceLinks:      serviceLinks,
		}); err != nil {
			log.Fatalf("tpl.ExecuteTemplate: %s", err.Error())
		}
	})

	/* ----- Starting dont-panic-11238 ----- */
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}

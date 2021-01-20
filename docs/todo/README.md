# TODO

## For TODO

* [ ] Organize better
* [ ] Create a way to customize granularity (category, priority ...)

---

## For dont-panic-11238 in general

* [X] Create a framework to scaffold the minimal code necessary to serve a function, application, blog, ...
    * [X] Fileserver
* [X] Create way to configure scaffold with a git repository. (Ex: manifest/blogs)
* [X] Define the PORT layout for services the service sub-types (still WiP but draft is done)
* [X] Define environment manifest types (still WiP but draft is done)
* [ ] Pseudo-randomly assign a port to a service within the services rather than assiging a specific port?
* [ ] Write tests
    * [ ] `manifests_test.go`
    * [ ] `services_test.go`
    * [ ] `fileserver_test.go`
    * [X] `dockerfile_test.go`
    * [X] `scan_test.go`
    * [ ] `utils_test.go`
    * [X] `dockercompose.go`
* [X] Implement [cmdtab](https://github.com/rwxrob/cmdtab)
* [ ] Automate way to change PORTs of existing projects like gitea
* [ ] GoDoc
* [ ] Review

---

## For `manifests.go`

* [X] Create
* [ ] Write test
* [X] Refactor `GenerateManifests` and abstract the default environment generated
* [X] Create localhost environments (default universes focused on a specific topic)
    * [X] Default
    * [ ] Data science
    * [ ] Livestreaming
    * [ ] Pentesting
    * [ ] Web development
    * [ ] Writing (blogging, notes, ...)

---

## For `services.go`

* [X] Create
* [ ] Write test
* [X] Implement `git pull` functionality in genService() if the service exists locally
	* [ ] Replace current implementation (using runCmd) with [go-git](https://github.com/go-git/go-git)
* [X] Change serviceType to `fn func(name, port, link string)` and allow a func to be passed to generate any necessary the servicePath repository
* [ ] Use `readme.go` to generate a default README.md file in manifestPath dir
* [ ] Remove service when its removed from the manifest file

---

## For `scan.go`

* [X] Create
* [ ] Write test

---

## For `dockercompose.go`

* [X] Create
* [X] Write test
* [X] Generate `docker-compose.yml` file based on manifest file entries
* [ ] Add functionality to append a service
* [ ] Add functionality to remove a service

---

## For `fileserver.go`

* [X] Create
* [X] Write test

---

## For `docker.go`

* [X] Create
* [X] Write test

---

## For `doc.go`

* [X] Create
* [ ] Write test
* [X] Replace `render.sh` with `doc.go` to execute [pandoc](https://github.com/jgm/pandoc)
* [ ] Create Render(src, target string) ->`pandoc $src -o $target`
* [ ] Create RenderWithTemplate(src, target, template string) -> `pandoc $src -o $target --template=$template`
* [ ] Replace pandoc with a pegn parser

---

## For `readme.go`

* [ ] Create
* [ ] Write test

---

## For `git.go`

* [ ] Create
* [ ] Write test

---

## For `docs` and root README.md

* [X] Use `doc.go` to generate HTML files

---

## For `docs/ports`

* [ ] Review layout and improve

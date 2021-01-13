---
title: TODO
---

## For TODO

* [ ] Organize better
* [ ] Create a way to customize granularity (category, priority ...)

## For `docs` and root README.md

* [ ] Use `pandoc.go` to generate HTML files

## For `docs/ports`

* [ ] Review layout and improve

## For `services.go`

* [ ] Change serviceType to `fn func(name, port, link string)` and allow a func to be passed to scaffold the servicePath repository
    * Ex: 
* [ ]  Use `readme.go` to generate a default README.md file in manifestPath dir

## For dont-panic-11238 in general

* [X] Create a framework to scaffold the minimal code necessary to serve a function, application, blog, ...
    * [X] Fileserver
* [X] Create way to configure scaffold with a git repository. (Ex: manifest/blogs)
* [X] Define the PORT layout for services the service sub-types (still WiP but draft is done)
* [X] Define environment manifest types (still WiP but draft is done)
* [X] Implement `git pull` functionality in services.go -> genService() if service exists
* [X] Generate `docker-compose.yml` file based on manifest file entries
    * [ ] Add functionality to append a service
    * [ ] Add functionality to remove a service
        * [ ] Remove service when its removed from the manifest file
* [X] Refactor `GenerateManifests` and abstract the default environment generated
* [ ] Create `readme.go` 
* [ ] Replace `render.sh` with `pandoc.go` to execute [pandoc](https://github.com/jgm/pandoc)
    * [ ] Render(src, target string) ->`pandoc $src -o $target`
    * [ ] RenderWithTemplate(src, target, template string) -> `pandoc $src -o $target --template=$template`
* [X] Create localhost environments (default universes focused on a specific topic)
    * [X] Default
    * [ ] Data science
    * [ ] Livestreaming
    * [ ] Pentesting
    * [ ] Web development
    * [ ] Writing (blogging, notes, ...)
* [ ] Write tests
    * [ ] `manifests_test.go`
    * [ ] `services_test.go`
    * [ ] `fileserver_test.go`
    * [ ] `dockerfile_test.go`
    * [ ] `scan_test.go`
    * [ ] `utils_test.go`
    * [ ] `dockercompose.go`
* [ ] Replace `Makefile` with [cmdtab](https://github.com/rwxrob/cmdtab)
* [ ] Automate way to change PORTs of existing projects like gitea
* [ ] GoDoc
* [ ] Review

---
title: dont-panic:11238
---

## A guide to the universe of ***your*** **localhost**.

Yes the name is a spin off of [The Hitch Hikers Guide to the Galaxy](). No I have not finished the trilogy of four yet, so please no spoilers.

## Why?

The hope behind dont-panic-11238 is to give people the ability to replicate the spaces or software of the internet that they observe/use. There is plenty of good content (cool tech, fun games, awesome blogs, good documentation, and a whole lot more) out their that is hosted via git, and because of that dont-panic-11238 enables you to have an up-to-date local copy of those repos. The idea then grows into your own universe or rather version of the internet that you have access to anytime, online/offline, and on any device (provided the device is capable of running dont-panic-11238 and all defined services of dont-panic-11238).

## Getting Started

`make` && `lynx localhost:11238`

or

`localhost:11238` on any browser

## Default Services (18 total)

* [X] dont-panic-11238
* [X] sbh (stateless password panager)
* [X] quotitioner (curlable and REST available quotes)
* [X] gitea (git web UI)
* [X] awesome-gitea (list of gitea projects)
* [X] www-archiver (http/s website archiver)
* [X] mediafs (videos, audio, images, gifs, ...)
* [X] library (html, pdfs, whitepapers, ...)
* [X] directories (lists of lists: individuals, projects, ...)
* [X] chocolate-doom (cause DOOM)
* [X] fragglet-blog (creator/maintainer of chocolate-doom/freedoom)
* [X] freedoom-docs (website for freedoom)
* [X] rwx.gg (beginner boosts)
* [X] rwxrob (creator/maintainer of readme.world/rwx.gg/pegn.dev)
* [X] jessfraz (crazy smart person who runs *everything* in docker - look up her rant on open source firmware)
* [X] ptp (code snippets)
* [X] quakespasm (cause Quake)

## Standard Services Port Layout

It is important to note that the already known used TCP and UDP ports are
omitted from these ranges and are assumed not to be taken if not absolutely
necessary.

**< 9000 not used**

* 9000-9999 - functions (used by 11238)
    * 9001 - sbh
    * 9429 - quotitioner
    * 9090 - gitea (currently 3000 (TODO))

* 10000-19999 - applications (like 11238)
    * 11238 (DONT PANIC!)
    * 11111 www-archiver

* 20000-29999 - personal
    * ...

* 30000-39999 - games
    * 30303 - doom (chocolate-doom)
    * 30304 - doom2 (crispy-doom)
    * 31313 - quake (quakespasm)
    * 31314 - quake2 (yamagi?)
    * 31315 - quake3 (quake3io)
    * 32323 - chess (smallchesslib)
    * 34343 - amongst-us (amongst-us)
    
* 40000-44999 - blogs
    * 40006 - fragglet-blog
    * 40007 - rwxrob-blog
    * 40008 - jessfraz-blog
    * 43000 - gitea-blog

* 45000-49999 - docs
    * 45000 - rwx.gg
	* 46000 - freedoom-docs
    * 46001 - linode-docs

* 50000-59999 - fileservers
    * 50000 - directories
    * 50001 - library
    * 50002 - ptp
    * 50700 - awesome-gitea

* 60000-65535 - TODO: ...

## TODO:

* [X] Create a framework to scaffold the minimal code necessary to serve a function, application, blog, ...
* [X] Create way to configure scaffold with a git repository. (Ex: manifest/blogs)
* [X] Define the PORT layout for services the service sub-types (Still WiP but draft is done)
* [X] Write tests
    * [ ] `manifests_test.go`
    * [ ] `services_test.go`
    * [ ] `fileserver_test.go`
    * [ ] `dockerfile_test.go`
    * [ ] `scan_test.go`
* [ ] Automate way to change PORTs of existing projects like gitea
* [ ] 
* [ ] GoDoc
* [ ] Review
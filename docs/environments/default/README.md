---
title: Default Environment
---

The default environment setup for a basic localhost universe. It includes a variety of services from the bias perspective of a thinker, a programmer, a gamer, a reader, and a questioner.

## Services
* functions
    * [X] sbh (stateless password panager)
    * [X] quotitioner (curlable and REST available quotes)

* applications
    * [X] dont-panic-11238
    * [X] gitea (git web UI)
    * [X] www-archiver (http/s website archiver)

* individuals-blog
    * [X] fragglet-blog (creator/maintainer of chocolate-doom/freedoom)
    * [X] rwxrob (creator/maintainer of readme.world/rwx.gg/pegn.dev)
    * [X] jessfraz (crazy smart person who runs *everything* in docker - look up her rant on open source firmware)

* projects-blog

* docs
    * [X] freedoom-docs (website for freedoom)
    * [X] rwx.gg (beginner boosts)

* fileserver
    * [X] mediafs (videos, audio, images, gifs, ...)
    * [X] library (html, pdfs, whitepapers, ...)
    * [X] directories (lists of lists: individuals, projects, ...)
    * [X] ptp (code snippets)
    * [X] awesome-gitea (list of gitea projects)

* games
    * [X] chocolate-doom (cause DOOM)
    * [X] quakespasm (cause Quake)


## Port Layout

* 9000-9999 - functions
    * 9001 - sbh
    * 9429 - quotitioner
    * 9999 - gitea (currently 3000 (TODO))

* 10000-19999 - applications
    * 11238 (DONT PANIC!)
    * 11111 www-archiver

* 20000-29999 - personal

* 30000-39999 - games
    * 30303 - doom (chocolate-doom)
    * 30304 - doom2 (chocolate-doom)
    * 31313 - quake (fteqw)
    * 31314 - quake2 (fteqw)
    * 31315 - quake3 (fteqw)
    * 32323 - chess (smallchesslib)
    * 34343 - amongst-us (amongst-us)

* 40000-43999 - (individuals) blogs
    * 40006 - fragglet-blog
    * 40007 - rwxrob-blog
    * 40008 - jessfraz-blog

* 43999-44999 - (projects) blogs
    * 43001 - go-blog
    * 43001 - gitea-blog

* 45000-49999 - docs
    * 45000 - rwx.gg
	* 46000 - freedoom-docs
    * 46001 - linode-docs
    * 49999 - godocs

* 50000-59999 - fileservers
    * 50000 - directories
    * 50001 - library
    * 50002 - ptp
    * 50700 - awesome-gitea

* 60000-65535 - TODO: ...

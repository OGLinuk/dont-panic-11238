# dont-panic:11238

## A guide to the universe of ***your*** **localhost(s)**.

Yes the name is a spin off of [The Hitch Hikers Guide to the Galaxy](https://en.wikipedia.org/wiki/The_Hitchhiker's_Guide_to_the_Galaxy). No I have not finished the trilogy of four yet, so please no spoilers.

## Why?

The hope behind dont-panic-11238 is to give people the ability to replicate the spaces or software of the internet that they observe/use. There is plenty of good content (cool tech, fun games, awesome blogs, good documentation, and a whole lot more) out their that is hosted via git, and because of that dont-panic-11238 enables you to have an up-to-date local copy of those repos. The idea then grows into your own universe or rather version of the internet that you have access to anytime, online/offline, and on any device (provided the device is capable of running dont-panic-11238 and all defined services of dont-panic-11238).

## **Pre-requisites**

* `>= go 1.15`
* `>= docker 19.03.8`
* `>= docker-compose 1.25.5`
* `>= git 2.25.1`

## ***Getting Started***

`make`

once all of the services have been downloaded then run

`make services`

finally goto [`localhost:11238`](http://localhost:11238) on any browser (lynx, firefox, ...)

***Note:*** there are some services (like quakespasm) that arent web based and
will have to be operated from its source.

## **Documentation**

[Available here](docs)

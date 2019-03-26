# Little Miss Bossy

A small health check server written in [Go](https://golang.org)

![](little-miss-bossy.jpg)

> _Image probably (c) by Roger Hargreaves_
>
> There is no particular link between this program its namesake. This book was
just next in the list

## What does it do?

* Runs a web server on port 65432
* Samples /proc/stat to see what's going on with your CPU
* Serves up a /health endpoint that returns 200 if CPU utilisation is 75% or
  below, and 503 if CPU utilisation is over 75%

## Why does this help?

Your load balancer can check this endpoint to see if the machine is overloaded.

## Future enhancements

* Configurability
* Cross-platform

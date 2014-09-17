quick-serve
===========

Quickly serve static websites

Install
=======

If you have Go installed, you can run `go install github.com/pguelpa/quick-serve`

If you don't have Go installed, you can use [Gobuild.IO](http://gobuild.io/download/github.com/pguelpa/quick-serve)

Usage
=====

To serve the static files in the current directory on port 8080:

`quick-serve`

Use the `PORT` and `ROOT` environment variables to override the defaults.

`PORT=8888 ROOT=$HOME/website quick-serve`

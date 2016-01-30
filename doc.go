/*
Volatile CLI is a tool for managing Volatile app source code.

Installation

	go get github.com/volatile/volatile

NOTE: The Volatile CLI is actually only compatible with Unix-based systems.

New app

	volatile new myapp
makes a new directory named "myapp" and puts bootstrap code in.

If your app will be an API, use the api argument:
	volatile new api myapp

Run app

When you are inside a Volatile app directory, just use
	volatile
to automatically recompile and rerun the app every time a file change.

A ./build file can be used to automatically execute a shell script before building and running your app.
Don't forget the shebang and the executability rights for this script!

Update Volatile

To get updated versions of the core (https://godoc.org/github.com/volatile/core) and all handlers and helpers (https://godoc.org/github.com/volatile/core#hdr-Handlers_and_helpers):
	volatile update
*/
package main

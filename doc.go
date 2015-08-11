/*
Volatile is the perfect foundation for any web app as it's designed to have the best balance between readability, flexibility and performance.

It provides a pure handler (or "middleware") stack so you can perform actions downstream, then filter and manipulate the response upstream.

For a complete documentation, see the [Volatile website](http://volatile.whitedevops.com).

Installation

In the terminal:

	$ go get github.com/volatile/volatile

Create a new app

In the terminal:

	$ volatile new myapp

… creates a new directory named "myapp", puts bootstrap code in, and runs it.

If your future app will be an API, use the "api" argument:

	$ volatile new api myapp

Run an app

In the terminal:

	$ volatile run

… recompiles and reruns the app in the current directory every time a file change.

Update Volatile

In the terminal:

	$ volatile update

… gets updated versions of the [Volatile Core](https://github.com/volatile/core) and all [official packages](https://github.com/volatile/core#official-handlers).
*/
package main

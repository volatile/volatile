/*
Volatile is the perfect foundation for any web app as it's designed to have the best balance between readability, flexibility and performance.

It provides a pure handler (or "middleware") stack so you can perform actions downstream, then filter and manipulate the response upstream.

For a complete documentation, see the Volatile website (http://volatile.whitedevops.com).

Getting started

1. Install Go

Before creating an app with Volatile, you must have a working Go installation.
Follow the official guide (https://golang.org/doc/install) (and don't forget to set your "GOPATH" (https://golang.org/doc/code.html) environment variable).

2. Install Volatile

Get the Volatile command line interface with the "go get" command:

	$ go get github.com/volatile/volatile

NOTICE: The Volatile command line interface is actually only compatible with Unix-based systems.

3. Create a new app

	$ volatile new myapp

… creates a new directory named "myapp" and puts bootstrap code in.

If your future app will be an API, use the "api" argument:

	$ volatile new api myapp

The app runs directly and is reachable at http://localhost:8080/.

4. In the near future

When you're inside a Volatile app directory, just use…

	$ volatile

… to automatically recompile and rerun the app every time a file change.
If a ./build file exists, it will be executed before building and running your app. Don't forget the shebang line on top of your script.

Finally, to get updated versions of the Core (https://github.com/volatile/core) and all official packages (http://volatile.whitedevops.com#handlers-and-helpers):

	$ volatile update
*/
package main

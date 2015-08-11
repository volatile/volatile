<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/volatile/logo.png" alt="Volatile" title="Volatile"><br><br></p>

Volatile is the perfect foundation for any web app as it's designed to have the best balance between **readability**, **flexibility** and **performance**.  

It provides a pure handler (or *middleware*) stack so you can perform actions downstream, then filter and manipulate the response upstream.

For a complete **documentation**, see the [Volatile website](http://volatile.whitedevops.com).

## Installation

```Shell
$ go get github.com/volatile/volatile
```

## Usage

### Create a new app

```Shell
$ volatile new myapp
```
… creates a new directory named "myapp" and puts bootstrap code in.

If your future app will be an API, use the `api` argument:

```Shell
$ volatile new api myapp
```

The app is run directly and is reachable at [localhost:8080](http://localhost:8080/).

### Run an app

```Shell
$ volatile run
```
… recompiles and reruns the app in the current directory every time a file change.

### Update Volatile

```Shell
$ volatile update
```
… gets updated versions of the [Volatile Core](https://github.com/volatile/core) and all [official packages](https://github.com/volatile/core#official-handlers).

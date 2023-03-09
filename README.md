# Go Tree Demo

## Content

* [Requirements](#requirements)
* [Building executable](#building-executable)
* [Usage](#usage)
* [Dev Notes](#dev-notes)
  + [Styling and Linting](#styling-and-linting)
* [Known issues](#known-issues)
  + [1. GoLand IDE complains about outdated Delve version for debugging.](#1-goland-ide-complains-about-outdated-delve-version-for-debugging)
* [Possible Future Improvements](#possible-future-improvements)

## Requirements

* Go version: 
  * v1.19
  * v1.20 (preferable)

## Building executable

* Make sure that `$GOROOT` and `$GOPATH` environmental variable paths are set. 
* Go to project directory.
* Run `go build ./...` to build application.
* Run `go install ./...`

## Usage

* Go to `$GOPATH/bin` folder.
* Run `./tree -h` command to read documentation.
* Run `./tree "path/to/dir"` to build tree representation of the chosen directory with default depth (= 5).
* Run `./tree -d 2 "path/to/dir"` to build tree representation of the chose directory with depth = 2.

## Dev Notes

### Styling and Linting 

* **GoFmt Style**
  * Project uses `gofmt` check in GH Actions to verify that all contributors are following the same format style.
  * For `GoLand` you can set `gofmt` File Watcher to apply specified formatting automatically.
* **Go Vet & StaticCheck**
  * Project uses mentioned above tools in GH Actions to verify that all contributors follow best GoLang practices.

## Known issues

### 1. GoLand IDE complains about outdated Delve version for debugging.

Error message:
```
undefined behavior - version of Delve is too old for Go version 1.20.0 (maximum supported version 1.19)
```

Solution:
* Install `dlv` binary with `go install github.com/go-delve/delve/cmd/dlv@latest`
* Set the `dlv.path=<path_to_dlv_executable>` under `Help > Edit Custom Properties`
* Restart GoLand

## Possible Future Improvements

* Improve the following:
  * Better (more readable) ASCII representation of the directory tree.
  * Better abstraction for Printer model to support different formatting.
* Add the following features:
  * Show full path prefix (if specified flag is set).
  * Show only directories (if specified flag is set).
  * Show file sizes (if specified flag is set).
  * Show permissions and ownership (if specified flag is set).
* Add test coverage check into GH Actions.
* Add automated release creation via GH Actions when something is merged to master branch.
* Add Docker image for containerization (?) -- as it's cmd tool now I don't see any benefits of doing so at the moment.

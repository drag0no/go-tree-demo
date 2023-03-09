# Go Tree Demo

## Requirements

* Go v1.20

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

### Styling & Linting 

* **GoFmt Style**
  * We use `gofmt` check in GH Actions to verify that all contributors are following the same format style.
  * For `GoLand` you can set `gofmt` File Watcher to apply specified formatting automatically.
* **Go Vet, GoLint & Static Check**
  * We use mentioned above tools in GH Actions to verify that all contributors are following best GoLang practices.

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

## TODO

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

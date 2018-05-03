# git (shallow)

Want to make sure `go get ...` always does a shallow git pull/clone? This is a little hack that will do it. Basically its just a wrapper for git, and if "pull/clone" comes up it makes sure to make it shallow. To make it work you have to install this as `$GOPATH/bin/git` and then *prepend* that to the PATH so that the wrapper will call the hardcoded git path.

```
$ go get github.com/schollz/git
$ export PATH=$GOPATH/bin:$PATH
```

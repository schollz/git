# How to use shallow clones for `go get`

This is a little hack to use shallow clones for new git checkouts with `go get`. Unfortunately for Gophers, [this has been an open issue for three years counting](https://github.com/golang/go/issues/13078) without a workable solution aside from patching the go toolchain yourself. This solution utilizes a *git* wrapper that determines if a pull/clone is happening and then makes sure it is shallow. 

To install it you do 

```
$ go get github.com/schollz/git
```

The *git*-wrapper tool is named *"git"* on purpose so that your GOPATH can be prepended to the PATH and then the *git*-wrapper substituted for the real *git* (`/usr/bin/git`). So then, to activate shallow cloning all you have to do is: 

```
$ export PATH=$GOPATH/bin:$PATH
```

which you can add to your `.bashrc` files if you want it to be permanent. This way, the wrapper will aways be used and the wrapper will force cloning to be shallow.

# Benchmarks

Here's a benchmark showing a 50% reduction in disk usage and thus a 50% reduction in time taken for a `go get`. You'll not get that much for smaller repositories, but its not bad.

## normal `go get`


```
% docker run -it golang:1.10 /bin/bash
root@d9208178f1fa:/go# time go get github.com/juju/juju/...
real    7m35.631s
user    1m40.059s
sys     0m45.436s
root@d9208178f1fa:/go# du -sh .
1.1G
```

## shallow `go get`

```
% docker run -it golang:1.10 /bin/bash
root@68135fb64a3e:/go# go get github.com/schollz/git
root@68135fb64a3e:/go# export PATH=$GOPATH/bin:$PATH
root@68135fb64a3e:/go# time go get github.com/juju/juju/...
real    3m0.335s
user    0m29.192s
sys     0m17.253s
root@d9208178f1fa:/go# du -sh .
499M    .
```

# Acknowledgements

Thanks [tscholl2](https://github.com/tscholl2) for the idea.

# License

MIT
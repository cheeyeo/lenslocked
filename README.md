### Web Development with Go

Building a web app in go from scratch


### Reload

Use `modd` module and `modd.conf` file to watch for code changes:

```
go install github.com/cortesi/modd/cmd/modd@latest
```

Need to check that GOPATH is set and added to PATH first:
```
export GOPATH=$HOME/go
export PATH=$GOPATH:PATH

modd --version

modd
```


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

* `http.HandlerFunc` is itself a type of func:
```
type HandlerFunc func(ResponseWriter, *Request)
```

its able to declare a ServeHTTP function where it calls itself:
```
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

so we can use our own `pathHandler` function and pass it into http.ListenAndServe() via `http.HandleFunc(pathHandler))

In summary:
1. The HandlerFunc type implements the Handler interface

2. We can convert our handler functions into HandlerFuncs

3. When a function is converted into a HandlerFunc, it has a ServeHTTP method which implements the Handler interface.


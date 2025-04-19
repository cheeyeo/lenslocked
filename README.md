### Web Development with Go

Building a web app in go from scratch

### Running postgresql container
```
docker compose -f cmd/exp/compose.yml up
```

### Reload

Use `modd` module and `modd.conf` file to watch for code changes:

```
go install github.com/cortesi/modd/cmd/modd@latest
```

Need to check that GOPATH is set and added to PATH first:
```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH

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


### On managing passwords

To handle passwords securely in web apps:

* Use HTTPs to secure domain

* Store hashed passwords. Never store encrypted or plaintext passwords.

* Add salt to passwords before hashing

* Using time-constant functions during authentication.


2 hashing functions to focus on:

* bcrypt - for hashing passwords
* HMAC - to hash session tokens

Always use password specific hashing function when setting up an authentication system

While it might be tempting to use bcrypt like we did with our passwords for hashing session tokens, this approach isn’t a good fit for session tokens. When using bcrypt, every hash has a unique salt added to it. This means we wouldn’t be able to hash a session token and then search our database for it, as it would have a different salt value.
This helps prevent rainbow table attacks, but we only need to worry about those with passwords, not session tokens

### Encryption functions are not hashing functions

At some point you may learn about encryption functions like AES, and at first
glance they might appear to be a good fit. After all, the way we code AES looks
somewhat similar to HMAC. I want to once again reiterate that encryption and
hashing are two different things, and encryption is not appropriate for an securing
passwords.

When you encrypt a password it is reversible. Yes, you do need the encryption key
to decrypt a password, but if someone has hacked your server it isn’t a stretch to
imagine that they might be able to figure out your encryption key. Once they have
that key they will be able to decrypt every password in the database. This is not
possible with hashing functions, even ones that use a key.

It is my suggestion that you stick with bcrypt, or if you want to explore alterna-
tives scrypt and argon2 are also great choices, albeit a little harder to use with
Go


#### CSRF






### REF

[Logging in go]: https://www.bytesizego.com/blog/guide-to-logging-in-go

[Using CSRF middleware]: https://github.com/gorilla/csrf

[Using postgresql container]: https://www.docker.com/blog/how-to-use-the-postgres-docker-official-image/
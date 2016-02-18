# go-jsonrpc

> This is Golang Jsonrpc(v2) Package

#### Usage

* Install Dependencies
```sh
  $ godep restore
```

* Start Server
```sh
  $ go run demo/server.go
```

* Run Tests (Client)
```sh
  $ go test -v demo/server_test.go
```


> the graceful restart feature of server is based on [endless](https://github.com/fvbock/endless).
Check the [examples](https://github.com/fvbock/endless/tree/master/examples) for more information.

#### Signals

The endless server will listen for the following signals: `syscall.SIGHUP`, `syscall.SIGUSR1`, `syscall.SIGUSR2`, `syscall.SIGINT`, `syscall.SIGTERM`, and `syscall.SIGTSTP`:

`SIGHUP` will trigger a fork/restart

`syscall.SIGINT` and `syscall.SIGTERM` will trigger a shutdown of the server (it will finish running requests)

`SIGUSR2` will trigger [hammerTime](https://github.com/fvbock/endless#hammer-time)

`SIGUSR1` and `SIGTSTP` are listened for but do not trigger anything in the endless server itself. (probably useless - might get rid of those two)



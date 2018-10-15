# WASM Dom in go

![Web Assembly](https://cdn-images-1.medium.com/max/400/1*fb19K-mU-ipu0NRpc_p2pA.png)

# Requirements

- go >= 1.11
- nodejs

# Building

```sh
make
```

watch and build with [entr](http://www.entrproject.org/)

```sh
ls *.go | entr make
```

# Running

You probably have python

```sh
python -m SimpleHTTPServer 8080
```

If you want/have [goexec](https://github.com/shurcooL/goexec#goexec)

```sh
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

open [http://localhost:8080](http://localhost:8080)

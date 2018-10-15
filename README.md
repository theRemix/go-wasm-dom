![Web Assembly](https://cdn-images-1.medium.com/max/400/1*fb19K-mU-ipu0NRpc_p2pA.png)

### [Live demo : https://theremix.github.io/go-wasm-dom](https://theremix.github.io/go-wasm-dom/index.html)

## Requirements

- go >= 1.11
- nodejs

## Building

```sh
make
```

watch and build with [entr](http://www.entrproject.org/)

```sh
ls *.go | entr make
```

## Running

You probably have python

```sh
python -m SimpleHTTPServer 8080
```

If you want/have [goexec](https://github.com/shurcooL/goexec#goexec)

```sh
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

open [http://localhost:8080](http://localhost:8080)

all:
	@GOOS=js GOARCH=wasm go build -o js/main.wasm .

.PHONY: static/lib.wasm serve

static/lib.wasm:
	GOARCH=wasm GOOS=js go build -o static/lib.wasm main.go

serve:
	go run server.go -dir "`pwd`/static/" -listen "0.0.0.0:6969"
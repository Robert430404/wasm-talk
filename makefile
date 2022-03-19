build-go:
	cd wasm && GOOS=js GOARCH=wasm go build -o ../build/print.wasm

start-dev:
	go run main.go
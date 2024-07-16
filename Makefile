build-example:
	GOOS=js GOARCH=wasm go build -o ./dist/cherry.wasm ./examples/incrementing_buttons/*.go 

run-fileserver:
	GOOS=linux GOARCH=amd64 go run ./fileserver/fileserver.go

build-api:
	cd api && \
	GOOS=linux GOARCH=amd64 go build -o ../build/main main.go
	cd build && \
	zip api.zip main && \
	rm main

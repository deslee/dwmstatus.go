all: build

build: main.go
	go build -o status *.go

clean:
	rm status

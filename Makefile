build:
	go build -o releases/cstore main.go

install: build
	cp releases/cstore /opt/bin
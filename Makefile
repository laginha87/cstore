build:
	go build -o releases/store main.go

install: build
	cp releases/store /opt/bin
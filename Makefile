run:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata static		
	go run *.go

build:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata static		
	go build -o go-go-server-info *.go

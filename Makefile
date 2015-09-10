run:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -o gogoserverinfo static		
	go run *.go

build:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -o gogoserverinfo static		
	go build *.go

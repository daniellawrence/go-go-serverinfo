run:    build
	./go-go-server-info

build:
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata static		
	go build -o go-go-server-info *.go

clean:
	rm bindata.go go-go-server-info

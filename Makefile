depends:
	cd src ; go get -v -d

all: depends
	go build src/toris.go

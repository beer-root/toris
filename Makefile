depends:
	cd src ; go get -v -d

all: depends
	go build -o toris src/*.go

clean:
	rm toris

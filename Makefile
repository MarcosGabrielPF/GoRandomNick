.PHONY: build install clean run

build:
	go build -o grn .
install:
	sudo install -m 755 grn /usr/local/bin/grn

run:
	go run .

clean:
	rm -f grn
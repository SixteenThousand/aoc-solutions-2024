adv24: *.go
	go build -o adv24 .
build: adv24
run: adv24
	./adv24
test: *.go
	go test -v

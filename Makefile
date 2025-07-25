# Super simple makefile
# TODO: running out of tiiiime

test:
	go test -v ./

docker-build:
	docker build -t "leakybucket":tag1 .

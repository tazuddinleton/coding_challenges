clean:
	rm ./bin/ccwc

build: clean
	go build -o ./bin/ccwc

test:
	go test -v -count=1

run-byte-count: 
	./bin/ccwc -c test.txt

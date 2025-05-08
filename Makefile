
.PHONY: all
all: 
	go run .

.PHONY: build
build:
	go build -o main .

.PHONY: test
test:
	go test -v -race ./... | tee result.log | tail -20

.PHONY: clean
clean:
	rm -f main
	rm -f *.out
	rm -f *.log
	rm -f *.exe
	rm -f *.test
	rm -f *.test.exe
	rm -f *.test.out
	rm -rf testdata
	rm -rf testdata.zip

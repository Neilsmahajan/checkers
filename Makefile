build:
	go build -o ./bin/checkers ./cmd/checkers/main.go

clean:
	rm -rf bin

run: build
	./bin/checkers

all: compile run

compile:
	echo "Compiling..."
	go build -o ./bin/timeit ./cmd/timeit/main.go

run:
	./bin/timeit


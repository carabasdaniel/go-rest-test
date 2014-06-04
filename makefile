all: build

build:
        docker build -no-cache -rm -t docker:4050/go:1.2 .

run:
        docker run -rm docker:4050/go:1.2
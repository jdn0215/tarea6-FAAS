build:
    GOBIN=${PWD}/functions go get ./...
    GOBIN=${PWD}/functions go install ./...

clean:
    rm -f functions/*
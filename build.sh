rm -rf ./bin
mkdir bin
go build -o bin/lib.so -buildmode=c-shared main.go

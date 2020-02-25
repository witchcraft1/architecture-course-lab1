default: out/example
clean:
	rm -rf out
	rm -f cmd/main/buildVersion.go
test:
	go vet && go test
out/example: implementation.go cmd/main/main.go
	mkdir -p out
	./buildVersion.sh
	go build -o out/example ./cmd/main

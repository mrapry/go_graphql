.PHONY : all clean format test cover

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

clean:
	[ -f go-graphql-osx ] && rm go-graphql-osx || true
	[ -f go-graphql-linux ] && rm go-graphql-linux || true
	[ -f go-graphql32.exe ] && rm go-graphql32.exe || true
	[ -f go-graphql64.exe ] && rm go-graphql64.exe || true
	[ -f coverage.txt ] && rm coverage.txt || true
	rm ./coverages/*.txt

go-graphql-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

go-graphql-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

go-graphql64.exe: main.go
	GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o $@

go-graphql32.exe: main.go
	GOOS=windows GOARCH=386 go build -ldflags '-s -w' -o $@

go-graphql-windows: go-graphql64.exe b2c32.exe

go-graphql: go-graphql-osx go-graphql-linux go-graphql-windows

test: ./helper
	go test -race -short \
		 ./helper \


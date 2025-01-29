# runs tests
.PHONY: test
test:
	go test ./...

# runs tests and shows coverage
.PHONY: coverage
coverage:
	go test -covermode=count -coverprofile coverage.out.tmp ./...
	cat coverage.out.tmp | grep -v "mocks.go" > coverage.out
	rm coverage.out.tmp
	go tool cover -func=coverage.out

# runs tests and shows coverage map
.PHONY: coveragemap
coveragemap:
	go test -coverprofile coverage.out.tmp ./...
	cat coverage.out.tmp | grep -v "mocks.go" > coverage.out
	rm coverage.out.tmp
	go tool cover -html=coverage.out

# Checks project and source code if everything is according to standard
.PHONY: check
check:
	@gofmt -l ${GOFILES} | read && echo "Code differs from gofmt's style" 1>&2 && exit 1 || true
	go vet ${GOPACKAGES}

# Runs gofmt -w on the project's source code, modifying any files that do not
# match its style.
.PHONY: fmt
fmt:
	gofmt -l -w ${GOFILES}

# Runs gofmt -s -w on the project's source code, modifying any files that do not
# match its style.
.PHONY: simplify
simplify:
	gofmt -l -s -w ${GOFILES}

# makes homebrew zips for mac
.PHONY: homebrew
homebrew:
	cp dist/gosweeper-darwin-amd64 dist/gosweeper
	zip -jm dist/gosweeper-${VERSION}-darwin-amd64.zip dist/gosweeper
	cp dist/gosweeper-darwin-arm64 dist/gosweeper
	zip -jm dist/gosweeper-${VERSION}-darwin-arm64.zip dist/gosweeper

#!/usr/bin/make -f

test:
	@go version
	go fmt ./...
	go mod tidy
	@echo
	go test -cover -count=1 -timeout=1s -race ./...

doc:
	printf '# ' > README.md && \
		head -n 1 go.mod | sed 's/^module //' >> README.md && \
		echo >> README.md && \
		echo >> README.md && \
		go doc -all list  | sed 's/^/\t/' >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md
		go doc -all set  | sed 's/^/\t/' >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md

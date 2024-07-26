#!/usr/bin/make -f

test:
	@go version
	go fmt ./...
	go mod tidy
	go test -cover -count=1 -timeout=1s -race ./...

doc:
	printf '# ' > README.md && \
		head -n 1 go.mod | sed 's/^module //' >> README.md && \
		echo >> README.md && \
		echo >> README.md && \
		go doc -all github.com/mdwhatcott/go-collections/list  | sed 's/^/\t/' >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md
		go doc -all github.com/mdwhatcott/go-collections/queue  | sed 's/^/\t/' >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md
		go doc -all github.com/mdwhatcott/go-collections/set  | sed 's/^/\t/' >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md
		go doc -all github.com/mdwhatcott/go-collections/stack  | sed 's/^/\t/' >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md

.PHONY: all
all: .byeworld

REV=a000002

.byeworld: byeworld-linux-amd64 Dockerfile
	docker build -t quay.io/tamara/byeworld:master-${REV} .
	touch $@

byeworld-linux-amd64: byeworld.go
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X main.rev=${REV}" -o $@ $*

.PHONY: clean
clean:
	rm -f .byeworld byeworld-linux-amd64


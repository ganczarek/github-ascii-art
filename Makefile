default: test

test: get_dependencies libgit2
	go test -v ./...

get_dependencies:
	go get github.com/stretchr/testify
	go get github.com/google/uuid
	go get github.com/deckarep/golang-set
	go get -d github.com/libgit2/git2go

.ONESHELL:
libgit2:
	GOPATH=`go env GOPATH`
	cd $${GOPATH}/src/github.com/libgit2/git2go
	git submodule update --init
	make install

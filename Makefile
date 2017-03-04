default: test

test: get_dependencies libgit2
	go test -v ./...

get_dependencies:
	go get github.com/stretchr/testify
	go get github.com/google/uuid
	go get github.com/deckarep/golang-set
	go get github.com/libgit2/git2go

.ONESHELL:
libgit2:
	go get -d github.com/libgit2/git2go
	GOPATH=`go env | ack 'GOPATH="(?P<gopath>.*)"' --output "$$+{gopath}"`
	echo $${GOPATH}
	cd $${GOPATH}/src/github.com/libgit2/git2go
	git submodule update --init
	make install
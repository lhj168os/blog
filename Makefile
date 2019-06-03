#export GOPATH=/home/server/goprojects3

all: blog-new

blog:
	@echo 'building blog-new ...'
	@go build blog-new
	@echo 'build blog-new done'
.PHONY: blog-new

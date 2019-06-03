#export GOPATH=/home/server/goprojects3

all: blog

blog:
	@echo 'building blog ...'
	@go build blog
	@echo 'build blog done'
.PHONY: blog-new

.PHONY: help

help:
	@echo 'make <target>'
	@echo ''
	@echo 'Targets:'
	@echo ''
	@echo '    help    Show this help'
	@echo ''
	@echo '    build   Build the fibonacci package'
	@echo '    clean   Clean up build files'
	@echo '    init    Initialize fibonacci module'
	@echo '    test    Run the fibonacci package tests'
	@echo ''

build: clean init test
	@cd fibonacci && mkdir ../bin; go build -o ../bin .

clean:
	@cd fibonacci && rm -rf ../bin go.*

init: clean
	@cd fibonacci && go mod init fibonacci; go mod tidy

test: clean init
	@cd fibonacci && go test fibonacci -iter 14 -v -cover

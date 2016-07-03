default: clearscreen main

main:
	@go run main.go
	@echo ""
	@echo ""

clearscreen:
	@clear

database: destroydb createdb 

gitpush:
	git push 
	git push --tags

gofmt:
	go fmt ./...

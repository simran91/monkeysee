default: clearscreen main-quick

full: clearscreen main-full

main-quick:
	@go run main.go rgb.png
	@echo ""
	@echo ""

main-full:
	@go run main.go
	@echo ""
	@echo ""

clearscreen:
	@clear

gitpush:
	git push 
	git push --tags

gofmt:
	go fmt ./...

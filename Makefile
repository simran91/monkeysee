default: clearscreen main-quick

full: clearscreen clean main-full

# main-quick:
# 	@cd samples && go run ../main.go rgb.png 
# 	@echo ""
# 	@echo ""

main-quick:
	@cd samples && go run ../main.go rgb.png
	# @cd samples && go run ../main.go seam-test.png
	# @cd samples && go run ../main.go waves.png
	# @cd samples && go run ../main.go flower.jpg
	@echo ""
	@echo ""

main-full:
	@cd samples && go run ../main.go *.png *.jpg *.gif
	@echo ""
	@echo ""

flower: clearscreen
	@cd samples && go run ../main.go flower.jpg
	@echo ""
	@echo ""

forest: clearscreen
	@cd samples && go run ../main.go forest.png
	@echo ""
	@echo ""

png: clearscreen
	@cd samples && go run ../main.go *.png
	@echo ""
	@echo ""

clean:
	rm -fr samples/autogenerated

clearscreen:
	@clear

gitpush:
	git push 
	git push --tags

gofmt:
	go fmt ./...

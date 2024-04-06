build:
	@go build -o ./dist/main.exe ./cmd/main.go
run: build
	@./dist/main.exe
watch:
	@air

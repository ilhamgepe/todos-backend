build:
	@go build -o ./dist/main.exe ./cmd/main.go
run: build
	@./dist/main.exe
watch:
	@air
mCreate:
	@migrate create -ext sql -dir ./database/migrations $(name)
mUp:
	@migrate -database "mysql://root:root@tcp(localhost:3306)/todos" -path ./database/migrations up
mDown:
	@migrate -database "mysql://root:root@tcp(localhost:3306)/todos" -path ./database/migrations down
mUpTotal:
	@migrate -database "mysql://root:root@tcp(localhost:3306)/todos" -path ./database/migrations up $(total)
mDownTotal:
	@migrate -database "mysql://root:root@tcp(localhost:3306)/todos" -path ./database/migrations down $(total)

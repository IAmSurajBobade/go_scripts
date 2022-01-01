checkup:
	go mod tidy
	go vet .\migrate_csv.go
	go fmt .\migrate_csv.go
	golint .\migrate_csv.go
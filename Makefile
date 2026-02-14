
fmt::
	
# Run the application
run::
	go run cmd/api-server/main.go

# Test the application
test:
	go clean -testcache && go test -v -cover ./...

# Run mod tidy to resolve deps
tidy::
	go mod tidy -v

# Live Reload
watch::
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
		air; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing air...'; \
		go install github.com/air-verse/air@latest; \
		air; \
		Write-Output 'Watching...'; \
	}"


run:
	go run cmd/main.go

generate:
	mockery --version && echo "binary exists, generating..." || echo "need to install first"
	# GOBIN=$(pwd) go install github.com/vektra/mockery/v2@latest && sudo mv mockery /usr/local/bin/
	go generate ./...



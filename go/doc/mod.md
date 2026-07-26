# go工程化

```bash
go mod init github.com/wxz6289/go/version
go mod edit -replace github.com/wxz6289/go/version=./version
go get github.com/wxz6289/go/version
go mod tidy
```
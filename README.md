
### interpretar
```batch
go run main.go
```

### compilar
```batch
go build -o main.exe main.go

set CGO_ENABLED=0 
set GOOS=linux
set GOARCH=amd64
go build -o main.bin main.go

set CGO_ENABLED=0 
set GOOS=darwin
set GOARCH=amd64
go build -o main.app main.go 

set CGO_ENABLED=0
set GOOS=freebsd
set GOARCH=386 
go build -o main.bsd main.go
```

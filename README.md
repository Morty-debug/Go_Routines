
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

### compilar
```bash
GOOS=windows GOARCH=386 CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc go build -o main.exe main.go
```

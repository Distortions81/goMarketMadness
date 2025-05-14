
#!/bin/bash
GOOS=windows GOARCH=amd64 go build -pgo=auto -trimpath -gcflags=all="-B" -ldflags="-s -w" -o goMarketMadness-win64.exe

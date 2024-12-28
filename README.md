how to install

```bash
# mac, linux
curl -fsSL https://raw.githubusercontent.com/ry0y4n/envlate/main/install.sh | sh

# windows
powershell -Command "Invoke-WebRequest -Uri https://raw.githubusercontent.com/ry0y4n/envlate/main/install.ps1 -OutFile install.ps1; .\install.ps1"
```

memo

```bash
go run -o envlate
./envlate --file path/to/envfile
```

goreleaser

```bash
# local
goreleaser release --snapshot --clean
```

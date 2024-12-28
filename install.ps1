$ErrorActionPreference = 'Stop'

[string]$Version

if (-not $Version) {
    Write-Host "Fetching the latest version..."
    try {
        $response = Invoke-WebRequest -Uri "https://api.github.com/repos/ry0y4n/envlate/releases/latest" -UseBasicParsing
        $json = $response.Content | ConvertFrom-Json
        $Version = $json.tag_name
        if (-not $Version) {
            Write-Error "Failed to fetch the latest version"
            exit 1
        }
    } catch {
        Write-Error "Failed to fetch the latest version"
        exit 1
    }
}

# TODO arm64, i386
$url = "https://github.com/ry0y4n/envlate/releases/download/$Version/envlate_Windows_x86_64.zip"
$output = "$env:temp\envlate.zip"
$installDir = "$env:LocalAppData\envlate"

Write-Host "Downloading envlate version $Version from $url"
Invoke-WebRequest -Uri $url -OutFile $output

Write-Host "Extracting envlate"
Expand-Archive -Path $output -DestinationPath $installDir

Write-Host "Adding envlate to PATH"
$oldPath = [Environment]::GetEnvironmentVariable('Path', [System.EnvironmentVariableTarget]::User)
$newPath = "$oldPath;$installDir"
[Environment]::SetEnvironmentVariable('Path', $newPath, [System.EnvironmentVariableTarget]::User)

Write-Host "Installation complete. Please restart your terminal."
$platforms = @(
    @{GOOS="darwin"; GOARCH="amd64"; Extension=""},
    @{GOOS="darwin"; GOARCH="arm64"; Extension=""}, # For M1 Macs
    @{GOOS="linux"; GOARCH="amd64"; Extension=""},
    @{GOOS="linux"; GOARCH="arm64"; Extension=""} # For 64-bit ARM Linux
)

foreach ($platform in $platforms) {
    $env:GOOS = $platform.GOOS
    $env:GOARCH = $platform.GOARCH
    $output = "plate-$($platform.GOOS)-$($platform.GOARCH)$($platform.Extension)"
    
    Write-Host "Building for $($platform.GOOS) $($platform.GOARCH)..."
    go build -o $output
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Error building for $($platform.GOOS) $($platform.GOARCH)"
    } else {
        Write-Host "Successfully built $output"
    }
}
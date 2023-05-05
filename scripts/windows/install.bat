@echo off

setlocal EnableDelayedExpansion

set "suimon_version=latest"

:: Retrieve the latest stable version of Go from the official website
for /f "usebackq delims=" %%v in (`powershell -Command "(Invoke-WebRequest -Uri 'https://golang.org/VERSION?m=text').Content.Trim()"`) do set "go_version=%%v"

echo Installing Suimon...
echo ======================================
echo.

:: Install chocolatey (if not already installed)
if not exist "%ProgramData%\chocolatey" (
    echo Chocolatey not found. Installing Chocolatey...
    powershell -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"
)

:: Install dependencies
choco upgrade -y --no-progress --limit-output --no-colors
choco install -y --no-progress --limit-output --no-colors wget jq git cmake postgresql llvm

:: Set environment variables
set "GOPATH=%USERPROFILE%\go"
set "Path=%Path%;C:\Program Files (x86)\LLVM\bin;%GOPATH%\bin"

cd %USERPROFILE% && (
    mkdir .suimon
    wget -O .suimon\suimon-testnet.yaml https://raw.githubusercontent.com/bartosian/suimon/main/static/suimon.template.yaml
)

:: Install Go
if exist "%ProgramFiles%\Go" (
    echo Go %go_version% is already installed.
) else (
    echo Installing Go %go_version%...
    echo ======================================
    echo.

    powershell -Command "(Invoke-WebRequest -Uri 'https://golang.org/dl/go%go_version%.windows-amd64.msi' -UseBasicParsing -OutFile go.msi)"
    msiexec /i go.msi /qn /l*v go_install.log
    del go.msi

    echo Go %go_version% has been installed successfully.
)

:: Get the latest tag from the GitHub API
$LATEST_TAG=(Invoke-WebRequest -Uri https://api.github.com/repos/bartosian/suimon/releases/latest).content | ConvertFrom-Json | Select-Object -expand TagName

:: Download the latest binary release from GitHub
$ErrorActionPreference = "Stop"
if (Invoke-WebRequest -Uri "https://github.com/bartosian/suimon/releases/download/$LATEST_TAG/suimon-windows-latest-arm64" -OutFile "suimon") {
    Write-Host "Error: Failed to download suimon binary"
    exit 1
}

:: Move the binary to the executable directory
Move-Item -Path ".\suimon" -Destination "C:\Windows\System32\" -Force

# Make the binary executable
$oldACL = Get-Acl "C:\Windows\System32\suimon.exe"
$newACL = New-Object System.Security.AccessControl.FileSecurity
$newACL.SetAccessRuleProtection($true, $false)
$newACL.SetOwner([System.Security.Principal.NTAccount]::new("BUILTIN\Administrators"))
$newACL.SetAccessRuleProtection($false, $true)
$rule = New-Object System.Security.AccessControl.FileSystemAccessRule("BUILTIN\Users","ExecuteFile","Allow")
$newACL.SetAccessRule($rule)
Set-Acl "C:\Windows\System32\suimon.exe" $newACL

echo.
echo ======================================
echo Suimon has been installed and configured successfully.
echo Before running Suimon, you will need to customize the 'suimon-testnet
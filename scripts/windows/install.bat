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
    wget -O .suimon\suimon-testnet.yaml https://raw.githubusercontent.com/bartosian/suimon/main/static/templates/suimon-testnet.yaml
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
$LATEST_TAG = ((Invoke-WebRequest -Uri https://api.github.com/repos/bartosian/suimon/releases/latest).Content | ConvertFrom-Json).TagName

:: Define the file name for the release based on the new naming format
$RELEASE_FILE = "suimon_Windows_x86_64.zip"

:: Download the latest binary release from GitHub
$ErrorActionPreference = "Stop"
try {
    Invoke-WebRequest -Uri "https://github.com/bartosian/suimon/releases/download/$LATEST_TAG/$RELEASE_FILE" -OutFile $RELEASE_FILE
} catch {
    Write-Host "Error: Failed to download suimon release"
    exit 1
}

:: Extract the binary from the .zip file
try {
    Expand-Archive -Path $RELEASE_FILE -DestinationPath ".\"
} catch {
    Write-Host "Error: Failed to extract suimon binary from $RELEASE_FILE"
    exit 1
}

:: Move the binary to the executable directory
try {
    Move-Item -Path ".\suimon.exe" -Destination "C:\Windows\System32\" -Force
} catch {
    Write-Host "Error: Failed to move suimon binary to C:\Windows\System32\"
    exit 1
}

:: Make the binary executable
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
echo 🚀 Suimon has been installed and configured successfully. 🎉
echo 📝 Before running Suimon, you will need to customize the 'suimon-testnet.yaml' file in the '$HOME/.suimon' directory with the values specific to your environment. 🛠️
echo 👉 To get started, run 'suimon help'. 💡
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

REM Install the suimon module
go install "github.com/bartosian/suimon@%suimon_version%"

REM Check for errors in the install command
if %errorlevel% neq 0 (
  REM Check if the error message says "module declares its path as"
  findstr /C:"module declares its path as" suimon.out >nul
  if %errorlevel% equ 0 (
    echo Error: module path mismatch
    echo Please update the import path for the suimon module
  ) else (
    echo Error: failed to install suimon
  )
)

echo
echo "======================================"
echo "Suimon has been installed and configured successfully."
echo "Before running Suimon, you will need to customize the 'suimon-testnet.yaml' file in the '$HOME/.suimon' directory with the values specific to your environment."
echo "To get started, run 'suimon help'."

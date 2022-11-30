@echo off
echo NOTE: see also:
echo https://github.com/golang/go/wiki/WindowsCrossCompiling
echo https://github.com/golang/go/wiki/InstallFromSource#install-c-tools
"./go-winres.exe" make
md releases

::win32
call :build GuessNumber windows_386

::win64
call :build GuessNumber windows_amd64

::linux_amd64
call :build GuessNumber linux_amd64

::linux_arm64(android)
call :build GuessNumber linux_arm64

::Apple
call :build GuessNumber darwin_amd64

set GOOS=
set GOARCH=
set CGO_ENABLED=0
goto :eof

:build
set APP=%1
set PLATFORM=%2
:: Split param into GOOS & GOARCH (see: http://ss64.com/nt/syntax-substring.html)
set GOARCH=%PLATFORM:*_=%
call set GOOS=%%PLATFORM:_%GOARCH%=%%
:: Build filename
set FNAME=%APP%_%PLATFORM%
if "%GOOS%"=="windows" set FNAME=%FNAME%.exe
:: Do the build
echo == %FNAME% ==
go build -i -v -o ./releases/%FNAME% .
goto :eof

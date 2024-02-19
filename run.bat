@echo off

rem Step 1: Build the Go program
go build

rem Step 2: Execute the resulting executable
.\printonapp.exe

rem Pause to keep the Command Prompt window open after execution (optional)
pause

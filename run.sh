#!/bin/bash

# Step 1: Build the Go program
go build

# Step 2: Execute the resulting executable
.\printonapp.exe

# Optionally, pause to keep the terminal window open after execution (not necessary in Linux)
read -p "Press Enter to exit"

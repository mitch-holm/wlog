#!/bin/bash

# Exit on error
set -e

echo "Building wlog..."

# Build the application
go build -o wlog cmd/wlog/main.go

# Create bin directory if it doesn't exist
mkdir -p ~/bin

# Install to user's bin directory
echo "Installing to ~/bin..."
mv wlog ~/bin/

# Make sure the bin directory is in PATH
if [[ ":"$PATH":" != *":$HOME/bin:"* ]]; then
    echo "Adding ~/bin to PATH in .zshrc..."
    echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
    echo "Please restart your terminal or run 'source ~/.zshrc' to update PATH"
fi

echo "Installation complete! You can now use 'wlog' from anywhere." 
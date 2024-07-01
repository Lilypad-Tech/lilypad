#!/bin/bash

# Detect your machine's architecture and set it as $OSARCH
OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "amd64"; else print "unsupported_arch"}')
export OSARCH

# Detect your operating system and set it as $OSNAME
OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "darwin"; else if ($1 == "Linux") print "linux"; else print "unsupported_os"}')
export OSNAME

# Get the latest version URL
LATEST_URL=$(curl -s https://api.github.com/repos/lilypad-tech/lilypad/releases/latest | grep "browser_download_url.*lilypad-$OSNAME-$OSARCH" | cut -d : -f 2,3 | tr -d \")

# Extract the latest version (including short SHA) from the URL
LATEST_VERSION=$(echo $LATEST_URL | sed -n 's#.*/download/\([^/]*\)/.*#\1#p')
echo "Latest version: $LATEST_VERSION"

# Get the current version
CURRENT_VERSION=$(/usr/local/bin/lilypad version | grep "Lilypad:" | awk '{print $2}')
echo "Current version: $CURRENT_VERSION"

# Check if CURRENT_VERSION has a value
if [ -z "$CURRENT_VERSION" ]; then
    echo "Error: Unable to determin CURRENT_VERSION."
    exit 1
fi

if [ "$CURRENT_VERSION" != "$LATEST_VERSION" ]; then
  echo "Updating lilypad binary from version $CURRENT_VERSION to $LATEST_VERSION"
  sudo systemctl stop lilypad-resource-provider
  echo "Stopped the service: sudo systemctl stop lilypad-resource-provider"
  
  curl -L -o lilypad "$LATEST_URL"
  echo "Downloaded the latest version from $LATEST_URL"
  
  chmod +x lilypad
  sudo mv lilypad /usr/local/bin/lilypad
  echo "Made the new binary executable and moved it to /usr/local/bin/lilypad"
  
  sudo systemctl start lilypad-resource-provider
  echo "Restarted the service: sudo systemctl start lilypad-resource-provider"
  
  echo "Update complete."
else
  echo "Lilypad binary is already up to date."
fi
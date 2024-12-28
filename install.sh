#!/bin/bash

# If no version is specified as a command line argument, fetch the latest version.
if [ -z "$1" ]; then
    VERSION=$(curl -s https://api.github.com/repos/ry0y4n/envlate/releases/latest | grep -o '"tag_name": *"[^"]*"' | sed 's/"tag_name": *"//' | sed 's/"//')
    if [ -z "$VERSION" ]; then
        echo "Failed to fetch the latest version"
        exit 1
    fi
else
    VERSION=$1
fi

OS=$(uname -s)
ARCH=$(uname -m)
URL="https://github.com/ry0y4n/envlate/releases/download/${VERSION}/envlate_${OS}_${ARCH}.tar.gz"

echo "Start to install."
echo "VERSION=$VERSION, OS=$OS, ARCH=$ARCH"
echo "URL=$URL"

TMP_DIR=$(mktemp -d)
curl -L $URL -o $TMP_DIR/envlate.tar.gz
tar -xzvf $TMP_DIR/envlate.tar.gz -C $TMP_DIR
sudo mv $TMP_DIR/envlate /usr/local/bin/envlate
sudo chmod +x /usr/local/bin/envlate

rm -rf $TMP_DIR

if [ -f "/usr/local/bin/envlate" ]; then
  echo "[SUCCESS] envlate $VERSION has been installed to /usr/local/bin"
else
  echo "[FAIL] envlate $VERSION is not installed."
fi
#!/bin/bash

mkdir -p ./deb/cvecli_0.0.1-1_amd64/usr/bin
cp ./cvecli_darwin_amd64/cvecli ./deb/cvecli_0.0.1-1_amd64/usr.bin
mkdir -p ./deb/cvecli_0.0.1-1_amd64/DEBIAN
cat > ./deb/cvecli_0.0.1-1_amd64/DEBIAN/control << EOF
Package: cvecli
Version: 0.0.1
Maintainer: kyle@thepublicclouds.com
Architecture: amd64
Homepage: https://github.com/wizedkyle/cvecli
Description: A CLI tool that allows CNAs to manage their organisation and submit CVEs
EOF
dpkg --build ./deb/cvecli_0.0.1-1_amd64

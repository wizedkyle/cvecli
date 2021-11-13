#!/bin/bash

while getopts "a:v:" flag
do
  case "$flag" in
    "a") architecture=${OPTARG};;
    "v") version=${{OPTARG}};;
  esac
done

mkdir -p "./deb/cvecli_$version-1_$architecture/usr/bin"
cp ./cvecli_darwin_amd64/cvecli "./deb/cvecli_$version-1_$architecture/usr.bin"
mkdir -p "./deb/cvecli_$version-1_$architecture/DEBIAN"
cat > "./deb/cvecli_$version-1_$architecture/DEBIAN/control" << EOF
Package: cvecli
Version: $version
Maintainer: kyle@thepublicclouds.com
Architecture: $architecture
Homepage: https://github.com/wizedkyle/cvecli
Description: A CLI tool that allows CNAs to manage their organisation and submit CVEs
EOF
dpkg --build "./deb/cvecli_$version-1_$architecture"

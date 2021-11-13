#!/bin/bash

architecture=""
version=""

while true; do
  case "$1" in
    -a)
      architecture=$1
      ;;
    -v)
      version=$1
      ;;
    --)
      shift
      break
      ;;
    esac
    shift
done

echo $architecture
echo $version

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

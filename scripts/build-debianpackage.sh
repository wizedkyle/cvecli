#!/bin/bash

architecture=""
version=""

while getopts ":a:v:" options; do
  case "${options}" in
    a)
      architecture=${OPTARG}
      ;;
    v)
      version=${OPTARG}
      ;;
    :)
      echo "Error: -${OPTARG} requires an argument"
      exit 1
      ;;
    *)
      exit 1
      ;;
  esac
done

echo "=> Creating debian package folder structure"
mkdir -p "./deb/cvecli_$version-1_$architecture/usr/bin"
echo "=> Copying cvecli binary"
cp ./cvecli_darwin_amd64/cvecli "./deb/cvecli_$version-1_$architecture/usr/bin"
echo "=> Creating debian control file"
mkdir -p "./deb/cvecli_$version-1_$architecture/DEBIAN"
cat > "./deb/cvecli_$version-1_$architecture/DEBIAN/control" << EOF
Package: cvecli
Version: $version
Maintainer: kyle@thepublicclouds.com
Architecture: $architecture
Homepage: https://github.com/wizedkyle/cvecli
Description: A CLI tool that allows CNAs to manage their organisation and submit CVEs
EOF
echo "=> Building debian package"
dpkg --build "./deb/cvecli_$version-1_$architecture"

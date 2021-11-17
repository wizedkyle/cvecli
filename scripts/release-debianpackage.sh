#!/bin/bash

releaseArchitectures=""
version=""

while getopts ":a:v:" options; do
  case "${options}" in
    a)
      architectures+=("$OPTARG")
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

echo "=> Creating apt repo folder"
mkdir ./aptcvecli
echo "=> Syncing S3 bucket locally"
aws s3 sync s3://aptthepublicclouds/cvecli ./aptcvecli --debug
echo "=> Creating pools directory"
mkdir -p ./aptcvecli/pool/main
for architecture in "${architectures[@]}"; do
  releaseArchitectures+=architecture
  releaseArchitectures+=" "
  echo "=> Moving $architecture debian package to local apt repo"
  mv "./cvecli_$version-1_$architecture.deb" "./aptcvecli/pool/main/cvecli_$version-1_$architecture.deb"
  echo "=> Creating $architecture packages directory"
  mkdir -p "./aptcvecli/dists/stable/main/binary-$architecture"
  echo "=> Removing old $architecture package files"
  rm "./aptcvecli/dists/stable/main/binary-$architecture/Packages"
  rm "./aptcvecli/dists/stable/main/binary-$architecture/Packages.gz"
  echo "=> Generate new $architecture package file"
  dpkg-scanpackages --arch "$architecture" ./aptcvecli/pool/ > "./aptcvecli/dists/stable/main/binary-$architecture/Packages"
  echo "=> Compressing $architecture package file"
  gzip -k "./aptcvecli/dists/stable/main/binary-$architecture/Packages"
done
echo "=> Removing old release files"
rm ./aptcvecli/dists/stable/Release
rm ./aptcvecli/dists/stable/Release.gpg
rm ./aptcvecli/dists/stable/InRelease
./aptcvecli/dists/stable/Release << EOF
Origin: apt.thepublicclouds.com
Suite: stable
Codename: stable
Version: $version
Architectures: $releaseArchitectures
Components: main
Description: A CLI tool that allows CNAs to manage their organisation and submit CVEs.
Date: $(date)
EOF
cat ./aptcvecli/dists/stabe/Release

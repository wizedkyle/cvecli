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
echo "=> Checking pools directory"
if [ -d "./aptcvecli/pool/main" ]; then
  echo "=> ./aptcvecli/pool/main already exists"
else
  echo "=> Creating pools directory"
  mkdir -p ./aptcvecli/pool/main
fi
for architecture in "${architectures[@]}"; do
  releaseArchitectures+=architecture
  releaseArchitectures+=" "
  echo "=> Moving $architecture debian package to local apt repo"
  mv "./cvecli_$version-1_$architecture.deb" "./aptcvecli/pool/main/cvecli_$version-1_$architecture.deb"
  echo "=> Checking for $architecture packages directory"
  if [ -d "./aptcvecli/dists/stable/main/binary-$architecture" ]; then
    echo "=> ./aptcvecli/dists/stable/main/binary-$architecture already exists"
  else
    mkdir -p "./aptcvecli/dists/stable/main/binary-$architecture"
  fi
  echo "=> Checking for old $architecture package files"
  if [ -f "./aptcvecli/dists/stable/main/binary-$architecture/Packages" ]; then
    echo "=> Removing ./aptcvecli/dists/stable/main/binary-$architecture/Packages"
    rm "./aptcvecli/dists/stable/main/binary-$architecture/Packages"
  else
    echo "=> ./aptcvecli/dists/stable/main/binary-$architecture/Packages does not exist"
  fi
  if [ -f "./aptcvecli/dists/stable/main/binary-$architecture/Packages.gz" ]; then
    echo "=> Removing ./aptcvecli/dists/stable/main/binary-$architecture/Packages.gz"
    rm "./aptcvecli/dists/stable/main/binary-$architecture/Packages.gz"
  else
    echo "=> ./aptcvecli/dists/stable/main/binary-$architecture/Packages.gz does not exist"
  fi
  echo "=> Generate new $architecture package file"
  dpkg-scanpackages --arch "$architecture" ./aptcvecli/pool/ > "./aptcvecli/dists/stable/main/binary-$architecture/Packages"
  echo "=> Compressing $architecture package file"
  gzip -k "./aptcvecli/dists/stable/main/binary-$architecture/Packages"
done
echo "=> Checking for old release files"
if [ -f "./aptcvecli/dists/stable/Release" ]; then
  echo "=> Removing ./aptcvecli/dists/stable/Release"
  rm ./aptcvecli/dists/stable/Release
else
  echo "=> ./aptcvecli/dists/stable/Release does not exist"
fi
if [ -f "./aptcvecli/dists/stable/Release.gpg" ]; then
  echo "=> Removing ./aptcvecli/dists/stable/Release.gpg"
  rm ./aptcvecli/dists/stable/Release.gpg
else
  echo "=> ./aptcvecli/dists/stable/Release.gpg does not exist"
fi
if [ -f "./aptcvecli/dists/stable/InRelease" ]; then
  echo "=> Removing ./aptcvecli/dists/stable/InRelease"
  rm ./aptcvecli/dists/stable/InRelease
else
  echo "=> ./aptcvecli/dists/stable/InRelease does not exist"
fi
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

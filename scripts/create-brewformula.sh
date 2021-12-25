#!/bin/bash

version=""

generate_hash() {
  file=$1
  sha256sum "$file" | cut -d" " -f1
}

while getopts ":v:" options; do
  case "${options}" in
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

cat << EOF > cvecli.rb
class Cvecli < Formula
  desc "cvecli"
  homepage "https://github.com/wizedkyle/cvecli"
  version "$version"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/wizedkyle/cvecli/releases/download/$version/cvecli_${version}_darwin_amd64.tar.gz"
      sha256 "$(generate_hash "cvecli_${version}_darwin_amd64.tar.gz")"

      def install
        bin.install "cvecli"
      end
    end

    if Hardware::CPU.arm?
      url "https://github.com/wizedkyle/cvecli/releases/download/$version/cvecli_${version}_darwin_arm64.tar.gz"
      sha256 "$(generate_hash "cvecli_${version}_darwin_arm64.tar.gz")"

      def install
        bin.install "cvecli"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/wizedkyle/cvecli/releases/download/$version/cvecli_${version}_linux_amd64.tar.gz"
      sha256 "$(generate_hash "cvecli_${version}_linux_amd64.tar.gz")"

      def install
        bin.install "cvecli"
      end
    end

    if Hardware::CPU.arm?
      url "https://github.com/wizedkyle/cvecli/releases/download/0.1.0/cvecli_${version}_linux_amd64.tar.gz"
      sha256 "$(generate_hash "cvecli_${version}_linux_arm64.tar.gz")"

      def install
        bin.install "cvecli"
      end
    end
  end
end
EOF

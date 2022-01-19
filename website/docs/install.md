# Install

`cvecli` is available for Linux, macOS and Windows on both amd64 and arm64 architectures.
If there is no pre-compiled binaries that meet your requirements you can build `cvecli` from source.

## Install the pre-compiled binary

### Linux

=== "brew"
    ```shell
    brew tap wizedkyle/homebrew-tap
    brew install wizedkyle/tap/cvecli
    ```

=== "apt"
    ```shell
    apt-key adv --fetch-keys https://apt.cvecli.app/public.key
    add-apt-repository "deb https://apt.cvecli.app/ stable main"
    apt-get update
    apt-get install cvecli
    ```

=== "dpm"
    ```shell
    ```

### macOS

`cvecli` supports both Apple Silicon and Intel Macs, brew will install the correct binary for your operating system.

The macOS binaries are signed and notarised which should prevent any issues with Gatekeeper.

=== "brew"
    ```shell
    brew tap wizedkyle/homebrew-tap
    brew install wizedkyle/tap/cvecli
    ```

### Windows

Currently the only way to install Windows binaries is to download them the [releases page](https://github.com/wizedkyle/cvecli/releases).

### Manually

You can install pre-compiled binaries for `cvecli` by navigating to the [releases page](https://github.com/wizedkyle/cvecli/releases).

## Running with Docker

=== "docker"
    Registries:
    
    - [`ghcr.io/wizedkyle/cvecli`](https://github.com/wizedkyle/cvecli/pkgs/container/cvecli)
    
    Example usage:
    
    ```shell
    ```

=== "podman"
    Registries:

    - [`ghcr.io/wizedkyle/cvecli`](https://github.com/wizedkyle/cvecli/pkgs/container/cvecli)
    
    Example usage:
    
    ```shell
    ```

## Compiling from source

If you want to compile from source you can perform the following steps:

**Clone:**

```shell
git clone https://github.com/wizedkyle/cvecli/cvecli
cd cvecli
```

**Get the Go dependencies:**

```shell
go mod tidy
```

**Build:**

```shell
go build ./cmd/cvecli
```

**Verify it works:**

```shell
./cvecli --version
```

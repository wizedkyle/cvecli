# CVECLI

`cvecli` allows you to interact with the [CVE Services API](https://github.com/CVEProject/cve-services) via the command line. 
It currently supports the following functionality for CNAs:
- Reserving CVE IDs
- Managing user accounts

As more features are released in CVE Services, `cvecli` will be updated to support these.

## Installation

### Container Image

Container images for `cvecli` are hosted in GitHub Packages and can be pulled with the following command:

```shell
docker pull ghcr.io/wizedkyle/cvecli:$VERSION-$ARCHITECTURE
```

### Linux

You can install `cvecli` via apt on debian based linux distributions by running the following commands:

```shell
apt-key adv --fetch-keys https://apt.cvecli.app/public.key
add-apt-repository "deb https://apt.cvecli.app/ stable main"
apt-get update
apt-get install cvecli
```

### macOS

You can install `cvecli` via homebrew on macOS by running the following commands:

```shell
brew tap wizedkyle/homebrew-tap
brew install wizedkyle/tap/cvecli
```

### Windows

You can download a signed binary file from the specific release you want.

### Build from source

You can build `cvecli` from source using the following commands:

```shell
git clone https://github.com/wizedkyle/cvecli.git
```

## Authentication

`cvecli` supports two authentication methods:
1. Environment Variables which are useful when using `cvecli` without user interaction
2. Credentails File which is preferred when using `cvecli` with user interaction

### Environment Variables

Environment variable authentication is useful when running cvecli in a CI/CD pipeline. The following environment variables need to be set to allow for proper authentication.

```
CVE_API_USER: example@example.com

CVE_API_KEY: AbCeFG123

CVE_ORGANIZATION: OrganizationName

CVE_ENVIRONMENT: https://cveawg.mitre.org/api or https://cveawg-test.mitre.org/api
```
The CVE Services environment URLs are as follows:
* Production: https://cveawg.mitre.org/api
* Test: https://cveawg-test.mitre.org/api

### Credentials File

`cvecli` can use a credential file stored on disk to authenticate which can be generated interactively using `cvecli configure`.
All details in the credentials file is encrypted at rest using AES256 encryption.

The credential file is stored in the following locations depending on your operating system.

```
Windows: C:\Users\<username>\.cvecli\credentials\creds.json

Macos: /User/<username>/.cvecli/credentials/creds.json

Linux: /Usr/<username>/.cvecli/credentials/creds.json
```

The contents of the credential file is as follows:

```
{
  "apiUser": "abcefghi",
  "apiKey": "AbCeFG123",
  "organization": "<organization name>",
}
```

## Documentation

Documentation for each command can be found by using `cvecli <command> --help`

## Contributing

Clone or fork the repo, make your changes and create a pull request.
I will then review it and all things looking good it gets merged!

If there is something in the code that you don't understand please feel free to email at kyle@thepublicclouds.com.


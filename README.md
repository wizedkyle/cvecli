# CVECLI

`cvecli` allows you to interact with the [CVE Services API](https://github.com/CVEProject/cve-services) via the command line. 
It currently supports the following functionality for CNAs:
- Reserving CVE IDs
- Managing user accounts

As more features are released in CVE Services, `cvecli` will be updated to support these.

## Installation

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

`cvecli` uses a credential file stored on disk to authenticate which can be generated interactively using `cvecli configure`.
All details in the credentails file is encrypted at rest using AES256 encryption.

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


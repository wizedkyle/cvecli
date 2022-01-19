`cvecli` supports two authentication methods:

- Environment Variables which are useful when using `cvecli` without user interaction
- Credentials File which is preferred when using `cvecli` with user interaction

## Environment Variables

Environment variable authentication is useful when running `cvecli` in a CI/CD pipeline. 
The following environment variables need to be set to allow for proper authentication.

```
CVE_API_USER: example@example.com

CVE_API_KEY: AbCeFG123

CVE_ORGANIZATION: OrganizationName

CVE_ENVIRONMENT: https://cveawg.mitre.org/api or https://cveawg-test.mitre.org/api
```
The CVE Services environment URLs are as follows:

- [Production](https://cveawg.mitre.org/api)
- [Test](https://cveawg-test.mitre.org/api)

## Credentials File

`cvecli` can use a credential file stored on disk to authenticate which can be generated interactively using `cvecli` configure. 
All details in the credentials file is encrypted at rest using AES256 encryption.

The credential file is stored in the following locations depending on your operating system.

```
Windows: C:\Users\<username>\.cvecli\credentials\creds.json

Macos: /User/<username>/.cvecli/credentials/creds.json

Linux: /Usr/<username>/.cvecli/credentials/creds.json
```

The contents of the credential file is as follows:

```json
{
  "apiUser": "abcefghi",
  "apiKey": "AbCeFG123",
  "organization": "<organization name>",
  "environment": "production or test environment url"
}
```

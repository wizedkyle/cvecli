# cvecli create-user

Creats a new user in the organization

The user account that is being used to authenticate needs to have the ADMIN role to perform this action

```shell
cvecli create-user [flags]
```

## Options

```
-f, --first-name    - Specify the first name of the user
-h, --help          - help for create-user
-l, --last-name     - Specify the last name of the user
-m, --middle-name   - Specify the middle name of the user (if applicable)
-r, --roles         - Specify the roles for the user comma separated. Valid roles are: 'ADMIN'. Only add the user as an ADMIN if you want them to have control over the organization
-s, --suffix        - Specify the suffix of the user (if applicable)
-u, --username      - Specify the email address of the user
```

## See also

* [cvecli](/cmd/cvecli) - A CLI tool that allows CNAs to manage their organisation and CVEs.


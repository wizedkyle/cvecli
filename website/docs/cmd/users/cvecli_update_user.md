# cvecli update-user

Updates a user from the organization

```shell
cvecli update-user [flags]
```

## Options

```
-a, --role-to-add    - Specify the role for the user. Valid roles are: 'ADMIN'. Only add the user as an ADMIN if you want them to have control over the organization
-e, --enabled        - Set to false if you want to disable the user
-f, --first-name     - Specify the first name of the user
-h, --help           - help for update-user
-l, --last-name      - Specify the last name of the user
-m, --middle-name    - Specify the middle name of the user (if applicable)
-n, --new-username   - Specify the new email address of the user
-r, --role-to-remove - Specify the role to remove from the user. Valid roles are: 'ADMIN'
-s, --suffix         - Specify the suffix of the user (if applicable)
-u, --username       - Specify the current email address of the user
```

## See also

* [cvecli](/cmd/cvecli) - A CLI tool that allows CNAs to manage their organisation and CVEs.

# cvecli

A CLI tool that allows CNAs to manage their organisation and CVEs.

## Synopsis

`cvecli` allows you to interact with the [CVE Services API](https://github.com/CVEProject/cve-services) via the command line.

`cvecli` currently supports the following functionality:

* Manage users in CNA organisation 
* Manage CVE IDs

`cvecli` allows staff of an CNA organization to manage their CVE IDs via the CLI or use `cvecli` as a part of a CI/CD pipeline.

## Options

```
-h, --help      help for cvecli
-j, --json      outputs the response in json
-v, --version   version for cvecli
```

## See also

* [cvecli check-id-quota](/cmd/cve-ids/cvecli_check_id_quota/) - Checks the CVE ID quotas for the organization
* [cvecli configure](/cmd/cvecli_configure) - Sets credentials for `cvecli`
* [cvecli create-user](/cmd/users/cvecli_create_user) - Creates a new user in the organization
* [cvecli get-cve-id](/cmd/cve-ids/cvecli_get_cve_id/) - Retrieves a CVE ID record by the ID
* [cvecli get-organization-info](/cmd/organization/cvecli_get_organization_info) - Retrieves information about the organization the user authenticating is apart of
* [cvecli get-user](/cmd/users/cvecli_get_user) - Retrieves information about a user in the organization
* [cvecli list-cve-ids](/cmd/cve-ids/cvecli_list_cve_ids) - Lists all CVE Ids associated to an organization.
* [cvecli list-users](/cmd/users/cvecli_list_users) - Retrieves all users from the organization
* [cvecli reserve-cve-id](/cmd/cve-ids/cvecli_reserve_cve_id) - Reserves a CVE ID for the organization
* [cvecli reset-secret](/cmd/users/cvecli_reset_secret) - Resets the secret for a user in the organization
* [cvecli update-user](/cmd/users/cvecli_update_user) - Updates a user record from the organization

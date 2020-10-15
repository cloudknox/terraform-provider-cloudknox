---
layout: "cloudknox"
page_title: "Provider: CloudKnox"
description: |-
   The CloudKnox provider is used to interact with actions supported in the CloudKnox API
---

# CloudKnox Provider

The CloudKnox provider is used to interact with actions supported in the CloudKnox API.


## Example Usage

```hcl
# Configure CloudKnox provider
provider "cloudknox" {
   shared_credentials_file = "./credentials.conf"
   profile = "alpha" 
}

# Create a role-policy
data "cloudknox_role_policy" "my_policy" {
   # ...
}
``` 

## Provider Setup

See the [Configuration Reference](/docs/guides/configuration_reference.md) guide to learn how to setup the provider configuration files. 

## Authentication

The CloudKnox provider supports 2 methods of providing credentials for authentication. See the [Authentication Reference](/docs/guides/authentication_reference.md) guide to learn more.

- Credentials file
- Environment variables

## Argument Reference

The following arguments are supported:

### Optional

* `shared_credentials_file` - (Optional) String containing the path and filename of the credentials file you would like to use. If no file is specified, the default credentials file or environment variables will be used. See the [Authentication Reference](/docs/guides/authentication_reference.md) guide to learn more.
* `profile` - (Optional) String containing the profile that is in the credentials file (default or specified). `profile` has a default value of `default` if not specified.




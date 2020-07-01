# Terraform Provider for Cloudknox

## Requirements
* Terraform 0.12.28
* Go 1.14.3

## Building the Provider

1. Clone the provider repository
2. Navigate inside the directory containing the source
3. Build the provider

```bash
sudo make build
```

* The provider will be built and stored in `~/.terraform.d/plugins`
* Logs are stored in `/var/log/cloudknox/application.log`
* Backend configuration file is created and stored in `/opt/cloudknox/terraform-provider-cloudknox-config.yaml`


## Using the Provider

### Initialization

Define the cloudnox terraform provider as shown below

```terraform
provider "cloudknox" {
    shared_credentials_file = "" //Optional
    profile = "" //Optional
}
```

### Credentials

#### Default Credentials File

Will be used if a `shared_credentials_file` is not provided

Place `creds.conf` in `~/.cnx/`

* `~/.cnx/` directory is created during build process

```HOCON
profiles {
    default {
        service_account_id = "######"
        access_key = "######"
        secret_key = "######"
    }

    other_profile {
        service_account_id = "######"
        access_key = "######"
        secret_key = "######"
    }
}
```

#### Shared Credentials File

Set the `shared_credentials_file` property in `main.tf` to the path containing a HOCON File filled out like this

```HOCON
profiles {
    default {
        service_account_id = "######"
        access_key = "######"
        secret_key = "######"
    }

    other_profile {
        service_account_id = "######"
        access_key = "######"
        secret_key = "######"
    }
}
```

#### Profiles

Set the `profile` property in `main.tf` to the profile you would like to use in your config file

`default` profile will be used if none is specified

#### Environment Variables

If no configuration file is specified and the default credentials file does not exist, the following environment variables will be checked for credentials.

Export these environment variables:

```bash
CNX_SERVICE_ACCOUNT_ID="#####"
CNX_ACCESS_KEY="#####"
CNX_SECRET_KEY="#####"
```

## Resources

### cloudknox_policy

#### Effects

Creates a `<name>.tf` file containing a terraform resource with a right-sized policy for AWS, GCP or Azure for the provided users or cloud resources.

#### Properties

- `name` : Name of the policy, can match the terraform resource name
- `output_path` : Directory where the terraform script will be outputted
- `auth_system_info` : Set to the following map

```
{
    id : Enter the id as a string
    type : Choose AWS, GCP or AZURE as a string (VCENTER NOT Currently Supported)
}
```

- `identity_type` : Choose `USER` or `RESOURCE`
- `identity_ids` : Provide a comma seperated list of strings containing `ids` of type `auth_system_info`
- `filter_history_days` : Number of days in the past to look at the actions of `identity_ids` to generate a policy
- `filter_preserve_reads` : Optional parameter used for `AZURE`
- `filter_history_start_time_millis` : Start time in unix time milliseconds to look at actions of `identity_ids`
- `filter_history_end_time_millis` : End time in unix time milliseconds to look at actions of `identity_ids`
- `request_params_scope` : Optional parameter for Cloudknox API
- `request_params_resource` : Optional parameter for Cloudknox API
- `request_params_resources` : Optional list of parameters for Cloudknox API
- `request_params_condition` : Optional parameter for Cloudknox API

---
**NOTE**

Use `filter_history_days` or `filter_history_start_time_millis` and `filter_history_end_time_millis` together as only one parameter will be considered when generating a policy. 

---


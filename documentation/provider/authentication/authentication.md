# CloudKnox Terraform Provider Authentication

## Overview

The CloudKnox Terraform Provider is a client that interacts CloudKnox's API, therefore each client must aquire credentials in order to use this provider. 

## Credential Aquisition

1. Login to the [CloudKnox Portal](app.cloudknox.io)
2. Click on your user account settings in the top right hand corner
3. Choose `Service Accounts` where you will be presented with a view of all active service accounts
4. Click on `Add Service Account` in the upper right hand corner, enter a `Name` and `Description` and click `Save`
    * The `Service Account` will appear in the table.
5. Access the `Service Account` options by clicking the 3 dot menu to the left of the `Service Account` with corresponding `Name`
6. `Add Access Key` to add an `Access Key` to this `Service Account`
    * Your `Service Account` can have a number of `Access Key`s
7. You will be presented with your `Access Key` and `Secret Key`
    * Note down your `Secret Key` as it is not viewable once you close the popup
    * Your active `Access Key`s  can be viewed anytime by clicking on the number of `Keys` associated with `Name`
8. At this point you have the 3 `Credentials` required for a `Service Account` to authenticate with the CloudKnox API and use the CloudKnox Terraform Provider

## Provider Credentials Setup

There are 2 main ways to use the `Credentials` you aquired with the CloudKnox Terraform Provider. Choose whichever option works with your workflow the best. 

### Credentials File

#### Default Credentials File

#### File Setup

Place `creds.conf` in `~/.cnx/`

* `~/.cnx/` directory is created during build process, but if it has been deleted or removed for any reason, ensure it exists. 

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

#### Provder Setup

Your provider declaration should look like this in `main.tf` if you want to use the `Default Credentials File`

```terraform
provider "cloudknox" {
    profile = "" //Optional
}
```

If no `profile` is provided, the `default` profile will be used. 


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

### Environment Variables

If no configuration file is specified and the default credentials file does not exist, the following environment variables will be checked for credentials.

Export these environment variables:

```bash
CNX_SERVICE_ACCOUNT_ID="#####"
CNX_ACCESS_KEY="#####"
CNX_SECRET_KEY="#####"
```
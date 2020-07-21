# CloudKnox Terraform Provider Authentication

## Overview

The CloudKnox Terraform Provider is a client that interacts CloudKnox's API, therefore each client must aquire credentials in order to use this provider.

## Credential Aquisition

1. Login to the [CloudKnox Portal](app.cloudknox.io)
2. Click on your user account settings in the top right hand corner
3. Choose `Service Accounts` where you will be presented with a view of all active service accounts
4. Click on `Add Service Account` in the upper right hand corner, enter a `Name` and `Description` and click `Save`
    * The `Service Account` will appear in the table.
5. Access the `Service Account` options by clicking the 3 dot menu to the right of the `Service Account` row with corresponding `Name`
6. `Add Access Key` to add an `Access Key` to this `Service Account`
    * Your `Service Account` can have a number of `Access Key`s
7. You will be presented with your `Access Key` and `Secret Key`
    * Note down your `Secret Key` as it is not viewable once you close the popup
    * Your active `Access Key`s  can be viewed anytime by clicking on the number of `Keys` associated with `Name`
8. At this point you have the 3 `Credentials` required to authenticate with the CloudKnox API and use the CloudKnox Terraform Provider

## Provider Credentials Setup

There are 2 main ways to use the `Credentials` you aquired with the CloudKnox Terraform Provider. Choose whichever option works with your workflow the best. 

### Credentials File

#### Default Credentials File

##### File Setup

Place `credentials.conf` in `~/.cloudknox/`

* `~/.cloudknox/` directory is created during build process, but if it has been deleted or removed for any reason, ensure it exists. 

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

##### Provder Setup

Your provider declaration should look like this in `main.tf` if you want to use the `Default Credentials File`.

```terraform
provider "cloudknox" {
    profile = "" //Optional
}
```

If no `profile` is provided, the `default` profile will be used. 


#### Shared Credentials File

##### File Setup

Create a a credentials file `credentials` in the directory of your choosing and ensure it is formatted as a HOCON file as shown below.

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

##### Provder Setup

Your provider declaration should look like this in `main.tf` if you want to use the `Shared Credentials File`

```terraform
provider "cloudknox" {
    shared_credentials_file = "/path/to/credentials" //Required
    profile = "" //Optional
}
```

If no `profile` is provided, the `default` profile will be used. 

#### Profiles

Set the `profile` property in `main.tf` to the profile you would like to use in either your `Default Credentials File` or `Shared Credentials File`. 

`default` profile will be used if none is specified

### Environment Variables

If no configuration file is specified and the default credentials file does not exist, the following environment variables will be checked for credentials.

##### Environment Variable Setup

1. Export these environment variables:

```bash
CNX_SERVICE_ACCOUNT_ID="#####"
CNX_ACCESS_KEY="#####"
CNX_SECRET_KEY="#####"
```
2. Delete the `Default Credentials File`
    * If the `Default Credentials File` is found in `~/.cloudknox/` then that file will be used so ensure that the file is deleted or renamed. 


##### Provder Setup

Your provider declaration should look like this in `main.tf` if you want to use `Environemnt Variables`

```terraform
provider "cloudknox" {
    // No parameters should be set
}
```




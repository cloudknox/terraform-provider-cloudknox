# Terraform Provider for Cloudknox

## Terraform Configuration
Create `main.tf` in the `terraform-provider-cloudknox` folder (TESTING)
```terraform
provider "cloudknox" {
    shared_credentials_file = "" //Optional
    profile = "" //Optional
}

resource "cloudknox_policy" "new-policy" { //meaningless resource for testing purposes
    address = "1" //Meaningless Parameter for testing purposes
}
```
### Credentials
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

#### Default Credentials File

Will be used if a `shared_credentials_file` is not provided

Place `creds.conf` in `~\.cnx\` folder or `C:\Users\%USER_PROFILE%\.cnx\`

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

CNX_SERVICE_ACCOUNT_ID="#####" \
CNX_ACCESS_KEY="#####" \
CNX_SECRET_KEY="#####"

### Testing (Cloudknox)

On Linux: 
```go 
go build -o terraform-provider-cloudknox
```
On Windows: 
```go
go build -o terraform-provider-cloudknox.exe
```
```bash
terraform init
terraform plan //Will Produce output at info.log
terraform apply //Will create the resource, not working as of now
```




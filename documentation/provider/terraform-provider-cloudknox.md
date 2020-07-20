# Terraform Provider for CloudKnox

## Setup

### Terraform

Declare the CloudKnox Terraform Provider as shown below in your `main.tf`

```terraform
provider "cloudknox" {
    shared_credentials_file = "" //Optional
    profile = "" //Optional
}
```

### Authentication

See [Authentication](./authentication/authentication.md) for details and examples on authentication setup and how to set the provider parameters accordingly. 



## Available Resources
* [cloudknox_role_policy](./resources/resource_cloudknox_role_policy/resource_cloudknox_role_policy.md)
    * [AWS](./resources/resource_cloudknox_role_policy/aws/resource_cloudknox_role_policy_aws.md)
    * [AZURE](./resources/resource_cloudknox_role_policy/azure/resource_cloudknox_role_policy_azure.md)
    * [GCP](./resources/resource_cloudknox_role_policy/gcp/resource_cloudknox_role_policy_gcp.md)
    * [VCENTER](./resources/resource_cloudknox_role_policy/vcenter/resource_cloudknox_role_policy_vcenter.md)


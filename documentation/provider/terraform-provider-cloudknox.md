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



## Data Sources
* [cloudknox_role_policy](./data_sources/data_source_role_policy/data_source_role_policy.md)
    * [AWS](./data_sources/data_source_role_policy/aws/data_source_role_policy_aws.md)
    * [AZURE](./data_sources/data_source_role_policy/azure/data_source_role_policy_azure.md)
    * [GCP](./data_sources/data_source_role_policy/gcp/data_source_role_policy_gcp.md)
    * [VCENTER](./data_sources/data_source_role_policy/vcenter/data_source_role_policy_vcenter.md)


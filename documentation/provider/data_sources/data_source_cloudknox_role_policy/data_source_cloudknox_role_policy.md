# CloudKnox Role-Policy Data Source

## Overivew

Creates a `<name>.tf` file containing a Cloud Provider Terraform Resource with a Least-Privilege Role-Policy for AWS, GCP or Azure for the provided identities

## Usage

The `cloudknox_role_policy` data source will create different outputs based on the authorization system. Parameters can be set based on authorization system type to support the same functionality available through the [CloudKnox Portal](app.cloudknox.io)'s JEP Controller

AWS

1. [Activity of User(s)](./aws/data_source_cloudknox_role_policy_aws_activity_users.md)
2. [Activity of Resources(s)](./aws/data_source_cloudknox_role_policy_aws_activity_resources.md)
3. [Activity of Role](./aws/data_source_cloudknox_role_policy_aws_activity_role.md)

Azure

1. [Activity of User(s)](./azure/data_source_cloudknox_role_policy_azure_activity_users.md)
2. [Activity of App(s)](./azure/data_source_cloudknox_role_policy_azure_activity_apps.md)


GCP

1. [Activity of User(s)](./gcp/data_source_cloudknox_role_policy_gcp_activity_users.md)
2. [Activity of Service Account(s)](./gcp/data_source_cloudknox_role_policy_gcp_activity_service_accounts.md) (Not Currently Supported)

vCenter

1. [Activity of User(s)](./vcenter/data_source_cloudknox_role_policy_vcenter_activity_users.md) (Not Currently Supported)


## Properties Overview

These are all the properties available to set during data source declaration. The above documentation will explain in detail what each parameter does and what they can be set as in order to achieve the desired result along with example data source declarations and outputs. 

- `name` : Name of the policy, can match the terraform data source name
- `output_path` : Directory where the terraform script will be outputted
- `auth_system_info` : Set to the following map

```
{
    id : Enter the id as a string
    type : Choose AWS, GCP or AZURE as a string (VCENTER NOT Currently Supported)
}
```

- `identity_type` : Identity type of the ids
- `identity_ids` : Provide a comma seperated list of strings containing `ids` of type `auth_system_info`
- `filter_history_days` : Number of days in the past to look at the actions of `identity_ids` to generate a policy
- `filter_preserve_reads` : Boolean to indicate preserve read permissions granted before (Only on Azure)
- `filter_history_start_time_millis` : Start time in unix time milliseconds to look at actions of `identity_ids`
- `filter_history_end_time_millis` : End time in unix time milliseconds to look at actions of `identity_ids`
- `request_params_scope` : Optional parameter for Cloudknox API
- `request_params_resource` : Optional parameter for Cloudknox API
- `request_params_resources` : Optional list of parameters for Cloudknox API
- `request_params_condition` : Optional parameter for Cloudknox API

---
**NOTE**

Not all parameters are required when declaring your `cloudknox_role_policy` data source. Some parameters only apply to certain Authorization System Types

Use `filter_history_days` or `filter_history_start_time_millis` and `filter_history_end_time_millis` together as only one parameter will be considered when generating a policy. 

---
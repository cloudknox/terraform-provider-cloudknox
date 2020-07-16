# CloudKnox Role-Policy Resource

## Overivew

Creates a `<name>.tf` file containing a Terraform resource with a Least-Privilege Role-Policy for AWS, GCP or Azure for the provided identi

## Usage

The `cloudknox_role_policy` resource will create different outputs based on the authorization system. Parameters can be set based on authorization system type to support the same functionality available through the [CloudKnox Portal](app.cloudknox.io)'s JEP Controller

AWS

1. [Activity of User(s)](./aws/resource_cloudknox_role_policy_aws_activity_users.md)
2. [Activity of Groups(s)](./aws/resource_cloudknox_role_policy_aws_activity_groups.md)
3. [Activity of Resources(s)](./aws/resource_cloudknox_role_policy_aws_activity_resources.md)
4. [Activity of Role](./aws/resource_cloudknox_role_policy_aws_activity_role.md)
5. [From Existing Policy](./aws/resource_cloudknox_role_policy_aws_from_existing_policy.md)
6. [New Policy](./aws/resource_cloudknox_role_policy_aws_new_policy.md)

Azure

1. [Activity of User(s)](./azure/resource_cloudknox_role_policy_azure_activity_users.md)
2. [Activity of Groups(s)](./azure/resource_cloudknox_role_policy_azure_activity_groups.md)
3. [Activity of App(s)](./azure/resource_cloudknox_role_policy_azure_activity_apps.md)
4. [From Existing Role](./azure/resource_cloudknox_role_policy_azure_from_existing_role.md)
5. [New Role](./azure/resource_cloudknox_role_policy_azure_new_role.md)

GCP

1. [Activity of User(s)](./gcp/resource_cloudknox_role_policy_gcp_activity_users.md)
2. [Activity of Groups(s)](./gcp/resource_cloudknox_role_policy_gcp_activity_groups.md)
3. [Activity of Service Account(s)](./gcp/resource_cloudknox_role_policy_gcp_activity_service_accounts.md)
4. [From Existing Role](./gcp/resource_cloudknox_role_policy_gcp_from_existing_role.md)
5. [New Role](./gcp/resource_cloudknox_role_policy_gcp_activity_new_role.md)

vCenter

1. [Activity of User(s)](./vcenter/resource_cloudknox_role_policy_vcenter_activity_users.md)
2. [Activity of Groups(s)](./vcenter/resource_cloudknox_role_policy_vcenter_activity_groups.md)
3. [From Existing Role](./vcenter/resource_cloudknox_role_policy_vcenter_from_existing_role.md)
4. [New Role](./vcenter/resource_cloudknox_role_policy_vcenter_new_role.md)

## Properties Overview

These are all the properties available to set during resource declaration. The above documentation will explain in detail what each parameter does and what they can be set to to achieve the desired result along with example resource declarations. 

- `name` : Name of the policy, can match the terraform resource name
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
- `filter_preserve_reads` : Optional parameter for Cloudknox API
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
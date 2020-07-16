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
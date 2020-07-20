# Terraform Provider for CloudKnox

## Requirements
* Terraform 0.12.28
* Go 1.14.3

## Building the Provider

1. Clone the provider repository
2. Navigate inside the directory containing the source
3. Build the provider

```bash
make build
```

* The provider will be built and stored in `~/.terraform.d/plugins`
* The API configuration file is required and must be created in `~/.cloudknox/api.conf`

The API configuration file should be populated as follows.
```HOCON
api {
    base_url: "https://olympus.aws-staging.cloudknox.io"
}
```

## Documentation

Detailed [documentation](./documentation/provider/terraform-provider-cloudknox.md) is provided. Please use this documentation for provider setup and usage.

## Debug

Logging is disabled by default. 

Set the following environment variables to enable logging and set `CNX_LOG_LEVEL` is set to one of the options below. 

```bash
export CNX_LOG_LEVEL={"ALL", "DEBUG", "ERROR", "INFO", "WARN", "NONE"}
export CNX_LOG_OUTPUT="/path/to/logs/application.log"
```

Disable logging by unsetting the environment variables

```bash
unset CNX_LOG_LEVEL
unset CNX_LOG_OUTPUT
```





# Terraform Provider for CloudKnox

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

[Documentation](./documentation/provider/terraform-provider-cloudknox.md) is seperated by `resource` and further seperated by `Authorization System Type` and is located in `./documention`







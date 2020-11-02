---
layout: "cloudknox"
page_title: "CloudKnox Provider Configuration Reference"
description: |-
  Configuration reference for the CloudKnox provider for Terraform.
---


# CloudKnox Provider Configuration Reference

## API Configuration

The API configuration file is required and must be created in `~/.cloudknox/api.conf` and populated as follows:

* `base_url` - (Required) Get the `base_url` from the CloudKnox Integrations page available in your user settings on the CloudKnox Portal. 

```HOCON
api {
    base_url: "http://base.url"
}
```

### Authentication

See the [Authentication](/providers/cloudknox/cloudknox/latest/docs/guides/authentication_reference.md) guide to learn more.



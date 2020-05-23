package main

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudknox.Provider,
	})
}

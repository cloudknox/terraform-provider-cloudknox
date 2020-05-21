package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/sakethkollu/terraform-provider-cloudknox/cloudknox"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudknox.Provider,
	})
}

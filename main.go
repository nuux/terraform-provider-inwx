package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"terraform-provider-inwx/inwx"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: inwx.Provider,
	})
}

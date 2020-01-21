package main

import (
	"github.com/hashicorp/terraform/plugin"
	"terraform-provider-inwx/inwx"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: inwx.Provider,
	})
}

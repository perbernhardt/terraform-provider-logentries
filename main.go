package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-logentries/logentries"
)

func main() {
	plugin.Serve(new(logentries.Provider))
}

package main

import (
	"github.com/paypal/dce-go/dce"

	// register the plugins
	_ "github.com/paypal/dce-go/plugin/example"
	_ "github.com/paypal/dce-go/plugin/general"
)

func main() {
	dce.Run()
}

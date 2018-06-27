package main

import (
	"github.com/7ojo/docker-machine-driver-upcloud/driver"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(upcloud.NewDriver("", ""))
}

package controllers

import (
	"fmt"

	"github.com/bartosian/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/suimon/internal/core/ports"
)

const version = "v1.1.0"

type VersionController struct {
	cliGateway *cligw.Gateway
}

func NewVersionController(
	cliGateway *cligw.Gateway,
) ports.VersionController {
	return &VersionController{
		cliGateway: cliGateway,
	}
}

func (c *VersionController) PrintVersion() {
	fmt.Println(version)
}

package controllers

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

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
	fmt.Println("v1.0.2")
}

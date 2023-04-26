package controllers

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

type RootController struct {
	cliGateway *cligw.Gateway
}

func NewRootController(
	cliGateway *cligw.Gateway,
) ports.RootController {
	return &RootController{
		cliGateway: cliGateway,
	}
}

func (c *RootController) BeforeStart() bool {
	return true
}

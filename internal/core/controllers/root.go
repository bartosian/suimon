package controllers

import (
	"github.com/bartosian/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/suimon/internal/core/ports"
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

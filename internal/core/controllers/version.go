package controllers

import (
	"log/slog"

	"github.com/bartosian/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/suimon/internal/core/ports"
)

const version = "v1.2.2"

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
	slog.Info("Suimon version", "version", version)
}

package monitor

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/config"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

type (
	Gateways struct {
		rpc        ports.RPCGateway
		geo        ports.GeoGateway
		prometheus ports.PrometheusGateway
		cli        *cligw.Gateway
	}

	Hosts struct {
		rpc       []host.Host
		node      []host.Host
		validator []host.Host
		peers     []host.Host
	}

	Builders struct {
		static  map[enums.TableType]ports.Builder
		dynamic map[enums.TableType]ports.Builder
	}

	Controller struct {
		logger log.Logger

		config   *config.Config
		hosts    Hosts
		gateways Gateways
		builders Builders
	}
)

func NewController(
	logger log.Logger,
	config *config.Config,
	rpcGW ports.RPCGateway,
	geoGW ports.GeoGateway,
	prometheusGW ports.PrometheusGateway,
	cliGW *cligw.Gateway,
) *Controller {
	return &Controller{
		logger: logger,
		config: config,
		gateways: Gateways{
			rpc:        rpcGW,
			geo:        geoGW,
			prometheus: prometheusGW,
			cli:        cliGW,
		},
	}
}

package monitor

import (
	"fmt"
	"sync"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/config"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

type (
	Gateways struct {
		cli *cligw.Gateway
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
		lock sync.RWMutex

		selectedTables     []enums.TableType
		selectedDashboards []enums.TableType

		config   *config.Config
		hosts    Hosts
		gateways Gateways
		builders Builders
	}
)

func NewController(
	config *config.Config,
	cliGW *cligw.Gateway,
) *Controller {
	return &Controller{
		config: config,
		gateways: Gateways{
			cli: cliGW,
		},
		builders: Builders{
			static:  make(map[enums.TableType]ports.Builder),
			dynamic: make(map[enums.TableType]ports.Builder),
		},
	}
}

// setBuilderForTable sets the builder for a given table type.
// It updates the `static` map in `c.builders` with the given builder for the specified table type.
func (c *Controller) setBuilderForTable(table enums.TableType, builder ports.Builder) {
	c.builders.static[table] = builder
}

// getHostsByTableType returns the list of hosts for a given table type.
// It acquires a read lock on the controller lock before accessing the hosts data.
func (c *Controller) getHostsByTableType(table enums.TableType) (hosts []host.Host, err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	switch table {
	case enums.TableTypeNode:
		return c.hosts.node, nil
	case enums.TableTypeValidator:
		return c.hosts.validator, nil
	case enums.TableTypeActiveValidators:
		return c.hosts.rpc[:1], nil
	case enums.TableTypeGasPriceAndSubsidy,
		enums.TableTypeValidatorsCounts,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:
		return c.hosts.rpc[:1], nil
	case enums.TableTypePeers:
		return c.hosts.peers, nil
	case enums.TableTypeRPC:
		return c.hosts.rpc, nil
	default:
		return nil, fmt.Errorf("unknown table type: %v", table)
	}
}

// setHostsByTableType updates the list of hosts for a given table type.
// It acquires a write lock on the controller lock before updating the hosts data.
func (c *Controller) setHostsByTableType(table enums.TableType, hosts []host.Host) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	switch table {
	case enums.TableTypeNode:
		c.hosts.node = hosts
	case enums.TableTypeValidator:
		c.hosts.validator = hosts
	case enums.TableTypePeers:
		c.hosts.peers = hosts
	case enums.TableTypeRPC:
		c.hosts.rpc = hosts
	default:
		return fmt.Errorf("unknown table type: %v", table)
	}

	return nil
}

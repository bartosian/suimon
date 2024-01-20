package monitor

import (
	"fmt"
	"sync"

	"github.com/bartosian/suimon/internal/core/domain/config"
	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/core/domain/release"
	"github.com/bartosian/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/suimon/internal/core/ports"
)

type Gateways struct {
	cli *cligw.Gateway
}

type Hosts struct {
	rpc       []host.Host
	node      []host.Host
	validator []host.Host
}

type Releases []Releases

type Builders struct {
	static  map[enums.TableType]ports.Builder
	dynamic map[enums.TableType]ports.Builder
}

type Controller struct {
	lock sync.RWMutex

	// nwtwork represents the currently selected network.
	network string

	// selectedConfig represents the currently selected configuration.
	selectedConfig config.Config

	// selectedTables stores the selected table types.
	selectedTables []enums.TableType

	// selectedDashboard represents the selected dashboard type.
	selectedDashboard enums.TableType

	// configs is a map of named configurations.
	configs map[string]config.Config

	// hosts stores different types of hosts.
	hosts Hosts

	// releases information
	releases []release.Release

	// gateways represent the available gateways.
	gateways Gateways

	// builders contain static and dynamic builders.
	builders Builders
}

// NewController creates a new instance of the Controller.
// It takes a map of configuration and a CLI gateway as input and returns a pointer to the Controller.
// The map of configuration is used to initialize the Controller's configs field.
// The CLI gateway is used to initialize the Controller's gateways field.
// The static and dynamic maps in the Builders field are initialized with empty maps.
// The newly created Controller instance is returned.
func NewController(
	configs map[string]config.Config,
	cliGW *cligw.Gateway,
) *Controller {
	return &Controller{
		configs: configs,
		gateways: Gateways{
			cli: cliGW,
		},
		builders: Builders{
			static:  make(map[enums.TableType]ports.Builder),
			dynamic: make(map[enums.TableType]ports.Builder),
		},
	}
}

func (c *Controller) getHostsByTableType(table enums.TableType) (hosts []host.Host, err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	switch table {
	case enums.TableTypeNode:
		return c.hosts.node, nil
	case enums.TableTypeValidator:
		return c.hosts.validator, nil
	case enums.TableTypeRPC:
		return c.hosts.rpc, nil
	case enums.TableTypeActiveValidators,
		enums.TableTypeGasPriceAndSubsidy,
		enums.TableTypeValidatorsParams,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports:

		if len(c.hosts.rpc) > 0 {
			return c.hosts.rpc[:1], nil
		}

		return nil, fmt.Errorf("no rpc hosts available for table type: %v", table)
	default:
		return nil, fmt.Errorf("unknown table type: %v", table)
	}
}

// setHostsByTableType sets the list of hosts for a given table type.
// It acquires a write lock on the controller lock before updating the hosts data.
// If the table type is unknown, it returns an error.
// The error is returned if the table type is unknown or if there is an error acquiring the lock.
func (c *Controller) setHostsByTableType(table enums.TableType, hosts []host.Host) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	switch table {
	case enums.TableTypeNode:
		c.hosts.node = hosts
	case enums.TableTypeValidator:
		c.hosts.validator = hosts
	case enums.TableTypeRPC:
		c.hosts.rpc = hosts
	default:
		return fmt.Errorf("unknown table type: %v", table)
	}

	return nil
}

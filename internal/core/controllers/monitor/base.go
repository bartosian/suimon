package monitor

import (
	"fmt"
	"sync"

	"github.com/bartosian/suimon/internal/core/domain/config"
	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/core/domain/metrics"
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
	builders Builders // 8 bytes

	// Pointers (8 bytes each on 64-bit systems), group together
	configs        map[string]config.Config // 8 bytes
	gateways       Gateways                 // 8 bytes
	selectedConfig config.Config            // size depends on the struct definition

	// Strings and slices (16 bytes each on 64-bit systems), group together
	network string // 16 bytes (pointer + len)

	// Enum and struct types, smaller than pointers, group together
	selectedDashboard enums.TableType   // size depends on the underlying type, typically int32 or int
	hosts             Hosts             // 8 bytes
	releases          []metrics.Release // 8 bytes

	selectedTables []enums.TableType // 16 bytes (pointer + len + cap)

	lock sync.RWMutex // 8 bytes on 64-bit, keep first due to its size and to avoid false sharing

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

// getHostsByTableType is a method of the Controller struct that returns a list of hosts for a given table type.
// It acquires a read lock on the controller lock before accessing the hosts data.
// The method uses a switch statement to determine which type of hosts to return based on the table type.
// The method returns a list of hosts and an error. The error is returned if the table type is unknown or if there are no RPC hosts available for the specified table types.
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
		enums.TableTypeValidatorReports,
		enums.TableTypeProtocol:
		if len(c.hosts.rpc) > 0 {
			return c.hosts.rpc[:1], nil
		}

		return nil, fmt.Errorf("no rpc hosts available for table type: %v", table)
	case enums.TableTypeReleases:
		return nil, fmt.Errorf("no hosts available for table type: %v", table)
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
	case enums.TableTypeActiveValidators,
		enums.TableTypeGasPriceAndSubsidy,
		enums.TableTypeValidatorsParams,
		enums.TableTypeValidatorsAtRisk,
		enums.TableTypeValidatorReports,
		enums.TableTypeProtocol,
		enums.TableTypeReleases:
		return nil
	default:
		return fmt.Errorf("unknown table type: %v", table)
	}

	return nil
}

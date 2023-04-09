package ports

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

type Builder interface {
	Init(table enums.TableType, hosts []host.Host) error
	Render() error
}

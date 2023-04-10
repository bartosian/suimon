package ports

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

type Builder interface {
	Init(hosts []host.Host) error
	Render()
}

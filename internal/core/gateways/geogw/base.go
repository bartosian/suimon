package geogw

import (
	"context"
	"net/http"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"

	"github.com/bartosian/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/suimon/internal/core/ports"
)

const (
	ipInfoCacheExp    = 5 * time.Minute
	httpClientTimeout = 4 * time.Second
)

type Gateway struct {
	ctx         context.Context
	accessToken string
	client      *ipinfo.Client
	cliGateway  *cligw.Gateway
}

func NewGateway(cliGW *cligw.Gateway, accessToken string) ports.GeoGateway {
	httpClient := &http.Client{Timeout: httpClientTimeout}
	infoCache := ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp))
	geoClient := ipinfo.NewClient(httpClient, infoCache, accessToken)

	return &Gateway{
		ctx:         context.Background(),
		accessToken: accessToken,
		client:      geoClient,
		cliGateway:  cliGW,
	}
}

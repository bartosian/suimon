package geogw

import (
	"context"
	"net/http"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const (
	ipInfoCacheExp    = 5 * time.Minute
	httpClientTimeout = 4 * time.Second
)

type Gateway struct {
	ctx         context.Context
	url         string
	accessToken string
	client      *ipinfo.Client
	logger      log.Logger
}

func NewGateway(logger log.Logger, url string, accessToken string) ports.GeoGateway {
	httpClient := &http.Client{Timeout: httpClientTimeout}
	infoCache := ipinfo.NewCache(cache.NewInMemory().WithExpiration(ipInfoCacheExp))
	geoClient := ipinfo.NewClient(httpClient, infoCache, accessToken)

	return &Gateway{
		ctx:         context.Background(),
		url:         url,
		accessToken: accessToken,
		client:      geoClient,
		logger:      logger,
	}
}

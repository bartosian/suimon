package prometheusgw

import (
	"context"
	"net/http"
	"time"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const httpClientTimeout = 4 * time.Second

type Gateway struct {
	ctx    context.Context
	url    string
	client *http.Client
	logger log.Logger
}

func NewGateway(logger log.Logger, url string) ports.PrometheusGateway {
	httpClient := http.Client{
		Timeout: httpClientTimeout,
	}

	return &Gateway{
		ctx:    context.Background(),
		url:    url,
		client: &httpClient,
		logger: logger,
	}
}

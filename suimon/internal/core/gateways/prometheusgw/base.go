package prometheusgw

import (
	"context"
	"net/http"
	"time"

	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

const httpClientTimeout = 4 * time.Second

type Gateway struct {
	ctx        context.Context
	url        string
	client     *http.Client
	cliGateway *cligw.Gateway
}

func NewGateway(cliGW *cligw.Gateway, url string) ports.PrometheusGateway {
	httpClient := http.Client{
		Timeout: httpClientTimeout,
	}

	return &Gateway{
		ctx:        context.Background(),
		url:        url,
		client:     &httpClient,
		cliGateway: cliGW,
	}
}

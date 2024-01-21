package prometheusgw

import (
	"context"
	"net/http"
	"time"

	"github.com/bartosian/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/suimon/internal/core/ports"
)

const httpClientTimeout = 3 * time.Second

type Gateway struct {
	ctx        context.Context
	client     *http.Client
	cliGateway *cligw.Gateway
	url        string
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

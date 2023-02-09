package checker

import (
	"errors"
	"sync"
	"time"

	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"github.com/oschwald/geoip2-golang"
)

const (
	httpClientTimeout = 2 * time.Second
	pathToGeoDB       = "./vendors/geodb/GeoLite2-Country.mmdb"
	pathToRPCList     = "rpc.yaml"
)

func (checker *Checker) ParseData() error {
	if len(checker.nodeYaml.Config.SeedPeers) == 0 {
		return errors.New("no peers found in config file")
	}

	filePath, err := filepath.Abs(pathToGeoDB)
	if err != nil {
		return err
	}

	db, err := geoip2.Open(filePath)
	if err != nil {
		return err
	}

	defer db.Close()

	checker.geoDbClient = db

	var (
		wg      sync.WaitGroup
		errChan = make(chan error)
	)

	wg.Add(3)

	go func() {
		defer wg.Done()

		if err := checker.parseRPCHosts(); err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()

		if err := checker.parseNode(); err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()

		if err := checker.parsePeers(); err != nil {
			errChan <- err
		}
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var errCounter int

	for parseErr := range errChan {
		err = multierror.Append(err, parseErr)

		errCounter++
	}

	if errCounter == 3 {
		return err
	}

	return nil
}

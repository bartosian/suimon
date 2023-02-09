package checker

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
)

func (checker *Checker) parseRPCHosts() error {
	var (
		wg      sync.WaitGroup
		hostCH  = make(chan RPCHost)
		rpcList []RPCHost
	)

	filePath, err := filepath.Abs(pathToRPCList)
	if err != nil {
		return err
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var result RPCList
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return err
	}

	hosts := result.GetByNetwork(checker.network)

	for _, host := range hosts {
		wg.Add(1)

		go func(host string) {
			defer wg.Done()

			newRPC := newRPCHost(host)
			doneCH := make(chan struct{})

			go func() {
				newRPC.GetTotalTransactionNumber()

				doneCH <- struct{}{}
			}()

			go func() {
				newRPC.GetLatestCheckpoint()

				doneCH <- struct{}{}
			}()

			for i := 0; i < 2; i++ {
				<-doneCH
			}

			defer close(doneCH)

			hostCH <- *newRPC
		}(host)
	}

	go func() {
		wg.Wait()
		close(hostCH)
	}()

	for rpc := range hostCH {
		rpc.SetStatus()

		rpcList = append(rpcList, rpc)
	}

	checker.rpcList = rpcList

	return nil
}

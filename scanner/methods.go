package scanner

import (
	"errors"
	"fmt"
	"github.com/heilart1n/cryptools/network"
)

func (scs *Scanners) AddScanner(scannerType string, scanner *Scanner) *Scanners {
	if scs.M == nil {
		scs.M = map[string]*Scanner{}
	}
	scs.M[scannerType] = scanner
	return scs
}

func (scs *Scanners) SetCurrentType(scannerType string) *Scanners {
	scs.CurrentType = scannerType
	return scs
}

func (scs *Scanners) Validate() error {
	if scs.CurrentType == "" {
		return errors.New("the current scanner is not installed")
	}
	if _, ok := scs.M[scs.CurrentType]; !ok {
		return errors.New("the currently installed scanner is not found in the map")
	}
	if len(scs.M) == 0 {
		return errors.New("at least one scanner must be added")
	}
	if scs.CurrentScanner().Endpoints[scs.CurrentNetwork] == "" {
		return errors.New("no scanner is installed for the network in use")
	}
	return nil
}

func (scs *Scanners) CurrentScanner() *Scanner {
	return scs.M[scs.CurrentType]
}

func (scs *Scanners) AddressUrl(address string) string {
	return scs.CurrentScanner().AddressUrl(scs.CurrentNetwork, address)
}

func (scs *Scanners) BlockUrl(block string) string {
	return scs.CurrentScanner().BlockUrl(scs.CurrentNetwork, block)
}

func (scs *Scanners) TxUrl(hash string) string {
	return scs.CurrentScanner().TxUrl(scs.CurrentNetwork, hash)
}

func (scanner *Scanner) AddressUrl(currentNetwork network.Type, address string) string {
	return fmt.Sprintf("%s/%s/%s", scanner.CurrentEndpoint(currentNetwork), scanner.AddressPath, address)
}

func (scanner *Scanner) BlockUrl(currentNetwork network.Type, block string) string {
	return fmt.Sprintf("%s/%s/%s", scanner.CurrentEndpoint(currentNetwork), scanner.BlockPath, block)
}

func (scanner *Scanner) TxUrl(currentNetwork network.Type, hash string) string {
	return fmt.Sprintf("%s/%s/%s", scanner.CurrentEndpoint(currentNetwork), scanner.TxPath, hash)
}

func (scanner *Scanner) CurrentEndpoint(currentNetwork network.Type) string {
	return scanner.Endpoints[currentNetwork]
}

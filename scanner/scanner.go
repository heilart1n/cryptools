package scanner

import "github.com/heilart1n/cryptools/network"

type (
	Scanners struct {
		CurrentNetwork network.Type        `json:"-" yaml:"-"`
		CurrentType    string              `json:"current_type" yaml:"current_type"`
		M              map[string]*Scanner `json:"l" yaml:"l"`
		DefaultScanID  string              `json:"-" yaml:"-"`
	}
	Scanner struct {
		Endpoints   map[network.Type]string `json:"endpoints" yaml:"endpoints"`
		AddressPath string                  `json:"address_path" yaml:"address_path"`
		TxPath      string                  `json:"tx_path" yaml:"tx_path"`
		BlockPath   string                  `json:"block_path" yaml:"block_path"`
	}
)

func NewSingle(endpoints map[network.Type]string, addressPath, txPath, blockPath string) *Scanner {
	return &Scanner{
		Endpoints:   endpoints,
		AddressPath: addressPath,
		TxPath:      txPath,
		BlockPath:   blockPath,
	}
}

func New(currentNetwork network.Type) *Scanners {
	return &Scanners{
		CurrentNetwork: currentNetwork,
	}
}

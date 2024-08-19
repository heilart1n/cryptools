package network

const (
	Mainnet Type = "mainnet"
	Testnet Type = "testnet"
	Devnet  Type = "devnet"
)

var (
	validTypeNames = map[Type]bool{
		Mainnet: true,
		Testnet: true,
		Devnet:  true,
	}
)

type (
	Network struct {
		RPCEndpoint   string `json:"rpc_endpoint" yaml:"rpc_endpoint"`
		WSSEndpoint   string `json:"wss_endpoint" yaml:"wss_endpoint"`
		UtilsEndpoint string `json:"utils_endpoint" yaml:"utils_endpoint"`
	}
	Networks struct {
		CurrentType Type              `yaml:"current" json:"current"`
		M           map[Type]*Network `yaml:"m" json:"m"`
	}
	Type string
)

func NewSingle() *Network {
	return &Network{}
}

func New() *Networks {
	return &Networks{}
}

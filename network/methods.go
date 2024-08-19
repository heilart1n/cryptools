package network

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/url"
)

func (network *Network) AddRPCEndpoint(e string) *Network {
	network.RPCEndpoint = e
	return network
}

func (network *Network) AddWSSEndpoint(e string) *Network {
	network.WSSEndpoint = e
	return network
}

func (network *Network) AddUtilsEndpoint(e string) *Network {
	network.UtilsEndpoint = e
	return network
}

func (networks *Networks) AddNetwork(nType Type, network *Network) *Networks {
	if networks.M == nil {
		networks.M = make(map[Type]*Network)
	}
	networks.M[nType] = network
	return networks
}

func (networks *Networks) SetCurrentType(t Type) *Networks {
	networks.CurrentType = t
	return networks
}

func (networks *Networks) Current() *Network {
	return networks.M[networks.CurrentType]
}

func (networks *Networks) Validate() error {
	if networks.CurrentType == "" {
		return errors.New("the current network is not installed")
	}
	if _, ok := networks.M[networks.CurrentType]; !ok {
		return errors.New("the currently installed network is not found in the map")
	}
	if len(networks.M) == 0 {
		return errors.New("at least one network must be added")
	}
	for nType, network := range networks.M {
		if !validTypeNames[nType] {
			return fmt.Errorf("invalid network type: %s", nType)
		}
		if err := network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (network *Network) Validate() error {
	return validation.ValidateStruct(network,
		validation.Field(&network.RPCEndpoint, validation.Required, validation.By(isValidURL)),
	)
}

func isValidURL(value interface{}) error {
	str, _ := value.(string)
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return fmt.Errorf("invalid URL")
	}
	return nil
}

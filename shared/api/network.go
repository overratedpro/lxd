package api

// NetworksPost represents the fields of a new LXD network
//
// API extension: network
type NetworksPost struct {
	NetworkPut `yaml:",inline"`

	Managed bool   `json:"managed" yaml:"managed"`
	Name    string `json:"name" yaml:"name"`
	Type    string `json:"type" yaml:"type"`
}

// NetworkPost represents the fields required to rename a LXD network
//
// API extension: network
type NetworkPost struct {
	Name string `json:"name" yaml:"name"`
}

// NetworkPut represents the modifiable fields of a LXD network
//
// API extension: network
type NetworkPut struct {
	Config map[string]string `json:"config" yaml:"config"`

	// API extension: entity_description
	Description string `json:"description" yaml:"description"`
}

// Network represents a LXD network
type Network struct {
	NetworkPut `yaml:",inline"`

	Name   string   `json:"name" yaml:"name"`
	Type   string   `json:"type" yaml:"type"`
	UsedBy []string `json:"used_by" yaml:"used_by"`

	// API extension: network
	Managed bool `json:"managed" yaml:"managed"`

	// API extension: clustering
	Status    string   `json:"status" yaml:"status"`
	Locations []string `json:"locations" yaml:"locations"`
}

// Writable converts a full Network struct into a NetworkPut struct (filters read-only fields)
func (network *Network) Writable() NetworkPut {
	return network.NetworkPut
}

// NetworkLease represents a DHCP lease
//
// API extension: network_leases
type NetworkLease struct {
	Hostname string `json:"hostname" yaml:"hostname"`
	Hwaddr   string `json:"hwaddr" yaml:"hwaddr"`
	Address  string `json:"address" yaml:"address"`
	Type     string `json:"type" yaml:"type"`
}

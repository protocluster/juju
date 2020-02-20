// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package network

import (
	"net"
	"sort"
	"strings"

	"github.com/juju/collections/set"
	"github.com/juju/errors"
)

// FanCIDRs describes the subnets relevant to a fan network.
type FanCIDRs struct {
	// FanLocalUnderlay is the CIDR of the local underlying fan network.
	// It allows easy identification of the device the fan is running on.
	FanLocalUnderlay string

	// FanOverlay is the CIDR of the complete fan setup.
	FanOverlay string
}

func newFanCIDRs(overlay, underlay string) *FanCIDRs {
	return &FanCIDRs{
		FanLocalUnderlay: underlay,
		FanOverlay:       overlay,
	}
}

// SubnetInfo is a source-agnostic representation of a subnet.
// It may originate from state, or from a provider.
type SubnetInfo struct {
	// CIDR of the network, in 123.45.67.89/24 format.
	CIDR string `json:"cidr" yaml:"cidr"`

	// Memoized value for the parsed network for the above CIDR.
	parsedCIDRNetwork *net.IPNet

	// ProviderId is a provider-specific subnet ID.
	ProviderId Id `json:"provider-id,omitempty" yaml:"provider-id,omitempty"`

	// ProviderSpaceId holds the provider ID of the space associated
	// with this subnet. Can be empty if not supported.
	ProviderSpaceId Id `json:"provider-space-id,omitempty" yaml:"provider-space-id,omitempty"`

	// ProviderNetworkId holds the provider ID of the network
	// containing this subnet, for example VPC id for EC2.
	ProviderNetworkId Id `json:"provider-network-id,omitempty" yaml:"provider-network-id,omitempty"`

	// VLANTag needs to be between 1 and 4094 for VLANs and 0 for
	// normal networks. It's defined by IEEE 802.1Q standard, and used
	// to define a VLAN network. For more information, see:
	// http://en.wikipedia.org/wiki/IEEE_802.1Q.
	VLANTag int `json:"vlan-tag" yaml:"vlan-tag"`

	// AvailabilityZones describes which availability zones this
	// subnet is in. It can be empty if the provider does not support
	// availability zones.
	AvailabilityZones []string `json:"zones,omitempty" yaml:"zones,omitempty"`

	// SpaceID is the id of the space the subnet is associated with.
	// Default value should be AlphaSpaceId. It can be empty if
	// the subnet is returned from an networkingEnviron. SpaceID is
	// preferred over SpaceName in state and non networkingEnviron use.
	SpaceID string `json:"space-id,omitempty" yaml:"space-id,omitempty"`

	// SpaceName is the name of the space the subnet is associated with.
	// An empty string indicates it is part of the AlphaSpaceName OR
	// if the SpaceID is set. Should primarily be used in an networkingEnviron.
	SpaceName string `json:"space-name,omitempty" yaml:"space-name,omitempty"`

	// FanInfo describes the fan networking setup for the subnet.
	// It may be empty if this is not a fan subnet,
	// or if this subnet information comes from a provider.
	FanInfo *FanCIDRs `json:"fan-info,omitempty" yaml:"fan-info,omitempty"`

	// IsPublic describes whether a subnet is public or not.
	IsPublic bool `json:"is-public,omitempty" yaml:"is-public,omitempty"`
}

// SetFan sets the fan networking information for the subnet.
func (s *SubnetInfo) SetFan(underlay, overlay string) {
	s.FanInfo = newFanCIDRs(overlay, underlay)
}

// FanLocalUnderlay returns the fan underlay CIDR if known.
func (s *SubnetInfo) FanLocalUnderlay() string {
	if s.FanInfo == nil {
		return ""
	}
	return s.FanInfo.FanLocalUnderlay
}

// FanOverlay returns the fan overlay CIDR if known.
func (s *SubnetInfo) FanOverlay() string {
	if s.FanInfo == nil {
		return ""
	}
	return s.FanInfo.FanOverlay
}

// Validate validates the subnet, checking the CIDR, and VLANTag, if present.
func (s *SubnetInfo) Validate() error {
	if s.CIDR == "" {
		return errors.Errorf("missing CIDR")
	} else if _, err := s.ParsedCIDRNetwork(); err != nil {
		return errors.Trace(err)
	}

	if s.VLANTag < 0 || s.VLANTag > 4094 {
		return errors.Errorf("invalid VLAN tag %d: must be between 0 and 4094", s.VLANTag)
	}

	return nil
}

// ParsedCIDRNetwork returns the network represented by the CIDR field.
func (s *SubnetInfo) ParsedCIDRNetwork() (*net.IPNet, error) {
	// Memoize the CIDR the first time this method is called or if the
	// CIDR field has changed.
	if s.parsedCIDRNetwork == nil || s.parsedCIDRNetwork.String() != s.CIDR {
		_, ipNet, err := net.ParseCIDR(s.CIDR)
		if err != nil {
			return nil, err
		}

		s.parsedCIDRNetwork = ipNet
	}
	return s.parsedCIDRNetwork, nil
}

// IsValidCidr returns whether cidr is a valid subnet CIDR.
func IsValidCidr(cidr string) bool {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err == nil && ipNet.String() == cidr {
		return true
	}
	return false
}

// FindSubnetIDsForAvailabilityZone returns a series of subnet IDs from a series
// of zones, if zones match the zoneName.
//
// Returns an error if no matching subnets match the zoneName.
func FindSubnetIDsForAvailabilityZone(zoneName string, subnetsToZones map[Id][]string) ([]Id, error) {
	matchingSubnetIDs := set.NewStrings()
	for subnetID, zones := range subnetsToZones {
		zonesSet := set.NewStrings(zones...)
		if zonesSet.Contains(zoneName) {
			matchingSubnetIDs.Add(string(subnetID))
		}
	}

	if matchingSubnetIDs.IsEmpty() {
		return nil, errors.NotFoundf("subnets in AZ %q", zoneName)
	}

	sorted := make([]Id, matchingSubnetIDs.Size())
	for k, v := range matchingSubnetIDs.SortedValues() {
		sorted[k] = Id(v)
	}

	return sorted, nil
}

// SubnetSet represents the classic "set" data structure, and contains Id.
// SubnetSet is used as a typed version to prevent string -> Id -> string
// conversion when using set.Strings
type SubnetSet map[Id]struct{}

// MakeSubnetSet creates and initializes a SubnetSet and populates it with
// initial values as specified in the parameters.
func MakeSubnetSet(values ...Id) SubnetSet {
	set := make(map[Id]struct{}, len(values))
	for _, id := range values {
		set[id] = struct{}{}
	}
	return set
}

// Add puts a value into the set.
func (s SubnetSet) Add(value Id) {
	s[value] = struct{}{}
}

// Size returns the number of elements in the set.
func (s SubnetSet) Size() int {
	return len(s)
}

// IsEmpty is true for empty or uninitialized sets.
func (s SubnetSet) IsEmpty() bool {
	return len(s) == 0
}

// Contains returns true if the value is in the set, and false otherwise.
func (s SubnetSet) Contains(id Id) bool {
	_, exists := s[id]
	return exists
}

// Values returns an unordered slice containing all the values in the set.
func (s SubnetSet) Values() []Id {
	result := make([]Id, len(s))
	i := 0
	for key := range s {
		result[i] = key
		i++
	}
	return result
}

// SortedValues returns an ordered slice containing all the values in the set.
func (s SubnetSet) SortedValues() []Id {
	values := s.Values()
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}

// InFan describes a network fan type.
const InFan = "INFAN"

// FilterInFanNetwork filters out any fan networks.
func FilterInFanNetwork(networks []Id) []Id {
	var result []Id
	for _, network := range networks {
		if !strings.Contains(network.String(), InFan) {
			result = append(result, network)
		}
	}
	return result
}

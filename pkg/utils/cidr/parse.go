package cidr

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

const (
	minIPV4CIDRBlockSize = 28
	maxIPV4CIDRBlockSize = 16
)

func GetIPV4GatewayNetmask(cidrBlock string) (string, string, error) {
	// The IPV4 CIDR block is of the format ip-addr/netmask
	ip, ipNet, err := net.ParseCIDR(cidrBlock)
	if err != nil {
		return "", "", errors.Wrapf(err,
			"getIPV4GatewayNetmask engine: unable to parse response for ipv4 cidr: '%s' from instance metadata", cidrBlock)
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return "", "", newParseIPV4GatewayNetmaskError("getIPV4GatewayNetmask", "engine",
			fmt.Sprintf("unable to parse ipv4 gateway from cidr block '%s'", cidrBlock))
	}

	maskOnes, _ := ipNet.Mask.Size()
	// As per
	// http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Subnets.html#VPC_Sizing
	// You can assign a single CIDR block to a VPC. The allowed block size
	// is between a /16 netmask and /28 netmask. Verify that
	if maskOnes > minIPV4CIDRBlockSize {
		return "", "", errors.Errorf("eni ipv4 netmask: invalid ipv4 cidr block, %d > 28", maskOnes)
	}
	if maskOnes < maxIPV4CIDRBlockSize {
		return "", "", errors.Errorf("eni ipv4 netmask: invalid ipv4 cidr block, %d <= 16", maskOnes)
	}

	// ipv4 gateway is the first available IP address in the subnet
	ip4[3] = ip4[3] + 1
	return ip4.String(), fmt.Sprintf("%d", maskOnes), nil
}

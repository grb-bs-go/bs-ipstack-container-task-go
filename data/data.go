package data

import (
	//"geoIP/logger"
	//"fmt"
	//"strconv"
	"inet.af/netaddr" // https://tailscale.com/blog/netaddr-new-ip-type-for-go
)

type GeoIP struct {
	IP        netaddr.IP
	Latitude  string
	Longitude string
}


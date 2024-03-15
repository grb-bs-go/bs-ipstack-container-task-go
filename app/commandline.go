package app

import (
	"encoding/json"
	"fmt"
	"geoIP/api"
	"geoIP/logger"
	"os"

	"inet.af/netaddr" // https://tailscale.com/blog/netaddr-new-ip-type-for-go
)

const UsageString string = "Usage: geoIP ip-address access-key" // to ease command line testing

func Start() string {
	// Read command-line args
	argsWithProg := os.Args
	logger.Log.Println("Command Line", argsWithProg)

	if len(argsWithProg) != 3 {
		fmt.Println(UsageString)
		fmt.Println("Example (IPv4): geopIP 80.44.77.120 a4js4jd2ld3eddKd3d")
		fmt.Println("Example (IPv6): geopIP 0:0:0:0:0:ffff:502c:4d78 a4js4jd2ld3eddKd3d")
		fmt.Println("JSON Response: {\"IP\":\"80.44.77.120\",\"Latitude\":\"52.569950103759766\",\"Longitude\":\"1.1133400201797485\"}")
		logger.Log.Println("Invalid command line input", argsWithProg)
		return UsageString
	}

	ip, err := netaddr.ParseIP(argsWithProg[1])
	if err != nil {
		fmt.Println(UsageString)
		logger.Log.Println(err)
		return UsageString
	}
	logger.Log.Println("Input IP is", ip)

	// Call API
	ipMeta, err := api.FindLocationFromIP(ip, argsWithProg[2])
	if err != nil {
		fmt.Println("Error calling apistack API", err)
		logger.Log.Println("Error calling apistack API", err)
		return UsageString
	}
	logger.Log.Println(ipMeta) // Struct output containing IP/Lat/Long values
	// We could output Struct data but for exercise convert output data to JSON
	output, err := json.Marshal(&ipMeta)
	if err != nil {
		fmt.Println("Error converting struct to JSON", err)
		logger.Log.Println("Error converting struct to JSON", err)
		return UsageString
	}
	// Command Line JSON response containing IP/Lat/Long values
	fmt.Println(string(output))

	return fmt.Sprint(string(output))
}

/*
 * API Package manages REST Interface to https://ipstack.com/
 *
 */
package api

import (
	"encoding/json"
	"fmt"
	"geoIP/data"
	"geoIP/logger"
	"inet.af/netaddr" // https://tailscale.com/blog/netaddr-new-ip-type-for-go
	"io"
	"net/http"
)

// In a 'real' impl this would be stored securely, Cloud Secret Store, Vault etc
const APIEndpoint = "http://api.ipstack.com/" // HTTPS not supported on Free Plan

func FindLocationFromIP(ip netaddr.IP, accessKey string) (data.GeoIP, error) {
	logger.Log.Println("findLocationFromIP ", ip)
	ipMeta := data.GeoIP{}

	// Call API
	targetURL := APIEndpoint + ip.String() + "?access_key=" + accessKey
	logger.Log.Println("Calling ", targetURL)
	// Test for protocol-level messaging error
	resp, err := http.Get(targetURL)
	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
		if resp == nil {
			logger.Log.Println("API connectivity issue")
		} else {
			logger.Log.Println("API HTTP error", err, "StatusCode", resp.StatusCode, "Status", resp.Status)
		}
		return ipMeta, err
	}
	defer resp.Body.Close()

	// Test for API Domain-level error
	var responseMap map[string]any
	readResponseMessage(resp.Body, &responseMap)
	succeedFlag, exists := responseMap["success"]
	if exists || succeedFlag == "false" {
		logger.Log.Printf(" - Error Response: %v", responseMap["error"])
		err := fmt.Errorf(" - Error Response: %v", responseMap["error"])
		return ipMeta, err
	}

	// Process a sunny day response
	logger.Log.Printf("SUCCESS Payload: %v", responseMap)
	ipMeta.IP = ip
	ipMeta.Latitude = fmt.Sprintf("%v", responseMap["latitude"])
	ipMeta.Longitude = fmt.Sprintf("%v", responseMap["longitude"])

	return ipMeta, err
}

// package conversion function
func readResponseMessage(response io.ReadCloser, responseMap *map[string]any) map[string]any {

	json.NewDecoder(response).Decode(&responseMap)
	logger.Log.Println("PAYLOAD")
	for k, v := range *responseMap {
		logger.Log.Printf("k: %s, v: %v\n", k, v)
	}
	return *responseMap
}

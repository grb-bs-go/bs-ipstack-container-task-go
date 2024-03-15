package main

import (
	"geoIP/app"
	"geoIP/logger"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Placeholder only for an extensive Go UnitTest Suite in a real project...
func TestNoCommandLineArgsRainyDay(t *testing.T) {
	logger.Log.Println("Running test TestNoCommandLineArgsRainyDay(t *testing.T) ...")
	os.Args = []string{}
	assert.Equal(t, app.UsageString, app.Start())
}

// One Command Line Arg, Program
func TestOneCommandLineArgRainyDay(t *testing.T) {
	logger.Log.Println("Running test TestOneCommandLineArgRainyDay(t *testing.T) ...")
	os.Args = []string{"geoIP"}
	assert.Equal(t, app.UsageString, app.Start())
}

// Two Command Line Args, Program
func TestTwoCommandLineArgRainyDay(t *testing.T) {
	logger.Log.Println("Running test estTwoCommandLineArgRainyDay(t *testing.T) ...")
	os.Args = []string{"geoIP","1.2.3.4"}
	assert.Equal(t, app.UsageString, app.Start())
}

// Four Command Line Args, too many
func TestFourCommandLineArgsRainyDay(t *testing.T) {
	logger.Log.Println("Running test TestFourCommandLineArgsRainyDay(t *testing.T) ...")
	os.Args = []string{"geoIP","1.2.3.4","token", "IPv6"}
	assert.Equal(t, app.UsageString, app.Start())
}

// Invalid value for an IP Address
func TestInvalidIPRainyDay(t *testing.T) {
	logger.Log.Println("Running test TestInvalidIPRainyDay(t *testing.T) ...")
	os.Args = []string{"geoIP","NotAnIPAddress","1234"}
	assert.Equal(t, app.UsageString, app.Start())
}




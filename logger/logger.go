package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	// set location of log file
	var logpath = "buildlogger.log"

	// flag.Parse()
	var logfile, err = os.Create(logpath)

	if err != nil {
		fmt.Println("Logger init failed!")
		os.Exit(1)
	}
	Log = log.New(logfile, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}

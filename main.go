/*
 * Simple Test Program for BS Interview Task.
 * @author: grb-bs-go
 * @date: 16/03/2024
 */
package main

import (
	"geoIP/app"
	"geoIP/logger"
	"time"
)

func main() {
	logger.Log.Println("Start geoIP - Time", time.Now())
	app.Start()
	logger.Log.Println("End geoIP - Time", time.Now())
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/showwin/speedtest-go/speedtest"
)

// create prometheus metrics
var dlSpeedMetric = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "download_speed", Help: "Most recent download speed measured by speedtest cli"})

func recordMetrics() {
	go func() {

		//run a loop forever
		for {
			//Declare a new speed tester
			var speedtestClient = speedtest.New()

			//Debug: print out something about the user
			//fmt.Println("Debug: User ISP = " + user.Isp)

			//Fetch a list of servers, ordered by nearest geographical location.
			//note, if using a VPN, the nearest geographical server will be in relation to the VPN server.
			serverList, err := speedtestClient.FetchServers()

			//Test for an error
			if err != nil {
				log.Fatal(err)
			}

			//Take the first server in the list
			nearestServer := serverList[0]

			//test the server
			nearestServer.TestAll()

			//set the gauge values for prometheus to scrape
			dlSpeedMetric.Set(nearestServer.DLSpeed)
			fmt.Println("DEBUG: downloadSpeed = " + fmt.Sprintf("%f", nearestServer.DLSpeed))

			//wait a bit so we don't spam the test server.
			//time.Sleep(60 * time.Second)
		}
	}()
}

func main() {

	//call the function to update the gauge
	recordMetrics()
	fmt.Println("registered metric")

	//create the /metrics endpoint for prometheus
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {

	//Declare a new speed tester
	var speedtestClient = speedtest.New()

	//Fetch a list of servers, ordered by nearest geographical location
	serverList, err := speedtestClient.FetchServers()

	//Test for an error
	if err != nil {
		fmt.Println(err)
	}

	//Debug: Printg out length of list of returned servers
	fmt.Println(serverList.Len())

	//Take the first server in the list
	nearestServer := serverList[0]

	//Debug: Print out something about the server chosen
	fmt.Println(nearestServer.Distance)

	//The download speed will be calculated and stored in the Sserver object for us to call later
	nearestServer.DownloadTest()

	//Print out the download speed from this server
	fmt.Println(nearestServer.DLSpeed)

	//Test to see if this stuff works
	fmt.Printf("This is a test!\n")
}

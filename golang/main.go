package main

import (
	"fmt"
	"sync"

	"github.com/showwin/speedtest-go/speedtest"
)

func printSpeeds(server *speedtest.Server, i int, waitGroup *sync.WaitGroup) {
	//build up the info about a server using the attached functions
	server.TestAll()

	//once this function has finished running, decrement the wait group counter
	defer waitGroup.Done()
	//Print out representation of server
	fmt.Printf("Server %v:\n", i)
	fmt.Printf("\t%v\n", server.String())
	fmt.Printf("\tDowload Speed: %v\n", server.DLSpeed)
	fmt.Printf("\tUpload Speed: %v\n", server.ULSpeed)
	fmt.Printf("\tLatency: %v\n", server.Latency)
	fmt.Println()
}

func main() {

	//Declare a new speed tester
	var speedtestClient = speedtest.New()

	//Find out stuff about the user
	user, err := speedtestClient.FetchUserInfo()

	if err != nil {
		fmt.Println("ERROR Found:")
		fmt.Println(err)
	}

	//Debug: print out something about the user
	fmt.Println("Debug: User ISP = " + user.Isp)

	//Fetch a list of servers, ordered by nearest geographical location
	serverList, err := speedtestClient.FetchServers()

	//Test for an error
	if err != nil {
		fmt.Println("ERROR Found:")
		fmt.Println(err)
	}

	//Debug: Printg out length of list of returned servers
	fmt.Printf("Debug: number of servers found = %v\n", serverList.Len())

	//Take the first server in the list
	//nearestServer := serverList[0]

	//create a wait group to allow goroutines to create threads of work
	var waitGroup sync.WaitGroup

	//loop through the servers in the returned server list
	for i, server := range serverList {

		// //build up the info about a server using the attached functions
		// server.TestAll()

		// //Print out JSON representation of server
		// fmt.Printf("Server %v:\n", i)
		// fmt.Printf("\t%v\n", server.String())
		// fmt.Printf("\tDowload Speed: %v\n", server.DLSpeed)
		// fmt.Printf("\tUpload Speed: %v\n", server.ULSpeed)
		// fmt.Printf("\tLatency: %v\n", server.Latency)
		// fmt.Println()

		//increment the wait group counter
		waitGroup.Add(1)

		//call a function that will start in its own goroutine
		go printSpeeds(server, i, &waitGroup)

		if i >= 4 {
			break
		}
	}

	//close the wait group
	waitGroup.Wait()
}

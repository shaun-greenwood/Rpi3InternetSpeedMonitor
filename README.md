# Rpi3InternetSpeedMonitor
Dedicated docker containers running Grafana, Prometheus and a custom Golang program to monitor internet speeds.
These containers are built using docker-compose.

## Installation
### Getting started
Clone this entire repository to your chosen directory on your Raspberry Pi 3.

### Docker and Docker Compose
Installing docker-compose on ARM architechture can be tricky.  Some guides online will say that this can only be achieved using pip.  However, this will only install docker-compose version 1 which is incompatible with the latest versions of docker.
To get around this, I have used a Docker Engine install script I found on Docker's website that automatically detects your version of the Raspberry Pi OS and gets the correct version of Docker and Docker Compose for you.
To install Docker and Docker Compose, simply run:

`$ ./Rpi3InternetSpeedMonitor/get-docker.sh`

### Configuration
This speed test monitor will store all speed tests so that you may look back over time and asses your internet's performance.  To that end, the Raspberry Pi's internal storage may not be enough and you may wish to choose an attached storage facility to store all of this data.  If you have an externally mounted directory you would like to use, edit the following file:

`docker-compose.override.yaml`

### Spin up Containers
Simply run the following command from the same directory that your docker-compose.yaml file is in:

`docker-compose up`

This should immediately start running three containers: the Golang speed monitor, the Prometheus database and the Grafana dashboards.

## Usage
This project was designed to be run in kiosk mode on a Raspberry Pi 3 with a 3.5 inch display attached.  However, there is no reason you couldn't use it as a headless system and simply browse to the webpage exposed by the Grafana container from whatever device you like.

If you would like to get the kiosk running on your 3.5 inch screen, run the following command:

`$ chromium-browser http://localhost:3000 --kiosk --window-size=480,320`

If you would like to view the dashboard from any other device on the same network you can use:

`http://[ip of pi]:3000`

## Golang Internet Speed Test
I used the speed test go package from showwin on github to create this speed test. The following two resources were invaluable in understanding how to use the package:

- [package description](https://pkg.go.dev/github.com/showwin/speedtest-go/speedtest#Server.DownloadTest)
- [example from README](https://github.com/showwin/speedtest-go/blob/master/example/main.go)


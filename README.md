# Rpi3InternetSpeedMonitor
Dedicated docker containers running Grafana, Prometheus and a custom Golang program to monitor internet speeds.

## Golang Internet Speed Test
I used the speed test go package from showwin on github to create this speed test. The following two resources were invaluable in understanding how to use the package:

- [package description](https://pkg.go.dev/github.com/showwin/speedtest-go/speedtest#Server.DownloadTest)
- [example from README](https://github.com/showwin/speedtest-go/blob/master/example/main.go)


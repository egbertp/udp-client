package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/egbertp/udp-client/helpers"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

// Variables to identiy the build
var (
	CommitHash string
	VersionTag string
	BuildTime  string
)

func main() {

	var (
		address = flag.String("address", "", "IP address and/or hostname")
		port    = flag.String("port", "9005", "UDP port")
	)
	flag.Parse()

	log.Printf("The version is: %s; the commit hash is: %s. Build time is: %s", VersionTag, CommitHash, helpers.ParseBuildTime(BuildTime).Format(time.RFC1123))

	// Combine address and port to the parsed address
	parsedAddress := *address + ":" + *port

	ServerAddr, err := net.ResolveUDPAddr("udp", parsedAddress)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		log.Printf("Sending packet with payload: %s to %s\n", msg, parsedAddress)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}

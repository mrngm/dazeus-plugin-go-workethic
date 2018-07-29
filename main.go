package main

import (
	"flag"
	"fmt"

	"github.com/dazeus/dazeus-go"
)

func main() {
	var toDaZeus = flag.String("dzconn", "unix:/tmp/dazeus.sock", "Set the connection parameters for DaZeus")
	var toChannel = flag.String("channel", "#example", "Send the messages to this channel on the first network")
	flag.Parse()
	fmt.Printf("Connecting to DaZeus on %s\n", *toDaZeus)

	dz, err := dazeus.ConnectWithLoggingToStdErr(*toDaZeus)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Going to send motivational messages to channel %s\n", *toChannel)

	networks, _ := dz.Networks()
	// Change this to your needs. It's partly hardcoded, yeah. Bite me.
	WorkEthics(networks[0], *toChannel, dz)
}

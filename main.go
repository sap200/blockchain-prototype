package main

import (
	// "blockchain/rest"
	// "os"
	// "strconv"
	// "blockchain/node"
	"blockchain/p2p"
	"os"
	"strconv"
	"log"
	"blockchain/types"
)



func main() {
	ip := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln("Invalid port")
	}
	// node := node.New()
	// api := rest.New(ip, uint16(port), node)
	// api.HandleRequests()
	alice := p2p.NewP2p(ip, port)
	if port != types.INITPORT {
		// dial to get peers list
		p2p.ConnectToServer(types.NewSocketConnector(types.INITIP, types.INITPORT))
	}

	alice.LaunchServer()

}


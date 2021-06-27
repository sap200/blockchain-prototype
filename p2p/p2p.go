package p2p

import (
	"fmt"
	"net"
	"log"
	"blockchain/types"
	"bufio"
)

type P2p struct {
	SocketConnector types.SocketConnector
	PeerAddressList []net.Conn
	peerSocketConnector []types.SocketConnector
}

func NewP2p(ip string, port int) P2p {
	return P2p{
		SocketConnector: types.NewSocketConnector(ip, port),
		PeerAddressList: []net.Conn{},
		peerSocketConnector: []types.SocketConnector{},
	}
}


func (p *P2p) LaunchServer() {
	fmt.Println("Listening on", p.SocketConnector.Address())
	ln, err := net.Listen("tcp", p.SocketConnector.Address())
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go p.handleIncomingConnection(conn)
	}
}

func ConnectToServer(s types.SocketConnector) {
	conn, err := net.Dial("tcp", s.Address())
	if err != nil {
		fmt.Println(err)
		return
	}
	// write a message to the connection
	go handleOutgoingConnection(conn)
}

func handleOutgoingConnection(conn net.Conn) {
	fmt.Println(conn, "Yes I am able to communicate")
	// REad from the connection
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	// Print what read
	fmt.Println(status)
}

func (p *P2p) handleIncomingConnection(conn net.Conn) {
	// Whenever there is an incoming connection fill the peer list with the given connection
	p.PeerAddressList = append(p.PeerAddressList, conn)
	fmt.Println("Peer-Address-List: ", p.PeerAddressList)
	fmt.Println("New connection from", conn.RemoteAddr())
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	// Print what read
	fmt.Println(status)
	// Write to connection
	fmt.Fprintf(conn, "Yes you connected to me.")
}


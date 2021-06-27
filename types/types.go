package types

import "strconv"

type SocketConnector struct {
	Ip string
	Port int
}

type Message struct {
	SenderConnector SocketConnector
	MsgType MSGType
	Data string
}

type MSGType int

const (
	 DISCOVERY MSGType = 0
)

const INITIP = "127.0.0.1"
const INITPORT = 10001

func NewSocketConnector(ip string, port int) SocketConnector {
	sc := SocketConnector{
		Ip: ip,
		Port: port,
	}

	return sc
}

func NewMessage(sc SocketConnector, type_ MSGType, data string) Message {
	msg := Message {
		SenderConnector: sc,
		MsgType: type_,
		Data: data,
	}

	return msg
}

func (s SocketConnector) Equals(s1 SocketConnector) bool {
	return (s.Ip == s1.Ip) && (s.Port == s1.Port)
}

func (s SocketConnector) Address() string {
	port := strconv.Itoa(s.Port)
	address := s.Ip + ":" + port
	return address
}

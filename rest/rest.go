package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"blockchain/node"
)

type RestAPI struct {
	Ip   string
	Port uint16
	Node node.Node

}

func New(ip string, port uint16, node node.Node) RestAPI {
	api := RestAPI{
		Ip:   ip,
		Port: port,
		Node: node,
	}

	return api
}

func homePage(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<h1>This is an API for the blockchain following methods are available </h1> <br><br>
                    a) <b style="color: red"> /blockchain</b>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;:returns the information about the blockchain<br>
                    b) <b style="color: red">/txPool</b>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;:returns the information about the current status of the transaction pool<br>
                    c) <b style="color: red">/issueTx</b>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;:It is a post request and is used to issue a transaction, a json file must be sent with a transaction field containing transaction<br>
	</html>
	`
	fmt.Fprintf(w, str)
}

func (api RestAPI) blockchain(w http.ResponseWriter, r *http.Request) {
	data := api.Node.Blockchain.ToJson()
	fmt.Fprintf(w, data)
}

func (api RestAPI) txPool(w http.ResponseWriter, r *http.Request) {
	data := api.Node.TxPool.ToJson()
	fmt.Fprintf(w, data)
}

func (api RestAPI) HandleRequests() {
	address := api.Ip + ":" + strconv.Itoa(int(api.Port))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/blockchain", api.blockchain)
	http.HandleFunc("/txPool", api.txPool)
	log.Fatal(http.ListenAndServe(address, nil))
}

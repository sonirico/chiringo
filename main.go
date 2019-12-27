package main

func main() {
	done := make(chan struct{})
	// the blockchain
	chain := NewChain()
	// the actual running node
	node := NewNode(chain, 256)
	// http web server
	http := newServer(node)
	http.setUp()
	// websocket web server
	ws := newWsServer(node)
	ws.setUp()
	go node.Run()
	go ws.Serve(config.WsPort)
	go http.Serve(config.HttpPort)
	<-done
}

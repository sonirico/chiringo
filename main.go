package main

func main() {
	done := make(chan struct{})
	node := NewNode(NewChain())
	web := newServer(node)
	web.setUp()
	go web.Serve()
	<-done
}

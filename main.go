package main

func main() {
	done := make(chan struct{}, 1)
	node := NewNode(NewChain())
	go node.ServeHTTP()
	<-done
}

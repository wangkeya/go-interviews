package main

func main() {
	var achan chan struct{}

	achan <- struct{}{}
}

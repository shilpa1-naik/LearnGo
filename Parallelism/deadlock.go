package main
import "fmt"

// The sending operation is a blocking operation and requires the receive channel to be ready before sending data to the channel. We’ll learn more about this when we’ll study channels in the next chapter!
func main() {
	mychannel := make(chan int)
	mychannel <- 10
	fmt.Println(<-mychannel)
}

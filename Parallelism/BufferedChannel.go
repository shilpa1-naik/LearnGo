package main
import "fmt"

// DEADLOCK
/*
func main() {
	mychannel := make(chan int)
	mychannel <- 10
	fmt.Println(<-mychannel)
        }
*/


package main
import "fmt"

func main() {
	mychannel := make(chan int)
	go func(){
		mychannel <- 10
	}()
	fmt.Println(<-mychannel)
}

// Buffered Channels are channels with a capacity/buffer. They can be created with the following syntax:

// No deadlock now
package main
import "fmt"

func main() {
	mychannel := make(chan int, 2)
	mychannel <- 10
	fmt.Println(<-mychannel)
}


// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

package main

import "fmt"

func main() {
	messages := make(chan string)
	
	// Here's a non-blocking receive. If a value is
	// available on the `messages` channel then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//this time we're creating a channel with the buffer of 2
	//and sending some messages to it
	notifications := make(chan string, 2)
	notifications <- "hello"
	notifications <- "hello again"

	//if there is something in the channel, we receive the mesage
	//otherwise default will be shown
	select {
	case msg := <-notifications:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	//We populated both channels. We are listening
	//on both channels using select.
	//the first channel that we're able to 
	//receive from will be the one we'll take 
	//the data from. In this case messagesA
	messagesA := make(chan string, 1)
	signalsA := make(chan bool, 1)	
	messagesA <- "some message"
	signalsA <- true	

	select {
	case msg := <-messagesA:
		fmt.Println("received message:", msg)		
	case sig := <-signalsA:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}

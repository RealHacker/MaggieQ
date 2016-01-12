package maggieQ
import "cmd/vet/testdata"

type Consumer struct {
	Name string
	Queue *Queue
}

type Channel struct {
	ID int64
	Consumers map[string]*Consumer
	messages chan WireMessage
}

func (ch Channel) closeChannel() {
	// close the message channel, stop accepting message
	close(ch.messages)
	// Remove all the consumers from its queue
	for _, consumer := range ch.Consumers {
		consumer.Queue.RemoveConsumer(consumer)
	}
}

func (ch Channel) messageHandler() {

}
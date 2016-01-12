package maggieQ

import "sync"

// The global queues store
var Queues QueueMap

type QueueMap struct {
	Queues map[string]*Queue
	sync.Mutex
}

type Queue struct {
	Consumers []*Consumer
}


func (q Queue) AddConsumser(consumer *Consumer) {

}

func (q Queue) RemoveConsumer(consumer *Consumer){

}
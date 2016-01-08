package maggieQ

import "sync"

// The global exchange store
var Exchanges ExchangeMap

type ExchangeMap struct {
	Exchanges map[string]*Exchange
	mutex sync.Mutex
}

type Exchange struct {

}

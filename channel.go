package maggieQ

type Consumer struct {
	Name string

}

type Channel struct {
	Consumers map[string]*Consumer

}

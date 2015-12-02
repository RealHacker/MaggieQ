package maggieQ


type WireMessage struct {
	Type int8
	Channel int16
	Length int32
	Payload []byte
}


func (msg *WireMessage) Marshall() []byte {

}

func (msg *WireMessage) Unmarshall(data []byte) error{

}

type MethodFrame struct {
	ClassID int16
	MethodID int16
	Arguments map[string]interface{}
}

func (frame *MethodFrame) Marshall() []byte {

}

func (frame *MethodFrame) Unmarshall(data []byte) error {

}

type HeaderFrame struct {
	ClassID int16
	BodySize int64
	// TODO: body parameters here
}

func (frame *HeaderFrame) Marshall() []byte {

}

func (frame *HeaderFrame) Unmarshall(data []byte) error {

}
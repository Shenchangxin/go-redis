package reply

type PongReply struct {
}

var pongBytes = []byte("+Pong\r\n")

func (p PongReply) ToBytes() []byte {
	return pongBytes
}

func MakePongReply() *PongReply {
	return &PongReply{}
}

type OkReply struct {
}

var okBytes = []byte("+OK\r\n")

func (p OkReply) ToBytes() []byte {
	return okBytes
}

var theOkReply = new(OkReply)

func MakeOkReply() *OkReply {
	return theOkReply
}

type NullBulkReply struct {
}

var nullBulkBytes = []byte("+$-1\r\n")

func (p NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

var emptyMultiBulkBytes = []byte("*0\r\n")

type EmptyMultiBulkReply struct {
}

func (p EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

type NoReply struct {
}

var noBytes = []byte("")

func (n NoReply) ToBytes() []byte {
	return noBytes
}

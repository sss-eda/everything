package messaging

type Message struct {
	kind    byte
	payload []byte
}

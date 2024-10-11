package service

type Message struct {
}

type Notifier interface {
	SendMessage(message Message)
}

type SMSService struct {
}

func (s *SMSService) SendMessage(message Message) {
	// send message
}

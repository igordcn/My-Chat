package model

import (
	"errors"
	"fmt"
)

type Message struct {
	content  string
	sender   *User
	receiver *User
}

func NewMessage(content string, from *User, to *User) (*Message, error) {
	message := &Message{content, from, to}
	if err := message.validate(); err != nil {
		return nil, err
	}
	return message, nil
}

func (message *Message) Content() string {
	return message.content
}

func (message *Message) Sender() *User {
	return message.sender
}

func (message *Message) Receiver() *User {
	return message.receiver
}

func (message *Message) toString() string {
	delimiter := "###"
	return fmt.Sprint(message.sender.Id()) + delimiter + message.sender.Name() + delimiter + fmt.Sprint(message.receiver.Id()) + delimiter + message.receiver.Name() + delimiter + message.content
}

func (message *Message) validate() error {
	if len(message.content) == 0 {
		return errors.New("Message Content should not be empty!")
	}
	if message.sender == nil {
		return errors.New("Message Sender is required!")
	}
	if message.receiver == nil {
		return errors.New("Message Receiver is required!")
	}
	return nil
}

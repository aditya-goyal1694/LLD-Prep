package main

import "fmt"

type Notification interface {
    Send() error
}

type EmailNotification struct {
    msg string
}

type SMSNotification struct {
    msg string
}

type PushNotification struct {
    msg string
}

func (e *EmailNotification) Send() error {
    fmt.Printf("Sending EMAIL: %s\n", e.msg)
    return nil
}

func (s *SMSNotification) Send() error {
    fmt.Printf("Sending SMS: %s\n", s.msg)
    return nil
}

func (p *PushNotification) Send() error {
    fmt.Printf("Sending PUSH notification: %s\n", p.msg)
    return nil
}

func main() {
	service := NewNotificationService()

	notifications := []*Notification{
		&EmailNotification{
			msg: "Welcome to our platform!",
		},
		&SMSNotification{
			msg: "Your OTP is 123456",
		},
		&PushNotification{
			msg: "You have a new message",
		},
	}

	for _, n := range notifications {
		if err := n.Send(n); err != nil {
			fmt.Println(err)
		}
	}
}

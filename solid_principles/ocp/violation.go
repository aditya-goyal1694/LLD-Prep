package main

import "fmt"

type Notification struct {
	Type    string
	Message string
}

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) Send(n Notification) error {
	if n.Type == "email" {
		fmt.Printf(
			"Sending EMAIL: %s\n",
			n.Message,
		)
	} else if n.Type == "sms" {
		fmt.Printf(
			"Sending SMS: %s\n",
			n.Message,
		)
	} else if n.Type == "push" {
		fmt.Printf(
			"Sending PUSH notification: %s\n",
			n.Message,
		)
	} else {
		return fmt.Errorf(
			"unsupported notification type: %s",
			n.Type,
		)
	}

	return nil
}

func main() {
	service := NewNotificationService()

	notifications := []Notification{
		{
			Type:    "email",
			Message: "Welcome to our platform!",
		},
		{
			Type:    "sms",
			Message: "Your OTP is 123456",
		},
		{
			Type:    "push",
			Message: "You have a new message",
		},
	}

	for _, n := range notifications {
		if err := service.Send(n); err != nil {
			fmt.Println(err)
		}
	}
}

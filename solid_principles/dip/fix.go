package main

import (
	"fmt"
	"time"
)

type OrderRepository interface {
    SaveOrder(orderID string) error
}

type MySQLDatabase struct{}

func (db *MySQLDatabase) SaveOrder(orderID string) error {
	fmt.Printf("Saving order %s to MySQL\n", orderID)
	return nil
}

type OrderNotifier interface {
    SendConfirmation(orderID string) error
}

type EmailService struct{}

func (e *EmailService) SendConfirmation(orderID string) error {
	fmt.Printf("Sending confirmation email for order %s\n", orderID)
	return nil
}

type OrderService struct {
	db    Database
	notification Notification
}

func NewOrderService(db Database, notification Notification) *OrderService {
	return &OrderService{
		db:    db,
		notification: notification,
	}
}

func (s *OrderService) PlaceOrder(orderID string) error {
	if orderID == "" {
		return fmt.Errorf("invalid order id")
	}

	if err := s.db.SaveOrder(orderID); err != nil {
		return err
	}

	if err := s.notification.SendConfirmation(orderID); err != nil {
		return err
	}

	fmt.Printf(
		"Order %s placed successfully at %s\n",
		orderID,
		time.Now().Format(time.RFC3339),
	)

	return nil
}

func main() {
	service := NewOrderService(&MySQLDatabase{},&EmailService{})

	if err := service.PlaceOrder("ORD-1001"); err != nil {
		fmt.Println(err)
	}
}
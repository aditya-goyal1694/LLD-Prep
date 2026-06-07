package publisher

import (
    "fmt"
	"lld/design_patterns/observer/subscriber"
)

type Publisher interface {
    Subscribe(sub subscriber.Subscriber)
    Unsubscribe(sub subscriber.Subscriber)
    NotifyAll(title string)
}

type Channel struct {
    Name string
    subscribers []subscriber.Subscriber
}

func NewChannel(name string) *Channel {
    return &Channel{
        Name: name,
        subscribers: make([]subscriber.Subscriber, 0),
    }
}

func (c *Channel) Subscribe(sub subscriber.Subscriber) {
    c.subscribers = append(c.subscribers, sub)
}

func (c *Channel) Unsubscribe(sub subscriber.Subscriber) {
    newSubscribers := make([]subscriber.Subscriber, 0)
    for _, subs := range c.subscribers {
        if sub == subs {
            continue
        }
        newSubscribers = append(newSubscribers, subs)
    }
    c.subscribers = newSubscribers
}

func (c *Channel) NotifyAll(title string) {
    var errs []error
    for _, subs := range c.subscribers {
        if err := subs.(subscriber.Subscriber).Update(fmt.Sprintf("New video titled: %s", title)); err != nil {
            errs = append(errs, err)
        }
    }
    
    if len(errs) > 0 {
        fmt.Printf("Error in notifying. Errs:\n")
        for _, err := range errs {
            fmt.Println(err.Error())
        }
    }
}

func (c *Channel) UploadVideo(title string) {
    fmt.Printf("Video titled %s has been successfully uploaded!\n", title)
    fmt.Println("Notifying all the subscribers")
    c.NotifyAll(title)
    fmt.Println("Subscribers have been successfully notified!")
}